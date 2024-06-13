package app

import (
	"context"
	"github.com/gdamore/tcell/v2"
	"team01/internal/client/components"
	"team01/internal/models"
	"team01/internal/proto/node"
	"team01/internal/util"
)

func (a *App) SetInputHandlers() {
	a.RootContainer.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// TODO добавить логику работу в зависимости от активного окна
		//focus := a.Tview.GetFocus()

		switch event.Key() {
		case tcell.KeyEnter:

			req := a.cRequest.ReqComponent.GetText()
			if req == ":q" {
				a.Tview.Stop()
				return event
			}

			response := &node.NodeResp{}
			var err error

			response, err = a.Cli.Request(context.Background(),
				&node.CliReq{
					Req:       ioToRequest(req),
					Addresses: nil,
				})

			if err != nil {
				response = models.ErrorResponse("Err: " + err.Error())
			}

			a.cHistory.AddItem(reqToResp(req, response), func() {
				a.cRequest.ReqComponent.SetText(req)
				a.Tview.SetFocus(a.cRequest.ReqComponent)
			})

			a.cRequest.ReqComponent.SetText("")
		}

		return event
	})

	a.CLog.Component.SetChangedFunc(func() {
		a.Tview.Draw()
	})

	a.cRequest.ReqComponent.SetChangedFunc(func(name string) {
		a.Tview.SetFocus(a.cRequest.ReqComponent)
	})
}

func reqToResp(req string, in *node.NodeResp) *components.Request {
	var res components.Request
	res.Request = req
	switch in.Code {
	case 1:
		res.Color = tcell.ColorGreen
	case 2:
		res.Color = tcell.ColorRed
	}
	res.Response = in.Value
	return &res
}

func ioToRequest(req string) *node.Request {
	reqSlice := util.RequestSplit(req)
	res := node.Request{}
	if len(reqSlice) >= 3 {
		res.Value = reqSlice[2]
	}
	if len(reqSlice) >= 2 {
		res.Uuid = reqSlice[1]
	}
	if len(reqSlice) >= 1 {
		res.Comm = reqSlice[0]
	}
	return &res
}
