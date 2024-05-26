package grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/io/grpc/node-rpc"
	"team01/internal/proto/node"
)

type serv struct {
	nodeRpc  *node_rpc.NodeRpc
	serv     *grpc.Server
	finished chan bool
}

func NewGrpcNode(bl *bl.BL, fin chan bool) *serv {
	return &serv{
		nodeRpc:  node_rpc.CreateNode(bl),
		finished: fin,
	}
}

func (s *serv) run() error {
	defer cfg.GetLogger().Info("f run end")

	s.serv = grpc.NewServer(grpc.UnaryInterceptor(s.nodeRpc.MiddleWare.ServerRequestInterceptor))
	node.RegisterNodeCommunicationServer(s.serv, s.nodeRpc)

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
		} else {
			cfg.GetLogger().Info("start node", zap.String("address", cfg.GetAddress()))
		}
		s.finished <- true
	}()

	// соединение с первой нодой, переданной при запуске
	address, err := cfg.GetConnectAddress()
	if err == nil {
		err = s.nodeRpc.ConnectTo(ctx, address)
	}
	s.nodeRpc.BL.Node.UpdLastNode()

	// горутина обработчик событий
	go func() {
		for {
			select {
			case <-s.nodeRpc.BL.Node.TickerLastNode().C:

				//ctxNew := context.WithValue(ctx, "from", cfg.GetAddress())
				cfg.GetLogger().Info("ticker done for " + s.nodeRpc.BL.Node.GetUnit().LastNode.Address)

				pingReq, err := s.nodeRpc.BL.Node.GetLastClient().Ping(ctx, &node.PingRequest{})
				if err != nil {
					// todo perepisat na deleteLastNode
					s.nodeRpc.BL.Node.DeleteNode(s.nodeRpc.BL.Node.GetUnit().LastNode.Address)
					continue
				}

				if pingReq.Res == false {
					// todo perepisat na deleteLastNode
					s.nodeRpc.BL.Node.DeleteNode(s.nodeRpc.BL.Node.GetUnit().LastNode.Address)
				}

			case <-ctx.Done():
				cfg.GetLogger().Info("ctx done")
				s.serv.Stop()

				return
			default:
			}
		}
	}()
}

//func () CloseConections()  {
//
//}

//func (s *serv) ConnectTo(ctx context.Context, address string) error {
//	if cfg.GetAddress() == address {
//		return nil
//	}
//	ctxNew, _ := context.WithTimeout(ctx, 5*time.Second)
//
//	if ok := s.nodeRpc.BL.Node.NodeIsKnown(address); !ok {
//		//TODO named logger, remove ctx background !!! <<<<<<<
//		conn, err := util.GetClient(ctxNew, address, s.nodeRpc.MiddleWare.ClientRequestInterceptor)
//		if err != nil {
//			return err
//		}
//		s.nodeRpc.BL.Node.AddNodeToKnown(address, node.NewNodeCommunicationClient(conn))
//		cfg.GetLogger().Info("OK", zap.String("Connect to", address))
//	}
//	return nil
//}
