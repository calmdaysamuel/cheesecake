package focus

import (
	"github.com/calmdaysamuel/cheesecake/widget"
	tea "github.com/charmbracelet/bubbletea"
)

var _ widget.Element = &Element{}

type Element struct {
	HasFocus    bool
	OnFocusGain func()
	OnKeyPress  func(key tea.KeyMsg)
	OnFocusLoss func()
	ID          string
}

func (e *Element) Init() {
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) Dispose() {}

func (e *Element) LoseFocus() {
	e.HasFocus = false
	if e.OnFocusLoss != nil {
		e.OnFocusLoss()
	}
}

func (e *Element) GainLocus() {
	e.HasFocus = true
	if e.OnFocusGain != nil {
		e.OnFocusGain()
	}
}

func (e *Element) OnKeyPressEvent(msg tea.KeyMsg) {
	if e.OnKeyPress != nil {
		e.OnKeyPress(msg)
	}
}

func (e *Element) InFocus() bool {
	return e.HasFocus
}
