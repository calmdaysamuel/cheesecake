package padding

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Child   widget.Widget
	BgColor lipgloss.Color
	Padding []int
	mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		parent:  m,
		ID:      random.ID(),
		Manager: m.Manager,
	}
}

func New(child widget.Widget, padding ...int) *Model {
	return &Model{
		Child:   child,
		Padding: padding,
	}
}

func WithMouseActions(m mouseactions.Manager) Option {
	return func(model *Model) {
		model.Manager = m
	}
}
