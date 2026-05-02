package row

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widgets/flex"
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
	Options  flex.Options
	children []widget.Widget
}

func (m *Model) Element() widget.RenderElement {
	return &flex.Element{
		Children: m.children,
		ID:       random.ID(),
		Options:  m.Options,
		FlexProvider: func(f widget.Flexible) int {
			val, _ := f.Flex()
			return val
		},
		SizeProvider: func(size size.Size) int {
			return size.Width
		},
		ConstraintsProvider: func(box constraints.Constraints) int {
			return box.MaxWidth
		},
		ChildJoiner: func(canvases ...canvas.Canvas) canvas.Canvas {
			return canvas.JoinHorizontal(
				lipgloss.Position(
					m.Options.CrossAxisAlignment,
				),
				canvases...,
			)
		},
	}
}

func New(children []widget.Widget, options ...flex.Option) *Model {
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
