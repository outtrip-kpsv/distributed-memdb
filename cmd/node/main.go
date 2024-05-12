package main

import (
	cfg "team01/internal/config"
	"team01/internal/node/bl"
	"team01/internal/node/db"
)

func main() {
	cfg.SetAppName("Node")
	cfg.GetLogger().Info("start " + cfg.GetAddress())

	dbRepo := db.NewDBRepo()
	dbRepo.Vault.GetArtifact("w")

	blRepo := bl.NewBL(dbRepo)

	//fmt.Println(cfg.GetAddress())

}
