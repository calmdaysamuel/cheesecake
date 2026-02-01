package layoutbuilder

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
	werror "github.com/palantir/witchcraft-go-error"
)

var _ widget.StatefulWidget = &StatefulModel{}

type StatefulModel struct {
	ChildFunc func(box constraints.Constraints) (widget.Widget, error)
}

type State struct {
	LastConstraints constraints.Constraints
}

func (m *StatefulModel) CreateState(ctx context.Context) (state.State, error) {
	return state.New(State{})
}

func (m *StatefulModel) Build(ctx context.Context, wcontext widgetcontext.Context, element state.State) (widget.Widget, error) {
	currentVal, err := state.Current[State](element)
	if err != nil {
		return nil, err
	}
	w, err := m.ChildFunc(currentVal.LastConstraints)
	if err != nil {
		return nil, werror.WrapWithContextParams(ctx, err, "failed to create widget from childFunc in layout builder.")
	}
	if w == nil {
		return nil, werror.ErrorWithContextParams(ctx, "failed to create widget from childFunc in layout builder. child is nil")
	}
	return &Model{
		Child:           w,
		LastConstraints: currentVal.LastConstraints,
		ConstraintsListener: func(c constraints.Constraints) error {
			if c == currentVal.LastConstraints {
				return nil // constraints must change to trigger widget rebuild
			}
			return state.Update(element, State{LastConstraints: c})
		},
	}, nil
}

func New(childFunc func(box constraints.Constraints) (widget.Widget, error)) *StatefulModel {
	return &StatefulModel{ChildFunc: childFunc}
}
