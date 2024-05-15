package grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/io/grpc/node-rpc"
	"team01/internal/node/io/grpc/util"
	"team01/internal/proto/node"
	"time"
)

type serv struct {
	nodeRpc *node_rpc.NodeRpc

	finished chan bool
}

func NewGrpcNode(bl *bl.BL, fin chan bool) *serv {
	return &serv{
		nodeRpc:  node_rpc.CreateNode(bl),
		finished: fin,
	}
}

func (s *serv) run() error {
	server := grpc.NewServer(grpc.UnaryInterceptor(s.nodeRpc.Midleware.ServerRequestInterceptor))
	node.RegisterNodeCommunicationServer(server, s.nodeRpc)

	lis, err := net.Listen("tcp", cfg.GetAddress())
	if err != nil {
		return err
	}
	err = server.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (s *serv) Run() {
	go func() {
		if err := s.run(); err != nil {
			cfg.GetLogger().Error("run error", zap.Error(err))
		}
		s.finished <- true
	}()
	cfg.GetLogger().Info("start node", zap.String("address", cfg.GetAddress()))
	address, err := cfg.GetConnectAddress()
	if err == nil {
		err = s.ConnectTo(address)
	}
	go func() {
		for {
			select {
			case <-s.nodeRpc.BL.Node.GetUnit().LastNode.Ticker.C:
				cfg.GetLogger().Info("ticker done for " + s.nodeRpc.BL.Node.GetUnit().LastNode.Address)
				pingReq, err := s.nodeRpc.BL.Node.GetUnit().KnowNodes[s.nodeRpc.BL.Node.GetUnit().LastNode.Address].Client.Ping(context.Background(), &node.PingRequest{})

				if err != nil {
					delete(s.nodeRpc.BL.Node.GetUnit().KnowNodes, s.nodeRpc.BL.Node.GetUnit().LastNode.Address)

					//s.nodeRpc.BL.Node.UpdLastNode()
					//cfg.GetLogger().Info("error", zap.Error(err))
					//delete(s.nodeRpc.BL.Node.GetUnit().KnowNodes, s.nodeRpc.BL.Node.GetUnit().LastNode.Address)
					//// todo обновить ласт ноду
					//oldTime := time.Now()
					//for k, v := range s.nodeRpc.BL.Node.GetUnit().KnowNodes {
					//
					//	cfg.GetLogger().Info("find time for last node check " + k)
					//	if v.Public.Ts.AsTime().Before(oldTime) {
					//		oldTime = v.Public.Ts.AsTime()
					//		s.nodeRpc.BL.Node.GetUnit().LastNode.Address = k
					//		cfg.GetLogger().Info("last node naw is " + k)
					//	} else {
					//		cfg.GetLogger().Info("ln old " + s.nodeRpc.BL.Node.GetUnit().LastNode.Address)
					//		//fmt.Println()
					//	}
					//}
					//s.nodeRpc.BL.Node.GetUnit().LastNode.Ticker.Stop()
					//tmp := time.Second*5 - (time.Now().Sub(oldTime))
					//cfg.GetLogger().Info("-------- " + tmp.String())
					////fmt.Println("t", tmp)
					//if tmp <= 0 {
					//	tmp = time.Millisecond
					//}
					//s.nodeRpc.BL.Node.GetUnit().LastNode.Ticker = time.NewTicker(tmp)
					continue
				} else {

				}

				if pingReq.Res == false {
					delete(s.nodeRpc.BL.Node.GetUnit().KnowNodes, s.nodeRpc.BL.Node.GetUnit().LastNode.Address)
				}

			default:
				continue

			}
		}
	}()
}

// TODO work1!! this!!!
func (s *serv) ConnectTo(address string) error {
	if cfg.GetAddress() == address {
		return nil
	}

	if ok := s.nodeRpc.BL.Node.NodeIsKnown(address); !ok {
		//TODO named logger, remove ctx background
		conn, err := util.GetClient(context.Background(), address, s.nodeRpc.Midleware.ClientRequestInterceptor)
		if err != nil {
			return err
		}
		s.nodeRpc.BL.Node.AddNodeToKnown(address, node.NewNodeCommunicationClient(conn))
		cfg.GetLogger().Info("OK", zap.String("Connect to", address))

	} else {

	}

	if len(s.nodeRpc.BL.Node.GetUnit().KnowNodes) == 1 {
		//todo getunit remove
		s.nodeRpc.BL.Node.GetUnit().LastNode.Address = address
		s.nodeRpc.BL.Node.GetUnit().LastNode.Ticker = time.NewTicker(time.Second * 5)
		cfg.GetLogger().Info("last node cr")
	}

	return nil
}
