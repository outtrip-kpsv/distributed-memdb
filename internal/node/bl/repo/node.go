package repo

import (
	"google.golang.org/protobuf/types/known/timestamppb"
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
}

type nodeBl struct {
	core *model.Unit
}

func (n nodeBl) AddNodeToKnown(address string, client node.NodeCommunicationClient) {
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

func (n nodeBl) NodeIsKnown(address string) bool {
	if _, ok := n.core.KnowNodes[address]; !ok {
		return false
	}
	return true
}

func (n nodeBl) GetKnowNode() *node.KnownNodes {
	return createReqKnownNodes(n.core.KnowNodes)
}

func (n nodeBl) GetUnit() *model.Unit {
	return n.core
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
	//knownNodesMap[u.Address] = &node.DataNode{
	//	Timestamp: timestamppb.Now(),
	//}

	return &node.KnownNodes{
		Nodes: knownNodesMap,
	}
}
