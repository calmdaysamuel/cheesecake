package text

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.StatefulWidget = &InteractiveModel{}

type InteractiveModel struct {
	Text         string
	HoverStyle   lipgloss.Style
	DefaultStyle lipgloss.Style
	ActiveStyle  lipgloss.Style
	Options      []Option
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
		return New(m.Text, append(m.Options, WithTextStyle(m.ActiveStyle), WithMouseActions(state.Manager))...)
	}
	if state.Manager.IsHovering() {
		return New(m.Text, append(m.Options, WithTextStyle(m.HoverStyle), WithMouseActions(state.Manager))...)
	}
	return New(m.Text, append(m.Options, WithTextStyle(m.DefaultStyle), WithMouseActions(state.Manager))...)
}

func Interactive(text string, defaultStyle, hoverStyle, activeStyle lipgloss.Style, opts ...Option) *InteractiveModel {
	return &InteractiveModel{
		Text:         text,
		DefaultStyle: defaultStyle,
		HoverStyle:   hoverStyle,
		ActiveStyle:  activeStyle,
		Options:      opts,
	}
}
