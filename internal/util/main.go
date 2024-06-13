package util

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"regexp"
	"sort"
	"strings"
	"team01/internal/proto/node"
	"time"
)

// GetClient для получения соединения с сервером gRPC
func GetClient(
	ctx context.Context,
	srv string,
	log chan string,
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

	for {
		select {
		case <-ticker.C:
			if conn.GetState() != connectivity.Ready {
				log <- fmt.Sprintf("Нет соединения с нодой: %s, следущая попытка через: %s", srv, time.Duration(nSeconds).String())
				tmp := time.Duration(nSeconds) * time.Millisecond

				ticker = time.NewTicker(tmp)

				nSeconds *= 2
				continue
			}
			log <- fmt.Sprintf("Успешное соединение с %s", srv)
			return conn, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// Timer показывает время прошедшее с момента последнего пинга переданной ноды
func Timer(node *node.DataNode) string {
	return timestamppb.Now().AsTime().Sub(node.Ts.AsTime()).String()
}

// KnowNodesToString преобразование списка известных нод к слайсу строк
func KnowNodesToString(nodes *node.KnownNodes) []string {
	var knowNodes []string
	for s := range nodes.Nodes {
		knowNodes = append(knowNodes, s)
	}
	sort.Strings(knowNodes)
	return knowNodes
}

// RequestSplit очистка запроса от лишних пробелов табов, и представление в виде слайса строк
func RequestSplit(req string) []string {
	r := regexp.MustCompile("\\s+")
	req = r.ReplaceAllString(strings.TrimSpace(req), " ")
	return strings.Split(req, " ")
}
