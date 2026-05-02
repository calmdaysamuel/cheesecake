package padding

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Options)
type Options struct {
	LeftPadding     int
	RightPadding    int
	TopPadding      int
	BottomPadding   int
	BackgroundColor lipgloss.Color
}

type Model struct {
	Child   widget.Widget
	Options Options
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:      random.ID(),
		options: m.Options,
		child:   m.Child,
	}
}

func New(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child: child,
	}
	for _, option := range options {
		option(&m.Options)
	}
	return m
}
