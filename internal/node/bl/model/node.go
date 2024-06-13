package model

import (
	"team01/internal/node/db/mem"
	"team01/internal/proto/node"
	"time"
)

type Unit struct {
	KnowNodes   map[string]*State
	ClusterSize int
	Status      node.NodeStatus

	Vault    mem.IVault
	LastNode NodeLastInfo
}

type State struct {
	//Connection *grpc.ClientConn
	Client node.NodeCommunicationClient
	Public node.DataNode
}

type NodeLastInfo struct {
	Address string
	Ticker  *time.Ticker
}
