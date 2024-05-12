package mem

import (
  "team01/internal/node/db"
  "team01/internal/proto/node"
)

type data struct {
  db map[string]node.DataNode
}

func NewData() db.IData {
  res := data{db: make(map[string]node.DataNode)}
  return &res
}
