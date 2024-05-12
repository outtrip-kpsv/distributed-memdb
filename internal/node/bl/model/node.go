package model

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	cfg "team01/internal/config"
	"team01/internal/node/db"
	"team01/internal/proto/node"
	"time"
)

type NodeLastInfo struct {
	Address string
	Ticker  *time.Ticker
}

type Unit struct {
	node.UnimplementedNodeCommunicationServer

	Address   string
	KnowNodes map[string]*State
	Vault     db.IVault
	LastNode  NodeLastInfo
}

type State struct {
	Public     node.DataNode
	Connection *grpc.ClientConn
	Client     node.NodeCommunicationClient
}

// ServerRequestInterceptor TODO move
// ServerRequestInterceptor middleware на стороне сервера
func (u *Unit) ServerRequestInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (response interface{}, err error) {
	//todo ??

	return handler(ctx, req)
}

func (u *Unit) ConnectTo(srv string, timeSt *timestamppb.Timestamp) error {
	if u.Address == srv {
		return nil
	}
	if _, ok := u.KnowNodes[srv]; !ok {
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
		stateTime := u.KnowNodes[srv].Public.Timestamp.AsTime()
		newGetTime := timeSt.AsTime()
		if stateTime.Before(newGetTime) {
			u.KnowNodes[srv].Public.Timestamp = timeSt
			cfg.GetLogger().Info("time upd", zap.String("node", srv), zap.Reflect("time", timeSt))
		}
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
		...grpc.CallOption) error, logger *zap.Logger) (*grpc.ClientConn, error) {

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

				logger.Info("Нет соединение с нодой:", zap.String("node", srv), zap.Duration("проверка подключения через", time.Duration(nSeconds)*time.Millisecond))
				tmp := time.Duration(nSeconds) * time.Millisecond
				//fmt.Println(tmp)
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

// CreateReqKnownNodes Функция для формирования knownNodes
func (u *Unit) CreateReqKnownNodes() *node.KnownNodes {
	knownNodesMap := make(map[string]*node.DataNode)
	for k, v := range u.KnowNodes {
		knownNodesMap[k] = &node.DataNode{Timestamp: v.Public.Timestamp}
	}
	// добавление скоего адресса в мапу
	knownNodesMap[u.Address] = &node.DataNode{
		Timestamp: timestamppb.Now(),
	}
	knownNodesMessage := &node.KnownNodes{
		Nodes: knownNodesMap,
	}
	return knownNodesMessage
}
