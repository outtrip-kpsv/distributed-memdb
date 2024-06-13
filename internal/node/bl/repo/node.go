package repo

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	cfg "team01/internal/config"
	"team01/internal/node/bl/model"
	"team01/internal/node/db"
	"team01/internal/proto/node"
	"time"
)

type INodeBL interface {
	GetUnit() *model.Unit //TODO remove
	GetInfo() *node.Info
	GetSizeCluster() int
	AddNodeToKnown(address string, client node.NodeCommunicationClient)
	NodeIsKnown(address string) bool
	UpdateKnowNode(ourNodes *node.KnownNodes) []string
	UpdTimePingNode(address string, ts *timestamppb.Timestamp)
	UpdLastNode()
	DeleteNode(address string)

	TickerLastNode() *time.Ticker
	GetLastClient() node.NodeCommunicationClient
}

type nodeBl struct {
	core *model.Unit
}

func NewNodeBL(dbRepo *db.DBRepo) INodeBL {
	return &nodeBl{core: newNode(dbRepo)}
}

func (n *nodeBl) GetInfo() *node.Info {
	return &node.Info{
		Address:   cfg.GetAddress(),
		Repl:      int32(cfg.GetRepl()),
		SizeVault: int32(n.core.Vault.GetSize()),
		Status:    n.core.Status,
		Env:       createReqKnownNodes(n.core.KnowNodes),
	}
}

func (n *nodeBl) GetSizeCluster() int {
	return n.core.ClusterSize
}

func (n *nodeBl) DeleteNode(address string) {
	cfg.GetLogger().Info("DEL " + address)

	delete(n.core.KnowNodes, address)
	n.core.ClusterSize--
	n.UpdLastNode()
	n.updStatus()

}

// GetUnit TODO remove
func (n *nodeBl) GetUnit() *model.Unit {
	return n.core
}

func (n *nodeBl) GetLastClient() node.NodeCommunicationClient {
	//TODO ---
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x2dec30b]

		goroutine 37 [running]:
		team01/internal/node/bl/repo.(*nodeBl).GetLastClient(0xc0001f9ee8?)
	*/
	return n.core.KnowNodes[n.core.LastNode.Address].Client
}

func (n *nodeBl) TickerLastNode() *time.Ticker {
	return n.core.LastNode.Ticker
}

func (n *nodeBl) AddNodeToKnown(address string, client node.NodeCommunicationClient) {
	n.core.KnowNodes[address] = &model.State{
		Public: node.DataNode{
			Ts: timestamppb.New(time.Now().Add(-time.Second * 5)),
		},
		Client: client,
	}
	n.core.ClusterSize++

	n.updStatus()
}

func (n *nodeBl) NodeIsKnown(address string) bool {
	if _, ok := n.core.KnowNodes[address]; !ok {
		return false
	}
	return true
}

func (n *nodeBl) UpdateKnowNode(ourNodes *node.KnownNodes) []string {
	var res []string
	for k, v := range ourNodes.Nodes {
		if k == cfg.GetAddress() {
			continue
		}
		if n.NodeIsKnown(k) {
			n.UpdTimePingNode(k, v.Ts)
		} else {
			res = append(res, k)
		}
	}
	return res
}

func (n *nodeBl) UpdTimePingNode(address string, ts *timestamppb.Timestamp) {
	stateTime := n.core.KnowNodes[address].Public.Ts.AsTime()
	if stateTime.Before(ts.AsTime()) {
		n.core.KnowNodes[address].Public.Ts = ts
		n.UpdLastNode()
	}
}

func (n *nodeBl) UpdLastNode() {

	switch len(n.core.KnowNodes) {
	case 0:
		n.TickerLastNode().Stop()
		return
	default:
		oldTime := time.Now()
		for k, v := range n.core.KnowNodes {
			if v.Public.Ts.AsTime().Before(oldTime) {
				oldTime = v.Public.Ts.AsTime()
				n.core.LastNode.Address = k
			}
		}
		n.core.LastNode.Ticker.Stop()
		tmp := time.Second*5 - (time.Now().Sub(oldTime))
		if tmp <= 0 {
			tmp = time.Millisecond
		}
		n.core.LastNode.Ticker = time.NewTicker(tmp)
	}
}

func newNode(dbRepo *db.DBRepo) *model.Unit {
	res := model.Unit{
		Vault:     dbRepo.Vault,
		KnowNodes: make(map[string]*model.State),
		LastNode: model.NodeLastInfo{
			Address: "",
			Ticker:  &time.Ticker{},
		},
		ClusterSize: 1,
	}
	return &res
}

// TODO добавить логики
func (n *nodeBl) updStatus() {
	if cfg.GetRepl() <= n.core.ClusterSize {
		n.core.Status = node.NodeStatus_LISTENER
	} else {
		n.core.Status = node.NodeStatus_UNKNOWN
	}
}

// createReqKnownNodes Функция для формирования knownNodes
func createReqKnownNodes(nodes map[string]*model.State) *node.KnownNodes {
	knownNodesMap := make(map[string]*node.DataNode)
	for k, v := range nodes {
		knownNodesMap[k] = &node.DataNode{Ts: v.Public.Ts}
	}
	// добавление своего адреса в мапу
	knownNodesMap[cfg.GetAddress()] = &node.DataNode{
		Ts: timestamppb.Now(),
	}

	return &node.KnownNodes{
		Nodes: knownNodesMap,
	}
}
