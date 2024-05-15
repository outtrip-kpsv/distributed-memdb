package grpc

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/io/grpc/middleware"
	"team01/internal/node/io/grpc/node-rpc"
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
	server := grpc.NewServer(grpc.UnaryInterceptor(middleware.ServerRequestInterceptor))
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

// TODO work1!! this!!!
func (s *serv) ConnectTo(address string) error {
	if cfg.GetAddress() == address {
		return nil
	}

	if ok := s.nodeRpc.BL.Node.NodeIsKnown(address); !ok {
		//TODO named logger
		conn, err := getClient(address, s.nodeRpc.Midleware.ClientRequestInterceptor)
		if err != nil {
			return err
		}

		s.nodeRpc.BL.Node.AddNodeToKnown(address, node.NewNodeCommunicationClient(conn))

		//var st = State{
		//	Public:     node.DataNode{},
		//	Connection: conn,
		//}
		//
		//u.KnowNodes[address] = &st
		//u.KnowNodes[address].Public.Ts = timestamppb.Now()
		//u.KnowNodes[address].Client = node.NewNodeCommunicationClient(conn)

		cfg.GetLogger().Info("OK", zap.String("Connect to", address))

	} else {

		//stateTime := u.KnowNodes[srv].Public.Ts.AsTime()
		//newGetTime := timestamppb.Now()
		//if stateTime.Before(newGetTime) {
		//	u.KnowNodes[srv].Public.Ts = timeSt
		//	cfg.GetLogger().Info("time upd", zap.String("node", srv), zap.Reflect("time", timeSt))
		//}
	}

	if len(u.KnowNodes) == 1 {
		u.LastNode.Address = srv
		u.LastNode.Ticker = time.NewTicker(time.Second * 5)
	}

	return nil
}

func getClient(
	srv string,
	interceptor func(
		context.Context,
		string,
		interface{},
		interface{},
		*grpc.ClientConn,
		grpc.UnaryInvoker,
		...grpc.CallOption) error) (*grpc.ClientConn, error) {

	ticker := time.NewTicker(time.Millisecond)
	timeout := time.After(5 * time.Second)
	nSeconds := 1

	conn, _ := grpc.Dial(
		srv,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor),
	)

	for {
		select {
		case <-ticker.C:
			ticker.Stop()
			if conn.GetState() != connectivity.Ready {

				cfg.GetLogger().Info("Нет соединение с нодой:", zap.String("node", srv), zap.Duration("проверка подключения через", time.Duration(nSeconds)*time.Millisecond))
				tmp := time.Duration(nSeconds) * time.Millisecond
				ticker = time.NewTicker(tmp)
				nSeconds *= 2
				continue
			}
			return conn, nil
		case <-timeout:
			return nil, errors.New("node: connection timeout")
		}
	}
}
