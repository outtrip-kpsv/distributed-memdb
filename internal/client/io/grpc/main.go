package grpc_cli

import (
	"context"
	"log"
	"team01/internal/client/app"
	cfg "team01/internal/config"
	"team01/internal/proto/node"
	"team01/internal/util"
	"team01/internal/util/journal"
	"time"
)

type Cli struct {
	KnownNodes map[string]interface{}
	Bl         *app.App
	journal    *journal.Journal
}

func NewConnect(app *app.App) *Cli {
	res := Cli{
		KnownNodes: make(map[string]interface{}),
		Bl:         app,
		journal:    journal.NewJournal(),
	}
	res.KnownNodes[cfg.GetAddressNode()] = nil
	return &res
}

func (c *Cli) Run() {
	c.Bl.CLog.LogWorker()
	go c.Connect("")
	if err := c.Bl.Start(); err != nil {
		panic(err)
	}
}

func (c *Cli) Connect(errAdr string) {
	for knownNode := range c.KnownNodes {
		if knownNode == errAdr {
			continue
		}

		ctxN, _ := context.WithTimeout(context.Background(), time.Second*5)
		connect, err := util.GetClient(ctxN, knownNode, c.Bl.CLog.Log, c.ClientRequestInterceptor)

		if err != nil {
			continue
		}
		c.Bl.GprcConn = connect
		c.Bl.Connect(node.NewClientCommunicationClient(connect))
		return
	}
	log.Println("все известные адреса не доступны")
	c.Bl.Tview.Stop()
}

func (c *Cli) UpdListConnect(env *node.KnownNodes) {
	res := make(map[string]interface{})
	c.KnownNodes = res
	for k := range env.Nodes {
		c.KnownNodes[k] = nil
	}
}
