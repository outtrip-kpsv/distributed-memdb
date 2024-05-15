package middleware

import (
	"context"
	"google.golang.org/grpc"
	"team01/internal/node/bl"
)

type IMiddleWare interface {
	ClientRequestInterceptor(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
}

type middleware struct {
	bl *bl.BL
}

func NewMiddlewares(bl *bl.BL) IMiddleWare {
	return &middleware{bl: bl}
}
