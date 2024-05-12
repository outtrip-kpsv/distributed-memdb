package repo

import (
	cfg "team01/internal/config"
	"team01/internal/node/bl/model"
	"team01/internal/node/db"
	"time"
)

type INodeBL interface {
	GetUnit() *model.Unit
}

type nodeBl struct {
	core *model.Unit
}

func (n nodeBl) GetUnit() *model.Unit {
	return n.core
}

func NewNodeBL(dbRepo *db.DBRepo) INodeBL {
	return &nodeBl{core: newNode(dbRepo)}
}

func newNode(dbRepo *db.DBRepo) *model.Unit {
	res := model.Unit{
		Address:   cfg.GetAddress(),
		Vault:     dbRepo.Vault,
		KnowNodes: make(map[string]*model.State),
		LastNode: model.NodeLastInfo{
			Address: "",
			Ticker:  &time.Ticker{},
		},
	}
	return &res
}
