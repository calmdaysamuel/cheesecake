package container

import (
	"cheesecake/random"
	"cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Width               int
	Height              int
	Child               widget.Widget
	BgColor             lipgloss.Color
	VerticalAlignment   lipgloss.Position
	HorizontalAlignment lipgloss.Position
}

func (m *Model) Element() widget.Element {
	return &Element{
		parent: m,
		ID:     random.ID(),
	}
}

func New(child widget.Widget, height, width int, options ...Option) *Model {
	m := &Model{
		Width:               width,
		Height:              height,
		Child:               child,
		VerticalAlignment:   lipgloss.Top,
		HorizontalAlignment: lipgloss.Left,
	}

	for _, option := range options {
		option(m)
	}
	return m
}

func WithBackgroundColor(color lipgloss.Color) Option {
	return func(model *Model) {
		model.BgColor = color
	}
}

func WithVerticalAlignment(position lipgloss.Position) Option {
	return func(model *Model) {
		model.VerticalAlignment = position
	}
}

func WithHorizontalAlignment(position lipgloss.Position) Option {
	return func(model *Model) {
		model.HorizontalAlignment = position
	}
}
