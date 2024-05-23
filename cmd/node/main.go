package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/db"
	"team01/internal/node/io/grpc"
)

// TODO add gracefully shutdown
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//ctx := context.Background()
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, syscall.SIGINT, syscall.SIGTERM)
	finished := make(chan bool, 1)
	//defer close(finished)
	//defer close(interruptCh)

	cfg.SetAppName("Node")
	//
	dbRepo := db.NewDBRepo()
	dbRepo.Vault.GetArtifact("w")
	//
	blRepo := bl.NewBL(dbRepo)
	srv := grpc.NewGrpcNode(blRepo, finished)
	srv.Run(ctx)
	fmt.Println("------")
	go func() {
		<-interruptCh
		fmt.Println("Received interrupt signal. Cleaning up...")
		cancel()
	}()
	<-finished

}
