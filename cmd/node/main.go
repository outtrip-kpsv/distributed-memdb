package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/db"
	"team01/internal/node/io/grpc_node"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, syscall.SIGINT, syscall.SIGTERM)
	finished := make(chan bool, 1)

	cfg.SetAppName("Node")

	dbRepo := db.NewDBRepo()
	blRepo := bl.NewBL(dbRepo)
	srv := grpc_node.NewGrpcNode(blRepo, finished)
	srv.Run(ctx)

	go func() {
		<-interruptCh
		cfg.GetLogger().Info("Received interrupt signal. Cleaning up...")
		cancel()
	}()
	<-finished

}
