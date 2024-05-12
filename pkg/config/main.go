package config

type Iconfig interface {
	GetConnectAddress() (string, error)
}

//func Init(options Iconfig) interface{} {
//	cfg := Iconfig{}
//
//	return
//}
//
//func init(opt Iconfig)  {
//	return
//}
