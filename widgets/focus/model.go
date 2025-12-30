package focus

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/widget"
	tea "github.com/charmbracelet/bubbletea"
)

var _ widget.StatefulWidget = &Model{}

type Option func(*Model)

type Model struct {
	Child       widget.Widget
	onFocusGain func()
	onFocusLoss func()
	onKeyPress  func(key tea.KeyMsg)
	ChildFunc   func(inFocus bool) widget.Widget
}

func (m *Model) Build(_ context.Context, state widget.State) widget.Widget {
	if m.ChildFunc != nil {
		return m.ChildFunc(state.(*widget.StatefulElementImpl[State]).Current().InFocus)
	}
	return m.Child
}

func (m *Model) Element() widget.Element {
	return &Element{
		StatefulElementImpl: widget.NewStatefulElement(State{
			InFocus:     false,
			OnKeyPress:  m.onKeyPress,
			OnFocusGain: m.onFocusGain,
			OnFocusLoss: m.onFocusLoss,
		}),
	}
}

func New(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child: child,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func NewBuilder(childFunc func(inFocus bool) widget.Widget, options ...Option) *Model {
	m := &Model{
		ChildFunc: childFunc,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func WithOnFocusGain(on func()) Option {
	return func(model *Model) {
		model.onFocusGain = on
	}
}

func WithOnFocusLoss(on func()) Option {
	return func(model *Model) {
		model.onFocusLoss = on
	}
}

func WithOnKeyPress(on func(msg tea.KeyMsg)) Option {
	return func(model *Model) {
		model.onKeyPress = on
	}
}
