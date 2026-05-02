package spacer

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

type Option func(*Options)
type Options struct {
	BackgroundColor lipgloss.Color
}

var _ widget.RenderWidget = &Model{}

type Model struct {
	flex    int
	Options Options
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:      random.ID(),
		flex:    m.flex,
		Options: m.Options,
	}
}

func New(flex int, options ...Option) *Model {
	m := &Model{flex: flex}
	for _, option := range options {
		option(&m.Options)
	}
	return m
}
