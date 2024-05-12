package db

import (
	"team01/internal/node/db/mem"
	"team01/internal/proto/node"
)

type DBRepo struct {
	Vault IVault
}

type IVault interface {
	GetArtifact(uuid string) *node.Artefact
}

func NewDBRepo() *DBRepo {
	return &DBRepo{
		Vault: mem.NewVault(),
	}
}

func (D DBRepo) GetArtifact(uuid string) *node.Artefact {
	//TODO implement me
	panic("implement me")
}
