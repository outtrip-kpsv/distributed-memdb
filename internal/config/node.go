package config

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"log"
	"os"
)

type OptionsSrv struct {
	Host        string `short:"h" long:"host" description:"хост" default:"localhost" env:"HOST"`
	Port        string `short:"p" long:"port" description:"порт" default:"8765" env:"PORT"`
	Log         string `long:"logger-create" description:"logger-create output" default:"debug" env:"LOG"`
	ConnectHost string `long:"chost" description:"хост существующего экземпляра" default:"" env:"CONN_HOST"`
	ConnectPort string `long:"cport" description:"порт существующего экземпляра" default:"" env:"CONN_PORT"`
	Repl        int    `long:"repl" description:"коэффициент репликации" default:"2" env:"K_REPL"`
}

var globalConf = ConfSrv{}

type ConfSrv struct {
	AppName string
	Options OptionsSrv
	Logger  *zap.Logger
}

func init() {
	var conf ConfSrv
	var opts OptionsSrv
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}
	logger := initLogger(opts.Log)
	conf.Options = opts
	conf.Logger = logger

	globalConf.Options = opts
	globalConf.Logger = logger
}

func SetAppName(name string) {
	globalConf.AppName = name
}

func GetLogger() *zap.Logger {
	return globalConf.Logger
}

func GetAddress() string {
	return fmt.Sprintf("%s:%s", globalConf.Options.Host, globalConf.Options.Port)
}

func GetConnectAddress() (string, error) {
	if len(globalConf.Options.ConnectHost) == 0 || len(globalConf.Options.ConnectPort) == 0 {
		return "", errors.New("not connect")
	}
	return fmt.Sprintf("%s:%s", globalConf.Options.ConnectHost, globalConf.Options.ConnectPort), nil
}

func GetRepl() int {
	return globalConf.Options.Repl
}
