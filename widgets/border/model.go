package border

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Child  widget.Widget
	Border lipgloss.Border
	Label  string
	Sides  []bool
	Style  lipgloss.Style
	mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		parent:  m,
		ID:      random.ID(),
		Manager: m.Manager,
	}
}

func New(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child:  child,
		Border: lipgloss.NormalBorder(),
		Sides:  []bool{true},
	}

	for _, option := range options {
		option(m)
	}
	return m
}

func WithBorder(border lipgloss.Border) Option {
	return func(model *Model) {
		model.Border = border
	}
}
func WithBorderTopLabel(label string) Option {
	return func(model *Model) {
		model.Label = label
	}
}

func WithBorderStyle(style lipgloss.Style) Option {
	return func(model *Model) {
		model.Style = style
	}
}

func WithSides(sides ...bool) Option {
	return func(model *Model) {
		model.Sides = sides
	}
}

func WithMouseActions(m mouseactions.Manager) Option {
	return func(model *Model) {
		model.Manager = m
	}
}
