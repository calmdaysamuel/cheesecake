package layoutbuilder

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.StatefulWidget = &StatefulModel{}

type StatefulModel struct {
	ChildFunc func(constraints constraints.Constraints) widget.Widget
}

type State struct {
	ID              string
	LastConstraints constraints.Constraints
}

func (s *State) Identifier() string {
	return s.ID
}

func (s *State) Init() {}

func (s *State) Dispose() {}

func (m *StatefulModel) Element() widget.State {
	return &State{ID: random.ID()}
}

func (m *StatefulModel) Build(ctx context.Context, element widget.State) widget.Widget {
	state := element.(*State)
	return &Model{
		ChildFunc:       m.ChildFunc,
		LastConstraints: state.LastConstraints,
		ConstraintsListener: func(c constraints.Constraints) {
			state.LastConstraints = c
		},
	}
}

func New(childFunc func(constraints constraints.Constraints) widget.Widget) *StatefulModel {
	return &StatefulModel{ChildFunc: childFunc}
}
