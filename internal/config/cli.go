package config

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

var cliConf = OptionsCli{}

type OptionsCli struct {
	Host string `short:"H" long:"host" description:"хост" default:"localhost" env:"HOST"`
	Port string `short:"P" long:"port" description:"порт" default:"8765" env:"PORT"`
}

func InitCli() {
	parser := flags.NewParser(&cliConf, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}
}

func GetAddressNode() string {
	return fmt.Sprintf("%s:%s", globalConf.Options.Host, globalConf.Options.Port)
}
