package node_rpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"sync"
	cfg "team01/internal/config"
	"team01/internal/node/io/grpc/util"
	"team01/internal/proto/node"
	"time"
)

func (n *NodeRpc) Ping(ctx context.Context, req *node.PingRequest) (*node.PingResult, error) {
	cfg.GetLogger().Info("method PING")
	return &node.PingResult{Res: true}, nil
}

func (n *NodeRpc) GetKnownNodes(ctx context.Context, req *node.KnownNodes) (*node.KnownNodes, error) {
	cfg.GetLogger().Info("method GetKnownNodes", zap.Reflect("nodes: ", req.Nodes))
	// TODO унести в бизнес логику
	// TODO context refactor
	ctxNew, _ := context.WithTimeout(ctx, 5*time.Second)

	var wg sync.WaitGroup

	ch := make(chan struct {
		address string
		conn    *grpc.ClientConn
		err     error
	})

	for k, v := range req.Nodes {
		if !n.BL.Node.NodeIsKnown(k) {
			//newNode = append(newNode, k)
			if k == cfg.GetAddress() {
				continue
			}
			wg.Add(1)
			cfg.GetLogger().Info("sss")
			go func(k string) {
				defer wg.Done()
				client, err := util.GetClient(ctxNew, k, n.MiddleWare.ClientRequestInterceptor)
				ch <- struct {
					address string
					conn    *grpc.ClientConn
					err     error
				}{k, client, err}
			}(k)
		} else {
			cfg.GetLogger().Info("++++")

			n.BL.Node.UpdTimePingNode(k, v.Ts)
			n.BL.Node.UpdLastNode()
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		if result.err != nil {
			cfg.GetLogger().Warn("Не удалось подключиться:", zap.Error(result.err))
			delete(n.BL.Node.GetUnit().KnowNodes, result.address)
			n.BL.Node.UpdLastNode()
			cfg.GetLogger().Info("22", zap.Reflect("map", n.BL.Node.GetUnit().KnowNodes))
			continue
		}
		n.BL.Node.AddNodeToKnown(result.address, node.NewNodeCommunicationClient(result.conn))
	}

	//for k, v := range req.Nodes {
	//	err := n.BL.ConnectTo(k, v.Timestamp)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//oldTime := time.Now()
	//for k, v := range req.Nodes {
	//	if k == u.Address {
	//		continue
	//	}
	//	//err := u.ConnectTo(k, v.Timestamp)
	//	//if err != nil {
	//	//  return err
	//	//}
	//	cfg.GetLogger().Info("find time for last node check " + k)
	//	if v.Ts.AsTime().Before(oldTime) {
	//		oldTime = v.Ts.AsTime()
	//		u.LastNode.Address = k
	//		u.Logger.Info("last node naw is " + k)
	//	} else {
	//		u.Logger.Info("ln old " + u.LastNode.Address)
	//		//fmt.Println()
	//	}
	//}
	//u.LastNode.Ticker.Stop()
	//tmp := time.Second*5 - (time.Now().Sub(oldTime))
	//cfg.GetLogger().Info("-------- " + tmp.String())
	////fmt.Println("t", tmp)
	//if tmp <= 0 {
	//	tmp = time.Millisecond
	//}
	//u.LastNode.Ticker = time.NewTicker(tmp)

	//nn := u.CreateReqKnownNodes()
	nn := n.BL.Node.GetKnowNode()
	return nn, nil
}
