package grpc

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/io/grpc/midleware"
	"team01/internal/node/io/grpc/node-rpc"
	"team01/internal/proto/node"
	"time"
)

type serv struct {
	nodeRpc  *node_rpc.NodeRpc
	finished chan bool
}

func NewGrpcNode(bl *bl.BL, fin chan bool) *serv {

	return &serv{
		nodeRpc:  node_rpc.CreateNode(bl),
		finished: fin,
	}

	//nodeServ := &model.NodeCore{}
	//
	////n := unit.NewNode(options.GetAddress(), options.Logger.Named("Node:"+options.GetAddress()))
	//s := grpc.NewServer(grpc.UnaryInterceptor(midleware.ServerRequestInterceptor))
	//
	////s := grpc.NewServer(grpc.UnaryInterceptor(n.ServerRequestInterceptor))
	//
	//node.RegisterNodeCommunicationServer(s, nodeServ)
	//
	////address, err := cfg.GetConnectAddress()
	////if err == nil {
	////	//todo ???
	////	err = bl.Node.GetUnit().ConnectTo(address, timestamppb.Now())
	////}
	////if err != nil {
	////	fmt.Println("Error", err)
	////}
	//
	//return nodeServ
}

func (s *serv) run() error {
	server := grpc.NewServer(grpc.UnaryInterceptor(midleware.ServerRequestInterceptor))
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
}

//TODO work1!! this!!!

func (s *serv) ConnectTo(address string) {
	if cfg.GetAddress() == srv {
		return nil
	}

	if _, ok := s.nodeRpc.BL.Node.GetUnit().KnowNodes[address]; !ok {
		//TODO named logger
		conn, err := getClient(srv, u.ClientRequestInterceptor, cfg.GetLogger().Named("rrr"))
		if err != nil {
			return err
		}

		var st = State{
			Public:     node.DataNode{},
			Connection: conn,
		}
		u.KnowNodes[srv] = &st

		u.KnowNodes[srv].Public.Ts = timestamppb.Now()
		u.KnowNodes[srv].Connection = conn
		u.KnowNodes[srv].Client = node.NewNodeCommunicationClient(conn)

		cfg.GetLogger().Info("OK", zap.String("Connect to", srv))
	} else {
		stateTime := u.KnowNodes[srv].Public.Ts.AsTime()
		newGetTime := timeSt.AsTime()
		if stateTime.Before(newGetTime) {
			u.KnowNodes[srv].Public.Ts = timeSt
			cfg.GetLogger().Info("time upd", zap.String("node", srv), zap.Reflect("time", timeSt))
		}
	}
	if len(u.KnowNodes) == 1 {
		u.LastNode.Address = srv
		u.LastNode.Ticker = time.NewTicker(time.Second * 5)
	}
}
