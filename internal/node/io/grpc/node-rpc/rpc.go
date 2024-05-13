package node_rpc

import (
	"team01/internal/node/bl"
	"team01/internal/proto/node"
)

type NodeRpc struct {
	node.UnimplementedNodeCommunicationServer
	BL *bl.BL
}

func CreateNode(bl *bl.BL) *NodeRpc {
	return &NodeRpc{
		BL: bl,
	}
}
