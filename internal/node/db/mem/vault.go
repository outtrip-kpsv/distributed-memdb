package mem

import (
  "team01/internal/node/db"
  "team01/internal/proto/node"
)

type vault struct {
  db map[string]node.Artefact
}

func NewVault() db.IVault {
  res := vault{db: make(map[string]node.Artefact)}
  return &res
}

func (v *vault) GetArtifact(uuid string) *node.Artefact {
  val, ok := v.db[uuid]
  if ok {
    return &val
  }
  return nil
}
