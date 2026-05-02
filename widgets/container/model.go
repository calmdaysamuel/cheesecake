package container

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
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
	*mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		parent:  m,
		ID:      random.ID(),
		Manager: m.Manager,
	}
}

func New(child widget.Widget) *Model {
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
