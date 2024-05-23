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
	GetKnowNode() *node.KnownNodes
	AddNodeToKnown(address string, client node.NodeCommunicationClient)
	NodeIsKnown(address string) bool
	UpdateKnowNode(ourNodes *node.KnownNodes) []string
	UpdTimePingNode(address string, ts *timestamppb.Timestamp)
	UpdLastNode()
}

type nodeBl struct {
	core *model.Unit
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
	return res
}

func (n *nodeBl) AddNodeToKnown(address string, client node.NodeCommunicationClient) {
	//var st = model.State{
	//	Public:     node.DataNode{},
	//}
	n.core.KnowNodes[address] = &model.State{
		Public: node.DataNode{Ts: timestamppb.Now()},
		Client: client,
	}
	//n.core.KnowNodes[address].Public.Ts = timestamppb.Now()
	//n.core.KnowNodes[address].Client = client

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

func (n *nodeBl) GetUnit() *model.Unit {
	return n.core
}

func (n *nodeBl) UpdTimePingNode(address string, ts *timestamppb.Timestamp) {
	//if ts == nil {
	//	ts = timestamppb.Now()
	//}

	stateTime := n.core.KnowNodes[address].Public.Ts.AsTime()
	if stateTime.Before(ts.AsTime()) {
		n.core.KnowNodes[address].Public.Ts = ts
		//cfg.GetLogger().Info("time upd", zap.String("node", address), zap.Reflect("time", ts.AsTime()))
	}
}

func (n *nodeBl) UpdLastNode() {
	oldTime := time.Now()
	for k, v := range n.core.KnowNodes {

		cfg.GetLogger().Info("find time for last node check " + k)
		if v.Public.Ts.AsTime().Before(oldTime) {
			oldTime = v.Public.Ts.AsTime()
			n.core.LastNode.Address = k
			//cfg.GetLogger().Info("last node naw is " + k)
		} else {
			//cfg.GetLogger().Info("ln old " + n.core.LastNode.Address)
			//fmt.Println()
		}
	}
	n.core.LastNode.Ticker.Stop()
	tmp := time.Second*5 - (time.Now().Sub(oldTime))
	//cfg.GetLogger().Info("-------- " + tmp.String())
	//fmt.Println("t", tmp)
	if tmp <= 0 {
		tmp = time.Millisecond
	}
	n.core.LastNode.Ticker = time.NewTicker(tmp)
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
