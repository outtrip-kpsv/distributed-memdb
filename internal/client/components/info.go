package components

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
	"team01/internal/proto/node"
	"team01/internal/util"
)

type InfoComponent struct {
	Container *tview.Flex
	Component *tview.TextView
	Info      *node.Info
}

func (c *InfoComponent) UpdateInfo(inf *node.Info) {
	c.Info = inf

	var infoText string
	switch c.Info.Status {
	case node.NodeStatus_LISTENER:
		infoText = fmt.Sprintf("Кластер без управляющей ноды")
	case node.NodeStatus_CANDIDATE, node.NodeStatus_CLIENT, node.NodeStatus_FOLLOWER:
		infoText = fmt.Sprintf("Кластер c управляющей нодой")
	case node.NodeStatus_UNKNOWN:
		infoText = fmt.Sprintf("Кластер не доступен")
	default:
		infoText = fmt.Sprintf("Статус")
	}

	//var knowNodes []string
	//for s := range c.Info.Env.Nodes {
	//	knowNodes = append(knowNodes, s+"\n                    ")
	//}

	text := fmt.Sprintf(
		`>> [blue] %s[white]
[yellow]адрес подключения:[white]	%s
[yellow]      статус узла:[white]	%s

[yellow] узлов в кластере:[white]	%d
[yellow]   известные ноды:[white]	%s
[yellow]   известные ноды:[white]`,

		infoText,
		c.Info.Address,
		c.Info.Status,
		len(c.Info.Env.Nodes),
		strings.Join(util.KnowNodesToString(c.Info.Env)[:], "\n                    "),
	)

	c.Component.SetText(text)
}

func CreateInfo() *InfoComponent {
	cp := tview.NewTextView().SetDynamicColors(true)

	ct := tview.NewFlex()
	ct.SetDirection(tview.FlexRow).SetBorder(true).SetTitle("Info")

	ct.AddItem(cp, 0, 1, false)
	res := InfoComponent{
		Container: ct,
		Component: cp,
	}
	return &res
}
