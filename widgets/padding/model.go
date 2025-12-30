package padding

import (
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
}

func (m *Model) Element() widget.Element {
	return &Element{
		parent: m,
		ID:     random.ID(),
	}
}

func New(child widget.Widget, padding ...int) *Model {
	return &Model{
		Child:   child,
		Padding: padding,
	}
}
