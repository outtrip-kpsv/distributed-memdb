package middleware

import (
	"context"
	"google.golang.org/grpc"
	cfg "team01/internal/config"
)

// ServerRequestInterceptor middleware на стороне сервера
func (m *middleware) ServerRequestInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (response interface{}, err error) {
	//todo ??

	m.bl.Node.UpdLastNode()
	cfg.GetLogger().Info("*******************************")

	return handler(ctx, req)
}
