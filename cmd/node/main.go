package main

import (
	cfg "team01/internal/config"
)

func main() {
	cfg.SetAppName("Node")
	cfg.GetLogger().Info("start " + cfg.GetAddress())
	//fmt.Println(cfg.GetAddress())
}
