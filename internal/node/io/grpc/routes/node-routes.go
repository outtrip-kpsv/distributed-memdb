package routes

import (
	"team01/internal/node/bl"
	"team01/internal/node/io/grpc/middleware"
)

type router struct {
	bl          *bl.BL
	middlewares middleware.IMiddleWare
}

func InitNodeRoutes(bl *bl.BL) {

}
