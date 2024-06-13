package grpc_cli

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"
	"strings"
	"team01/internal/proto/node"
)

// ClientRequestInterceptor клиентский мидлваре
func (c *Cli) ClientRequestInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	cc.Target()
	if cc.GetState() != connectivity.Ready {
		c.Bl.CLog.Log <- "соединение с узлом " + cc.Target() + " потерянно"
		c.Connect(cc.Target())
		cc = c.Bl.GprcConn
	}

	// Создание метаданных
	md := metadata.Pairs("from", "client")
	ctx = metadata.NewOutgoingContext(ctx, md)

	if cc.GetState() == connectivity.Ready && method != "/ClientCommunication/GetInfoNode" {
		ourNodeInfo := &node.Info{}
		err := cc.Invoke(ctx, "/ClientCommunication/GetInfoNode", &node.PingRequest{}, ourNodeInfo, opts...)

		if err != nil {
			c.Bl.CLog.Log <- err.Error()
			return err
		}
		c.Bl.CInfo.UpdateInfo(ourNodeInfo)
		c.UpdListConnect(ourNodeInfo.Env)
		if c.Bl.CInfo.Info.Status == node.NodeStatus_UNKNOWN {
			return errors.New("кластер не доступен")
		}
	}

	if method == "/ClientCommunication/Request" && c.Bl.CInfo.Info.Status == node.NodeStatus_LISTENER {

		// хэш функция

		/*
				TODO при отсутствии записи в мапе проходить по всем нодам чтобы исключить повторные записи
				хэш кластера отсортированный список нод
				жэш ууид
				------
				заморозить значение / удалить и записать в соответствии с новой схемой


			hash uuid/ hash cluster
		*/
		request := req.(*node.CliReq)
		response := node.NodeResp{}

		addr := c.journal.HashRequest(request.Req.Uuid, c.Bl.CInfo.Info.Env)
		addr = addr[:c.Bl.CInfo.Info.Repl]

		c.Bl.CLog.Log <- strings.Join(addr, "=")

		request.Req.HashCluster = c.journal.GetClusterHash(request.Req.Uuid)

		resultErr := invoker(ctx, method, request, &response, cc, opts...)
		// TODO !!! повторять запросы до достижения результата

		if resultErr != nil {
			c.Bl.CLog.Log <- resultErr.Error()

		}
		return resultErr
	}

	resultErr := invoker(ctx, method, req, reply, cc, opts...)
	if resultErr != nil {
		c.Bl.CLog.Log <- resultErr.Error()

	}
	return resultErr
}
