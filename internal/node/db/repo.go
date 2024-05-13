package db

import (
	"team01/internal/node/db/mem"
)

type DBRepo struct {
	Vault    mem.IVault
	DataNode IDataNode
}

type IDataNode interface {
}

func NewDBRepo() *DBRepo {
	return &DBRepo{
		Vault: mem.NewVault(),
	}
}
