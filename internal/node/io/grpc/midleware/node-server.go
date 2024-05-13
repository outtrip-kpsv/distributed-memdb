package midleware

import (
	"context"
	"google.golang.org/grpc"
)

// ServerRequestInterceptor middleware на стороне сервера
func ServerRequestInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (response interface{}, err error) {
	//todo ??

	return handler(ctx, req)
}
