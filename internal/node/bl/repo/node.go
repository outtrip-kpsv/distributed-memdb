package repo

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	cfg "team01/internal/config"
	"team01/internal/node/bl/model"
	"team01/internal/node/db"
	"team01/internal/proto/node"
	"time"
)

type INodeBL interface {
	GetUnit() *model.Unit //TODO remove
	GetKnowNode() *node.KnownNodes
	AddNodeToKnown(address string, client node.NodeCommunicationClient)
	NodeIsKnown(address string) bool
	UpdateKnowNode(ourNodes *node.KnownNodes) []string
	UpdTimePingNode(address string, ts *timestamppb.Timestamp)
	UpdLastNode()
	DeleteNode(address string)

	TickerLastNode() *time.Ticker
	GetLastClient() node.NodeCommunicationClient
	ConnectToNodes(addresses []string)
}

type nodeBl struct {
	core *model.Unit
}

func (n *nodeBl) DeleteNode(address string) {
	delete(n.core.KnowNodes, address)
	n.UpdLastNode()
}

func (n *nodeBl) ConnectToNodes(addresses []string) {
	for _, address := range addresses {
		fmt.Println(address)
	}
}

// GetUnit TODO remove
func (n *nodeBl) GetUnit() *model.Unit {
	return n.core
}

func (n *nodeBl) GetLastClient() node.NodeCommunicationClient {
	return n.core.KnowNodes[n.core.LastNode.Address].Client
}

func (n *nodeBl) TickerLastNode() *time.Ticker {
	return n.core.LastNode.Ticker
}

func (n *nodeBl) AddNodeToKnown(address string, client node.NodeCommunicationClient) {
	n.core.KnowNodes[address] = &model.State{
		Public: node.DataNode{Ts: timestamppb.Now()},
		Client: client,
	}
}

func (n *nodeBl) NodeIsKnown(address string) bool {
	if _, ok := n.core.KnowNodes[address]; !ok {
		return false
	}
	return true
}

func (n *nodeBl) GetKnowNode() *node.KnownNodes {
	return createReqKnownNodes(n.core.KnowNodes)
}

func (n *nodeBl) UpdateKnowNode(ourNodes *node.KnownNodes) []string {
	var res []string
	for k, v := range ourNodes.Nodes {
		if n.NodeIsKnown(k) {
			n.UpdTimePingNode(k, v.Ts)
		} else {
			res = append(res, k)
		}
	}
	n.UpdLastNode()
	return res
}

func (n *nodeBl) UpdTimePingNode(address string, ts *timestamppb.Timestamp) {
	stateTime := n.core.KnowNodes[address].Public.Ts.AsTime()
	if stateTime.Before(ts.AsTime()) {
		n.core.KnowNodes[address].Public.Ts = ts
	}
}

func (n *nodeBl) UpdLastNode() {
	switch len(n.core.KnowNodes) {
	case 0:
		return
	case 1:
		for k := range n.core.KnowNodes {
			n.core.LastNode.Address = k
		}
		n.core.LastNode.Ticker = time.NewTicker(time.Second * 5)
	default:
		oldTime := time.Now()
		for k, v := range n.core.KnowNodes {
			cfg.GetLogger().Info("find time for last node check " + k)
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
	cfg.GetLogger().Info("Upd last node")
}

func NewNodeBL(dbRepo *db.DBRepo) INodeBL {
	return &nodeBl{core: newNode(dbRepo)}
}

func newNode(dbRepo *db.DBRepo) *model.Unit {
	res := model.Unit{
		Vault:     dbRepo.Vault,
		KnowNodes: make(map[string]*model.State),
		LastNode: model.NodeLastInfo{
			Address: "",
			Ticker:  &time.Ticker{},
		},
	}
	return &res
}

// createReqKnownNodes Функция для формирования knownNodes
func createReqKnownNodes(nodes map[string]*model.State) *node.KnownNodes {
	knownNodesMap := make(map[string]*node.DataNode)
	for k, v := range nodes {
		knownNodesMap[k] = &node.DataNode{Ts: v.Public.Ts}
	}
	// добавление скоего адресса в мапу
	knownNodesMap[cfg.GetAddress()] = &node.DataNode{
		Ts: timestamppb.Now(),
	}

	return &node.KnownNodes{
		Nodes: knownNodesMap,
	}
}
