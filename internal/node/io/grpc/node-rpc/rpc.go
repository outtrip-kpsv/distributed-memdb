package node_rpc

import (
	"team01/internal/node/bl"
	m "team01/internal/node/io/grpc/middleware"
	"team01/internal/proto/node"
)

type NodeRpc struct {
	node.UnimplementedNodeCommunicationServer
	MiddleWare m.IMiddleWare
	BL         *bl.BL
}

func CreateNode(bl *bl.BL) *NodeRpc {
	return &NodeRpc{
		BL:         bl,
		MiddleWare: m.NewMiddlewares(bl),
	}
}
