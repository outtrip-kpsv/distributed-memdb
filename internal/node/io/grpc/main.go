package grpc

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/proto/node"
)

type nodeCore struct {
	BL       *bl.BL
	finished chan bool
}

func NewGrpcNode(bl *bl.BL) *nodeCore {
	lis, err := net.Listen("tcp", cfg.GetAddress())

	if err != nil {
		cfg.GetLogger().Fatal("Failed to listen: %v", zap.Error(err))
	}
	//n := unit.NewNode(options.GetAddress(), options.Logger.Named("Node:"+options.GetAddress()))
	s := grpc.NewServer(grpc.UnaryInterceptor(bl.Node.GetUnit().ServerRequestInterceptor))

	//s := grpc.NewServer(grpc.UnaryInterceptor(n.ServerRequestInterceptor))

	address, err := cfg.GetConnectAddress()
	if err == nil {
		err = bl.Node.GetUnit().ConnectTo(address, timestamppb.Now())
	}
	if err != nil {
		fmt.Println("Error", err)
	}

	node.RegisterNodeCommunicationServer(s, n)

	return &nodeCore{BL: bl}

}
