package grpc_node

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/io/grpc_node/handlers"
	"team01/internal/node/io/grpc_node/middleware"
	"team01/internal/proto/node"
	"team01/internal/util"
	"time"
)

type serv struct {
	serv *grpc.Server
	mw   middleware.IMiddleWare

	newNode        chan string
	log            chan string
	finished       chan bool
	controllerNode *handlers.NodeController
	controllerCli  *handlers.CliController
}

func NewGrpcNode(bl *bl.BL, fin chan bool) *serv {
	newNodes := make(chan string, 10)
	return &serv{
		mw:             middleware.NewMiddlewares(bl, newNodes),
		newNode:        newNodes,
		log:            make(chan string, 20),
		finished:       fin,
		controllerNode: handlers.NewNodeController(bl, newNodes),
		controllerCli:  handlers.NewCliController(bl),
	}
}

func (s *serv) run() error {
	s.serv = grpc.NewServer(grpc.UnaryInterceptor(s.mw.ServerRequestInterceptor))
	node.RegisterNodeCommunicationServer(s.serv, s.controllerNode)
	node.RegisterClientCommunicationServer(s.serv, s.controllerCli)

	lis, err := net.Listen("tcp", cfg.GetAddress())
	if err != nil {
		return err
	}
	err = s.serv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (s *serv) Run(ctx context.Context) {
	// горутина стартующая сервер
	go func() {
		defer cfg.GetLogger().Info("server stop")
		if err := s.run(); err != nil {
			cfg.GetLogger().Error("run error", zap.Error(err))
		}
		s.finished <- true
	}()

	// соединение с первой нодой, переданной при запуске
	address, err := cfg.GetConnectAddress()
	if err == nil {
		s.newNode <- address
	}
	status := make(chan string)
	go func() {
		for {
			time.Sleep(time.Second * 5)
			info := s.controllerNode.Bl.Node.GetInfo()
			status <- fmt.Sprintf(
				"Известных нод: %d, Размер хранилища: %d, Статус: %s",
				s.controllerNode.Bl.Node.GetSizeCluster(),
				info.SizeVault,
				info.Status.String())
		}

	}()

	// горутина обработчик событий
	go func() {
		for {
			select {
			case log := <-s.log:
				cfg.GetLogger().Info(log)

			case <-s.controllerNode.Bl.Node.TickerLastNode().C: // завершения таймера последней, по времени опроса ноды

				cfg.GetLogger().Info("ticker done for " + s.controllerNode.Bl.Node.GetUnit().LastNode.Address)
				for k, v := range s.controllerNode.Bl.Node.GetUnit().KnowNodes {
					cfg.GetLogger().Info(k + " : " + util.Timer(&v.Public))
				}
				_, err = s.controllerNode.Bl.Node.GetLastClient().Ping(ctx, &node.PingRequest{})
				if err != nil {
					cfg.GetLogger().Info(err.Error())
					continue
				}

			case nodeConnect := <-s.newNode: // подключение к новым нодам
				n, _ := context.WithTimeout(ctx, time.Second*5)
				err := s.connectTo(n, nodeConnect)
				if err != nil {
					cfg.GetLogger().Info(err.Error())
				}

			case stat := <-status:
				cfg.GetLogger().Info(stat)

			case <-ctx.Done(): // завершение работы текущей ноды
				cfg.GetLogger().Info("ctx done")
				// TODO послать запрос отключения ко всем нодам, закрыть соединения
				s.serv.GracefulStop()
				return

			}
		}
	}()
}

func (s *serv) connectTo(ctx context.Context, address string) error {
	if cfg.GetAddress() == address {
		return nil
	}

	conn, err := util.GetClient(ctx, address, s.log, s.mw.ClientRequestInterceptor)
	if err != nil {
		return err
	}
	nodeCli := node.NewNodeCommunicationClient(conn)

	s.controllerNode.Bl.Node.AddNodeToKnown(address, nodeCli)
	if err != nil {
		return err
	}
	s.controllerNode.Bl.Node.UpdLastNode()

	return nil
}
