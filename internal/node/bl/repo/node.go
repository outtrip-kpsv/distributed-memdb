package repo

import (
	"team01/internal/node/bl/model"
	"team01/internal/node/db"
	"team01/internal/proto/node"
	"time"
)

type INodeBL interface {
	GetUnit() *model.Unit //TODO remove
	GetKnowNode() *node.KnownNodes
}

type nodeBl struct {
	core *model.Unit
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
