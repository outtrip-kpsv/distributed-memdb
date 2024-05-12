package bl

import (
	"team01/internal/node/bl/repo"
	"team01/internal/node/db"
)

type BL struct {
	Node repo.INodeBL
}

func NewBL(dbRepo *db.DBRepo) *BL {
	return &BL{Node: repo.NewNodeBL(dbRepo)}
}
