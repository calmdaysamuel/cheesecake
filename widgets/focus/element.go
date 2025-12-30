package focus

import (
	"github.com/calmdaysamuel/cheesecake/widget"
	tea "github.com/charmbracelet/bubbletea"
)

var _ widget.StatefulElement = &Element{}

type State struct {
	InFocus     bool
	OnFocusGain func()
	OnKeyPress  func(key tea.KeyMsg)
	OnFocusLoss func()
}
type Element struct {
	*widget.StatefulElementImpl[State]
}

func (e *Element) LoseFocus() {
	e.SetState(func(oldState State) State {
		oldState.InFocus = false
		if oldState.OnFocusLoss != nil {
			oldState.OnFocusLoss()
		}
		return oldState
	})
}

func (e *Element) GainLocus() {
	e.SetState(func(oldState State) State {
		oldState.InFocus = true
		if oldState.OnFocusGain != nil {
			oldState.OnFocusGain()
		}
		return oldState
	})
}

func (e *Element) OnKeyPressEvent(msg tea.KeyMsg) {
	if e.Current().OnKeyPress != nil {
		e.Current().OnKeyPress(msg)
	}
}

func (e *Element) InFocus() bool {
	return e.Current().InFocus
}
