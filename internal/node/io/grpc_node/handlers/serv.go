package handlers

import (
	"context"
	"google.golang.org/grpc"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/proto/node"
)

type NodeController struct {
	node.UnimplementedNodeCommunicationServer
	Bl          *bl.BL
	needConnect chan string
}

func NewNodeController(bl *bl.BL, conn chan string) *NodeController {
	return &NodeController{
		Bl:          bl,
		needConnect: conn,
	}
}

func (n *NodeController) Ping(ctx context.Context, req *node.PingRequest) (*node.PingResult, error) {
	return &node.PingResult{Res: true}, nil
}

func (n *NodeController) GetInfo(ctx context.Context, req *node.Info) (*node.Info, error) {
	newNodes := n.Bl.Node.UpdateKnowNode(req.Env)
	for _, newNode := range newNodes {
		if newNode == cfg.GetAddress() {
			continue
		}
		if int(req.Repl) != cfg.GetRepl() {
			cfg.GetLogger().Info("Разный коэффициент репликации, соединение не возможно")
			continue
		}
		n.needConnect <- newNode
	}
	res := n.Bl.Node.GetInfo()

	return res, nil
}

//func (n *NodeController) Set(ctx context.Context, setReq *node.SetRequest) (*node.NodeResp, error) {
//	result := n.Bl.Node.GetUnit().Vault.SetArtifact(setReq.Uuid, setReq.Val)
//	return &node.NodeResp{
//		Value:       "OK",
//		Code:        0,
//		CountResult: 1,
//		Result:      map[string]bool{cfg.GetAddress(): result},
//	}, nil
//}

func (n *NodeController) RequestP2P(ctx context.Context, in *node.CliReq, opts ...grpc.CallOption) (*node.NodeResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NodeController) Get(ctx context.Context, in *node.Request, opts ...grpc.CallOption) (*node.ArtResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NodeController) Set(ctx context.Context, in *node.Request, opts ...grpc.CallOption) (*node.ArtResp, error) {
	art, result := n.Bl.Node.GetUnit().Vault.SetArtifact(in.Uuid, in.Value, in.HashCluster)
	var respStatus node.RespStatus
	if result {
		respStatus = node.RespStatus_OK
	} else {
		respStatus = node.RespStatus_NO
	}
	return &node.ArtResp{
		Art: &art,
		Res: respStatus,
	}, nil
}

func (n *NodeController) Delete(ctx context.Context, in *node.Request, opts ...grpc.CallOption) (*node.ArtResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NodeController) Repl(ctx context.Context, in *node.Artefact, opts ...grpc.CallOption) (*node.PingResult, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NodeController) ProxyRequest(ctx context.Context, in *node.CliReq, opts ...grpc.CallOption) (*node.NodeResp, error) {
	//TODO implement me
	panic("implement me")
}

func returnCode(in bool) uint64 {
	if in {
		return 1
	} else {
		return 2
	}
}
