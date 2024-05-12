package midleware

import "team01/internal/node/bl"

type IMiddleWare interface {
}

type midelwares struct {
	bl *bl.BL
}

func NewMidlewares(bl *bl.BL) IMiddleWare {
	return &midelwares{bl: bl}
}
