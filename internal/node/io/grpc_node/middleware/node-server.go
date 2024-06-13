package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	cfg "team01/internal/config"
)

// ServerRequestInterceptor middleware на стороне сервера
func (m *middleware) ServerRequestInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (response interface{}, err error) {
	// Получение метаданных из контекста
	md, ok := metadata.FromIncomingContext(ctx)
	var from string
	if ok {
		from = md["from"][0]

	} else {
		from = "none"
	}
	if from != "client" && m.bl.Node.NodeIsKnown(from) {
		m.bl.Node.UpdTimePingNode(from, timestamppb.Now())
	}

	cfg.GetLogger().Info("method -> " + info.FullMethod + " from: " + from)

	return handler(ctx, req)
}
