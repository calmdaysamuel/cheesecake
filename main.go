package main

import (
	"context"
	"fmt"

	"github.com/calmdaysamuel/cheesecake/application"
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/text"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	_ = application.Start(&Model{})
}

var _ widget.StatefulWidget = &Model{}

type Model struct {
}
type State struct {
	ID        string
	TextStyle lipgloss.Style
}

func (s *State) Identifier() string {
	return s.ID
}

func (s *State) Init() {}

func (s *State) Dispose() {}

func (m *Model) Element() widget.State {
	return &State{}
}

func (m *Model) Build(ctx context.Context, element widget.State) widget.Widget {
	state := element.(*State)
	return text.New("hello", text.WithTextStyle(state.TextStyle), text.WithMouseActions(mouseactions.Manager{
		OnHoverStateChangeFunc: nil,
		OnClickFunc: func() {
			fmt.Println("hello clicked")
			state.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("5"))
		},
	}))
}
