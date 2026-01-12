package spacer

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.StatefulWidget = &InteractiveModel{}

type InteractiveModel struct {
	HoverStyle   lipgloss.Color
	DefaultStyle lipgloss.Color
	ActiveStyle  lipgloss.Color
	Children     []widget.Widget
	Options      []Option
	Flex         int
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

func (m *InteractiveModel) Element() widget.State {
	s := &State{Manager: &mouseactions.Manager{}, ID: random.ID()}
	return s
}

func (m *InteractiveModel) Build(ctx context.Context, element widget.State) widget.Widget {
	state := element.(*State)
	if state.Manager.IsActive() {
		return New(m.Flex, append(m.Options, WithBackgroundColor(m.ActiveStyle), WithMouseActions(state.Manager))...)
	}
	if state.Manager.IsHovering() {
		return New(m.Flex, append(m.Options, WithBackgroundColor(m.HoverStyle), WithMouseActions(state.Manager))...)
	}
	return New(m.Flex, append(m.Options, WithBackgroundColor(m.DefaultStyle), WithMouseActions(state.Manager))...)
}

func Interactive(
	flex int,
	defaultBackgroundColor,
	hoverBackgroundColor,
	activeBackgroundColor lipgloss.Color,
	options ...Option,
) *InteractiveModel {
	return &InteractiveModel{
		Flex:         flex,
		DefaultStyle: defaultBackgroundColor,
		HoverStyle:   hoverBackgroundColor,
		ActiveStyle:  activeBackgroundColor,
		Options:      options,
	}
}
