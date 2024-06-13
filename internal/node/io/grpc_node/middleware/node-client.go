package middleware

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"
	cfg "team01/internal/config"
	"team01/internal/proto/node"
)

//todo все гетноде переписать

// ClientRequestInterceptor клиентский мидлваре
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
		err := cc.Close()
		if err != nil {
			return err
		}
		cfg.GetLogger().Info("3")
		m.bl.Node.DeleteNode(cc.Target())
		m.bl.Node.UpdLastNode()
		return errors.New("отсутствует соединение с нодой")
	}
	// Создание метаданных
	md := metadata.Pairs("from", cfg.GetAddress())
	ctx = metadata.NewOutgoingContext(ctx, md)

	if method != "/NodeCommunication/GetInfo" {
		ourNodeInfo := &node.Info{}
		thisNodeInfo := m.bl.Node.GetInfo()
		err := cc.Invoke(ctx, "/NodeCommunication/GetInfo", thisNodeInfo, ourNodeInfo, opts...)

		if err != nil {
			cfg.GetLogger().Info("ERR", zap.Error(err))
			return err
		}
		if int(ourNodeInfo.Repl) != cfg.GetRepl() {
			m.bl.Node.DeleteNode(ourNodeInfo.Address)
			return errors.New("разный коэффициент репликации")
		}
		_ = m.bl.Node.UpdateKnowNode(ourNodeInfo.Env)

	}

	resultErr := invoker(ctx, method, req, reply, cc, opts...)

	return resultErr
}
