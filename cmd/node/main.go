package main

import (
  cfg "team01/internal/config"
  "team01/internal/node/db"
)

func main() {
  cfg.SetAppName("Node")
  cfg.GetLogger().Info("start " + cfg.GetAddress())

  dbRepo := db.NewDBRepo()
  dbRepo.Vault.SetArtifact()

  //fmt.Println(cfg.GetAddress())

}
