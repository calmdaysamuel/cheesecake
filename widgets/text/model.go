package text

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Options)
type Options struct {
	ForegroundColor lipgloss.Color
	BackgroundColor lipgloss.Color
	Alignment       lipgloss.Position
	Bold            bool
	Faint           bool
	Underline       bool
	UnderlineSpaces bool
	Italic          bool
	ShouldWrap      bool
}

type Model struct {
	Text    string
	Options Options
	*mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:      random.ID(),
		Manager: m.Manager,
		options: m.Options,
		text:    m.Text,
	}
}

func New(text string, options ...Option) widget.Widget {
	m := &Model{Text: text}
	for _, option := range options {
		option(&m.Options)
	}
	return m
}
