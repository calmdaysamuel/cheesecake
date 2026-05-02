package border

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Options)
type Options struct {
	EnableLeftBorder   bool
	EnableRightBorder  bool
	EnableTopBorder    bool
	EnableBottomBorder bool
	BackgroundColor    lipgloss.Color
	BorderColor        lipgloss.Color
	BorderStyle        lipgloss.Border
}
type Model struct {
	Child   widget.Widget
	Options Options
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:      random.ID(),
		child:   m.Child,
		options: m.Options,
	}
}

func New(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child: child,
		Options: Options{
			EnableLeftBorder:   true,
			EnableRightBorder:  true,
			EnableTopBorder:    true,
			EnableBottomBorder: true,
			BorderStyle:        lipgloss.NormalBorder(),
		},
	}

	for _, option := range options {
		option(&m.Options)
	}
	return m
}
