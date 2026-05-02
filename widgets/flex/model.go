package flex

import (
	"slices"
	"strings"

	"github.com/calmdaysamuel/cheesecake/crossaxis"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/text"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Options)
type Options struct {
	Spacing            int
	CrossAxisAlignment crossaxis.Alignment
	BackgroundColor    lipgloss.Color
	ReverseChildren    bool
}
type Model struct {
	Options  Options
	children []widget.Widget
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		Children: m.children,
		ID:       random.ID(),
		Options:  m.Options,
	}
}

func New(children []widget.Widget, options ...Option) *Model {
	m := &Model{}
	for _, option := range options {
		option(&m.Options)
	}
	withSpacing := make([]widget.Widget, 0)

	// Add spacing for row
	for i, child := range children {
		withSpacing = append(withSpacing, child)
		if i < len(children)-1 {
			withSpacing = append(withSpacing, text.New(strings.Repeat(" ", m.Options.Spacing)))
		}
	}
	if m.Options.ReverseChildren {
		slices.Reverse(withSpacing)
	}
	m.children = withSpacing
	return m
}
