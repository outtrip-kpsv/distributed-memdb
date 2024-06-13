package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Request struct {
	Request  string
	Response string
	Color    tcell.Color
}

type HistoryComponent struct {
	Container *tview.Flex
	Component *tview.List
	Items     []Request
}

// AddItem a new request into requests list
func (rl *HistoryComponent) AddItem(req *Request, cb func()) {
	rl.Component.InsertItem(0, req.Request, "    - ["+req.Color.String()+"] "+req.Response, 0, cb).SetSecondaryTextColor(tcell.ColorBlue)
	rl.Items = append(rl.Items, *req)
}

func CreateHistory() *HistoryComponent {
	cp := tview.NewList().SetSelectedFocusOnly(true)

	ct := tview.NewFlex()
	ct.SetDirection(tview.FlexRow).SetBorder(true).SetTitle("history")

	ct.AddItem(cp, 0, 1, false)

	return &HistoryComponent{
		Container: ct,
		Component: cp,
		Items:     make([]Request, 0),
	}
}
