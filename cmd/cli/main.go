package main

import (
	clientApp "team01/internal/client/app"
	"team01/internal/client/io/grpc"
)

func main() {

	app := clientApp.CreateApp()

	cli := grpc_cli.NewConnect(app)

	cli.Run()

}
