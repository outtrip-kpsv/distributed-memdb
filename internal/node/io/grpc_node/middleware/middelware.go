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

	ServerRequestInterceptor(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (response interface{}, err error)
}

type middleware struct {
	bl       *bl.BL
	needConn chan string
}

func NewMiddlewares(bl *bl.BL, needConnect chan string) IMiddleWare {
	return &middleware{bl: bl, needConn: needConnect}
}
