package alignment

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
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

func Center(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child:               child,
		VerticalAlignment:   lipgloss.Center,
		HorizontalAlignment: lipgloss.Center,
	}

	for _, option := range options {
		option(m)
	}
	return m
}

func TopLeft(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child:               child,
		VerticalAlignment:   lipgloss.Top,
		HorizontalAlignment: lipgloss.Left,
	}

	for _, option := range options {
		option(m)
	}
	return m
}

func TopRight(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child:               child,
		VerticalAlignment:   lipgloss.Top,
		HorizontalAlignment: lipgloss.Right,
	}

	for _, option := range options {
		option(m)
	}
	return m
}

func BottomLeft(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child:               child,
		VerticalAlignment:   lipgloss.Bottom,
		HorizontalAlignment: lipgloss.Left,
	}

	for _, option := range options {
		option(m)
	}
	return m
}

func BottomRight(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child:               child,
		VerticalAlignment:   lipgloss.Bottom,
		HorizontalAlignment: lipgloss.Right,
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
