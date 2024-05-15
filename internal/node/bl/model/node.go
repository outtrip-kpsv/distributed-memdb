package model

import (
	"team01/internal/node/db/mem"
	"team01/internal/proto/node"
	"time"
)

type NodeLastInfo struct {
	Address string
	Ticker  *time.Ticker
}

type Unit struct {
	KnowNodes map[string]*State
	Vault     mem.IVault
	LastNode  NodeLastInfo
}

type State struct {
	Public node.DataNode
	//Connection *grpc.ClientConn
	Client node.NodeCommunicationClient
}

//// ServerRequestInterceptor TODO move
//// ServerRequestInterceptor middleware на стороне сервера
//func (u *Unit) ServerRequestInterceptor(
//	ctx context.Context,
//	req interface{},
//	info *grpc.UnaryServerInfo,
//	handler grpc.UnaryHandler,
//) (response interface{}, err error) {
//	//todo ??
//
//	return handler(ctx, req)
//}

//func getClient(
//	srv string,
//	interceptor func(
//		context.Context,
//		string,
//		interface{},
//		interface{},
//		*grpc.ClientConn,
//		grpc.UnaryInvoker,
//		...grpc.CallOption) error, logger *zap.Logger) (*grpc.ClientConn, error) {
//
//	ticker := time.NewTicker(time.Millisecond)
//	timeout := time.After(5 * time.Second)
//	nSeconds := 1
//	conn, _ := grpc.Dial(
//		srv,
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//		grpc.WithUnaryInterceptor(interceptor),
//	)
//	for {
//		select {
//		case <-ticker.C:
//			ticker.Stop()
//			if conn.GetState() != connectivity.Ready {
//
//				logger.Info("Нет соединение с нодой:", zap.String("node", srv), zap.Duration("проверка подключения через", time.Duration(nSeconds)*time.Millisecond))
//				tmp := time.Duration(nSeconds) * time.Millisecond
//				//fmt.Println(tmp)
//				ticker = time.NewTicker(tmp)
//				nSeconds *= 2
//				continue
//			}
//			return conn, nil
//		case <-timeout:
//			return nil, errors.New("node: connection timeout")
//		}
//	}
//}

//// CreateReqKnownNodes Функция для формирования knownNodes
//func (u *Unit) CreateReqKnownNodes() *node.KnownNodes {
//	knownNodesMap := make(map[string]*node.DataNode)
//	for k, v := range u.KnowNodes {
//		knownNodesMap[k] = &node.DataNode{Ts: v.Public.Ts}
//	}
//	// добавление скоего адресса в мапу
//	knownNodesMap[u.Address] = &node.DataNode{
//		Ts: timestamppb.Now(),
//	}
//	knownNodesMessage := &node.KnownNodes{
//		Nodes: knownNodesMap,
//	}
//	return knownNodesMessage
//}
