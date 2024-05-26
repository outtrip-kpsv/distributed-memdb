package util

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	cfg "team01/internal/config"
	"time"
)

// GetClient для получения соединения с сервером gRPC
func GetClient(
	ctx context.Context,
	srv string,
	interceptor func(
		context.Context,
		string,
		interface{},
		interface{},
		*grpc.ClientConn,
		grpc.UnaryInvoker,
		...grpc.CallOption) error) (*grpc.ClientConn, error) {

	ticker := time.NewTicker(time.Millisecond)

	defer ticker.Stop()

	nSeconds := 1

	conn, _ := grpc.Dial(
		srv,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor),
	)
	fmt.Println(ctx)

	for {
		select {
		case <-ticker.C:
			if conn.GetState() != connectivity.Ready {
				cfg.GetLogger().Info("Нет соединения с нодой:", zap.String("node", srv), zap.Duration("проверка подключения через", time.Duration(nSeconds)*time.Millisecond))
				tmp := time.Duration(nSeconds) * time.Millisecond
				ticker = time.NewTicker(tmp)
				nSeconds *= 2
				continue
			}
			cfg.GetLogger().Info("CONNECT", zap.String("address: ", srv))
			return conn, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
