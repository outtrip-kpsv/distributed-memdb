package node_rpc

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"sync"
	cfg "team01/internal/config"
	"team01/internal/node/io/grpc/util"
	"team01/internal/proto/node"
	"time"
)

func (n *NodeRpc) Ping(ctx context.Context, req *node.PingRequest) (*node.PingResult, error) {
	cfg.GetLogger().Info("method PING from: ")
	fmt.Println("context: ", ctx.Value("from"))
	return &node.PingResult{Res: true}, nil
}

func (n *NodeRpc) GetKnownNodes(ctx context.Context, req *node.KnownNodes) (*node.KnownNodes, error) {
	cfg.GetLogger().Info("method GetKnownNodes", zap.Reflect("nodes: ", req.Nodes))
	// TODO унести в бизнес логику
	// TODO context refactor

	newNodes := n.BL.Node.UpdateKnowNode(req)

	var wg sync.WaitGroup

	//ch := make(chan struct {
	//	address string
	//	conn    *grpc.ClientConn
	//	err     error
	//})

	for _, newNode := range newNodes {
		if newNode == cfg.GetAddress() {
			continue
		}
		wg.Add(1)
		//cfg.GetLogger().Info("sss")
		go func(k string) {
			defer wg.Done()
			//TODO ADD TIMEOUT
			//ctxNew, _ := context.WithTimeout(ctx, 5*time.Second)
			//defer cancel()
			fmt.Println(ctx)
			err := n.ConnectTo(ctx, k)
			if err != nil {
				cfg.GetLogger().Info(err.Error())
			}
			//client, err := util.GetClient(ctxNew, k, n.MiddleWare.ClientRequestInterceptor)
			//ch <- struct {
			//	address string
			//	conn    *grpc.ClientConn
			//	err     error
			//}{k, client, err}

		}(newNode)
	}

	go func() {
		wg.Wait()
		//close(ch)
	}()

	//for result := range ch {
	//	if result.err != nil {
	//		cfg.GetLogger().Warn("Не удалось подключиться:", zap.Error(result.err))
	//		delete(n.BL.Node.GetUnit().KnowNodes, result.address)
	//
	//		//cfg.GetLogger().Info("22", zap.Reflect("map", n.BL.Node.GetUnit().KnowNodes))
	//		continue
	//	}
	//	n.BL.Node.AddNodeToKnown(result.address, node.NewNodeCommunicationClient(result.conn))
	//	cfg.GetLogger().Info("OK f gkn", zap.String("Connect to", result.address))
	//
	//}
	nn := n.BL.Node.GetKnowNode()
	return nn, nil
}

func (s *NodeRpc) ConnectTo(ctx context.Context, address string) error {
	if cfg.GetAddress() == address {
		return nil
	}

	if ok := s.BL.Node.NodeIsKnown(address); !ok {
		//TODO named logger, remove ctx background !!! <<<<<
		ctxN, _ := context.WithTimeout(context.Background(), time.Second*5)
		conn, err := util.GetClient(ctxN, address, s.MiddleWare.ClientRequestInterceptor)
		if err != nil {
			return err
		}
		s.BL.Node.AddNodeToKnown(address, node.NewNodeCommunicationClient(conn))
		cfg.GetLogger().Info("OK", zap.String("Connect to", address))
	}
	return nil
}
