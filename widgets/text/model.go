package text

import (
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Text      string
	style     lipgloss.Style
	styleFunc func(idx int, char rune) *lipgloss.Style
	*mouseactions.Manager
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		parentWidget: m,
		ID:           random.ID(),
		Manager:      m.Manager,
	}
}

func New(text string, options ...Option) *Model {
	m := &Model{Text: text}
	for _, option := range options {
		option(m)
	}
	m.style = m.style.UnsetMargins().UnsetBorderStyle().UnsetPadding()
	return m
}

func WithTextStyle(s lipgloss.Style) Option {
	return func(model *Model) {
		model.style = s
	}
}
func WithTextStyleFunc(styleFunc func(idx int, char rune) *lipgloss.Style) Option {
	return func(model *Model) {
		model.styleFunc = styleFunc
	}
}

func WithMouseActions(m *mouseactions.Manager) Option {
	return func(model *Model) {
		model.Manager = m
	}
}
