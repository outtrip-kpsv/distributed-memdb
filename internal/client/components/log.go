package components

import (
	"fmt"
	"github.com/rivo/tview"
)

type LogComponent struct {
	Container *tview.Flex
	Component *tview.TextView
	Log       chan string
}

func CreateLog() *LogComponent {
	cp := tview.NewTextView().SetDynamicColors(true)

	ct := tview.NewFlex()
	ct.SetDirection(tview.FlexRow).SetBorder(true).SetTitle("Log")

	ct.AddItem(cp, 0, 1, false)
	res := &LogComponent{
		Container: ct,
		Component: cp,
		Log:       make(chan string, 20),
	}

	return res
}

func (l *LogComponent) LogWorker() {
	go func() {
		for {
			select {
			case log := <-l.Log:
				text := fmt.Sprintf("%s\n%s", l.Component.GetText(true), log)
				l.Component.SetText(text)
				l.Component.ScrollToEnd()
			}
		}
	}()
}
