package app

import (
	"context"
	"github.com/rivo/tview"
	"google.golang.org/grpc"
	"team01/internal/client/components"
	"team01/internal/proto/node"
)

type App struct {
	Tview         *tview.Application
	RootContainer *tview.Flex
	viewContainer *tview.Flex
	Cli           node.ClientCommunicationClient
	GprcConn      *grpc.ClientConn
	cHistory      *components.HistoryComponent
	cRequest      *components.RequestComponent
	CInfo         *components.InfoComponent
	CLog          *components.LogComponent
}

// CreateApp Create application and initialize components
func CreateApp() *App {
	a := App{
		Tview:         tview.NewApplication(),
		RootContainer: tview.NewFlex().SetDirection(tview.FlexRow),
		viewContainer: tview.NewFlex(),
		cRequest:      components.CreateRequest(),
		CInfo:         components.CreateInfo(),
		cHistory:      components.CreateHistory(),
		CLog:          components.CreateLog(),
	}

	a.viewContainer.
		AddItem(a.cHistory.Container, 0, 3, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(a.CInfo.Container, 0, 2, false).
			AddItem(a.CLog.Container, 0, 1, false),
			0, 2, false)

	a.RootContainer.
		AddItem(a.viewContainer, 0, 1, false).
		AddItem(a.cRequest.Container, 3, 0, true)

	a.Tview.SetRoot(a.RootContainer, true).SetFocus(a.cRequest.ReqComponent).EnableMouse(true)
	a.SetInputHandlers()

	return &a
}

func (a *App) Connect(cli node.ClientCommunicationClient) {
	a.Cli = cli
	infoNode, err := a.Cli.GetInfoNode(context.Background(), &node.PingRequest{})
	if err != nil {
		return
	}
	a.CInfo.UpdateInfo(infoNode)
}

// Start application
func (a *App) Start() error {
	return a.Tview.Run()
}
