package main

import (
	"fmt"
	cfg "team01/internal/config"
)

func main() {

	settings := cfg.Get()
	fmt.Println(settings.GetAddress())
}
