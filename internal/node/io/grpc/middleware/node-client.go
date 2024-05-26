package middleware

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	cfg "team01/internal/config"
	"team01/internal/proto/node"
)

//todo все гетноде переписать

// ClientRequestInterceptor клентский мидлваре выполняемый перед запросом
func (m *middleware) ClientRequestInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {

	if cc.GetState() != connectivity.Ready {
		cfg.GetLogger().Info("Нет соединение с нодой", zap.String("address", cc.Target()))
		m.bl.Node.DeleteNode(cc.Target())
		return errors.New("Server connection is not ready")
	}

	resultErr := invoker(ctx, method, req, reply, cc, opts...)
	// после запроса
	if resultErr == nil && method != "/NodeCommunication/GetKnownNodes" {
		ourNode := &node.KnownNodes{}
		thisNode := m.bl.Node.GetKnowNode()
		err := cc.Invoke(ctx, "/NodeCommunication/GetKnownNodes", thisNode, ourNode, opts...)

		if err != nil {
			cfg.GetLogger().Info("ERR", zap.Error(err))
			return err
		}
		newNode := m.bl.Node.UpdateKnowNode(ourNode)
		fmt.Println(newNode)

	}

	return resultErr
}
