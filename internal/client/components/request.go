package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type RequestComponent struct {
	Container    *tview.Flex
	ReqComponent *tview.InputField
}

func CreateRequest() *RequestComponent {
	ct := tview.NewFlex()
	ct.SetDirection(tview.FlexRow).SetBorder(true).SetTitle("Request")

	req := tview.NewInputField().SetFieldBackgroundColor(tcell.ColorGreen).SetFieldTextColor(tcell.ColorBlack)
	req.SetTitle("Response")
	ct.
		AddItem(req, 0, 1, false)

	return &RequestComponent{
		Container:    ct,
		ReqComponent: req,
	}
}
