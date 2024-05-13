package mem

//
//type dataNode struct {
//	db map[string]*State
//}
//
//func NewVault() db.IVault {
//	res := vault{db: make(map[string]node.Artefact)}
//	return &res
//}
//
//func (v *vault) GetArtifact(uuid string) *node.Artefact {
//	val, ok := v.db[uuid]
//	if ok {
//		return &val
//	}
//	return nil
//}
//
//type State struct {
//	Public     node.DataNode
//	Connection *grpc.ClientConn
//	Client     node.NodeCommunicationClient
//}
