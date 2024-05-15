package main

import (
	"fmt"
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/db"
	"team01/internal/node/io/grpc"
)

// TODO add gracefully shutdown
func main() {
	cfg.SetAppName("Node")

	dbRepo := db.NewDBRepo()
	dbRepo.Vault.GetArtifact("w")

	blRepo := bl.NewBL(dbRepo)
	finished := make(chan bool)
	srv := grpc.NewGrpcNode(blRepo, finished)
	srv.Run()
	fmt.Println("------")
	<-finished

}
