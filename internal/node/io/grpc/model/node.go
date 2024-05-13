package model

import (
	"team01/internal/proto/node"
)

type NodeCore struct {
	node.UnimplementedNodeCommunicationServer
}
