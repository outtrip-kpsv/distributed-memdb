package middleware

import (
	"context"
	"google.golang.org/grpc"
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

	return handler(ctx, req)
}
