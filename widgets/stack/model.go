package stack

import (
	"slices"

	"github.com/calmdaysamuel/cheesecake/crossaxis"
	"github.com/calmdaysamuel/cheesecake/mainaxis"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Options)
type Options struct {
	CrossAxisAlignment crossaxis.Alignment
	MainAxisAlignment  mainaxis.Alignment
	BackgroundColor    lipgloss.Color
	ReverseChildren    bool
}

type Model struct {
	Children []widget.Widget
	Options  Options
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		Children: m.Children,
		ID:       random.ID(),
		Options:  m.Options,
	}
}

func New(children []widget.Widget, options ...Option) *Model {
	m := &Model{}
	for _, option := range options {
		option(&m.Options)
	}

	if m.Options.ReverseChildren {
		slices.Reverse(children)
	}
	m.Children = children
	return m
}
