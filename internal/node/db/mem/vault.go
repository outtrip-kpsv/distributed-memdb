package mem

import (
	"team01/internal/proto/node"
)

type IVault interface {
	GetArtifact(uuid string) *node.Artefact
	SetArtifact(uuid, val, hashCluster string) (node.Artefact, bool)
	GetSize() int
}

type vault struct {
	db   map[string]node.Artefact
	size int
}

func NewVault() IVault {
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

// SetArtifact возвращает false если значение было переписано, и true есди новое значение
func (v *vault) SetArtifact(uuid, val, hashCluster string) (node.Artefact, bool) {
	_, ok := v.db[uuid]
	if !ok {
		v.size++
	}
	v.db[uuid] = node.Artefact{Name: val, HashCluster: hashCluster}

	return node.Artefact{Name: val, HashCluster: hashCluster}, !ok
}

func (v *vault) GetSize() int {
	return v.size
}

func add() {

}
