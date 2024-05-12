package routes

import (
	"team01/internal/node/bl"
	"team01/internal/node/io/grpc/midleware"
)

type router struct {
	bl          *bl.BL
	middlewares midleware.IMiddleWare
}

func InitNodeRoutes(bl *bl.BL) {

}
