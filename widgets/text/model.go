package text

import (
	"strings"

	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Text      string
	Style     lipgloss.Style
	StyleFunc func(idx int, char rune) *lipgloss.Style
	mouseactions.Manager
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
	m.Style = m.Style.UnsetMargins().UnsetBorderStyle().UnsetPadding()
	return m
}

func Square(sideLength int, options ...Option) *Model {
	row := strings.Repeat(" ", sideLength*2)
	columns := make([]string, 0)
	for i := 0; i < sideLength; i++ {
		columns = append(columns, row)
	}
	m := &Model{Text: strings.Join(columns, "\n")}
	for _, option := range options {
		option(m)
	}
	m.Style = m.Style.UnsetMargins().UnsetBorderStyle().UnsetPadding()
	return m
}

func Place(text string, width, height int, options ...Option) *Model {
	m := &Model{Text: lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, text)}
	for _, option := range options {
		option(m)
	}
	m.Style = m.Style.UnsetMargins().UnsetBorderStyle().UnsetPadding()
	return m
}

func WithTextStyle(s lipgloss.Style) Option {
	return func(model *Model) {
		model.Style = s
	}
}
func WithTextStyleFunc(styleFunc func(idx int, char rune) *lipgloss.Style) Option {
	return func(model *Model) {
		model.StyleFunc = styleFunc
	}
}

func WithMouseActions(m mouseactions.Manager) Option {
	return func(model *Model) {
		model.Manager = m
	}
}
