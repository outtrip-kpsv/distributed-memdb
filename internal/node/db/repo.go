package db

import (
  "team01/internal/node/db/mem"
  "team01/internal/proto/node"
)

type DBRepo struct {
  Data  IData
  Vault IVault
}

type IVault interface {
  GetArtifact(uuid string) *node.Artefact
}

type IData interface {
}

func NewDBRepo() *DBRepo {
  return &DBRepo{
    Data:  mem.NewData(),
    Vault: mem.NewVault(),
  }
}
