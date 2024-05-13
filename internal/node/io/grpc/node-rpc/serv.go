package node_rpc

import (
	"context"
	cfg "team01/internal/config"
	"team01/internal/proto/node"
)

func (n *NodeRpc) Ping(ctx context.Context, req *node.PingRequest) (*node.PingResult, error) {
	cfg.GetLogger().Info("method PING")
	return &node.PingResult{Res: true}, nil
}

func (n *NodeRpc) GetKnownNodes(ctx context.Context, req *node.KnownNodes) (*node.KnownNodes, error) {
	cfg.GetLogger().Info("method GetKnownNodes")
	// TODO унести в бизнес логику
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
