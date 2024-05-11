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

func Get() ConfSrv {
	return globalConf
}

func (c *ConfSrv) GetAddress() string {
	return fmt.Sprintf("%s:%s", c.Options.Host, c.Options.Port)
}

func (c *ConfSrv) GetConnectAddress() (string, error) {
	if len(c.Options.ConnectHost) == 0 || len(c.Options.ConnectPort) == 0 {
		return "", errors.New("not connect")
	}
	return fmt.Sprintf("%s:%s", c.Options.ConnectHost, c.Options.ConnectPort), nil
}
