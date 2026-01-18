package mousestates

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/column"
	tea "github.com/charmbracelet/bubbletea"
)

type Option func(*Model)

var _ widget.StatefulWidget = &Model{}

type Model struct {
	Default widget.Widget
	Hover   widget.Widget
	Active  widget.Widget
	OnClick func(button tea.MouseButton)
}

type State struct {
	ID string
	*mouseactions.Manager
}

func (s *State) Identifier() string {
	return s.ID
}

func (s *State) Init() {}

func (s *State) Dispose() {}

func (m *Model) Element() widget.State {
	s := &State{Manager: &mouseactions.Manager{OnClick: m.OnClick}, ID: random.ID()}
	return s
}

func (m *Model) Build(ctx context.Context, element widget.State) widget.Widget {
	state := element.(*State)
	if state.Manager.IsActive() && m.Active != nil {
		return column.New([]widget.Widget{m.Active}, column.WithMouseActions(state.Manager))
	}
	if state.Manager.IsHovering() {
		return column.New([]widget.Widget{m.Hover}, column.WithMouseActions(state.Manager))
	}
	return column.New([]widget.Widget{m.Default}, column.WithMouseActions(state.Manager))
}

func New(defaultWidget, hoverWidget widget.Widget, opts ...Option) *Model {
	m := &Model{
		Default: defaultWidget,
		Hover:   hoverWidget,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func WithOnClick(onClick func(button tea.MouseButton)) Option {
	return func(model *Model) {
		model.OnClick = onClick
	}
}

func WithActiveWidget(w widget.Widget) Option {
	return func(model *Model) {
		model.Active = w
	}
}
