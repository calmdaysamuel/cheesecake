package focus

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
	werror "github.com/palantir/witchcraft-go-error"
)

var _ widget.StatefulWidget = &Scope{}

type Scope struct {
	Child func(ctx context.Context, hasFocus bool) (widget.Widget, error)
}

type ScopeState struct {
	InFocus bool
}

func (s *Scope) CreateState(ctx context.Context) (state.State, error) {
	return state.New(ScopeState{})
}

func (s *Scope) Build(ctx context.Context, widgetContext widgetcontext.Context, widgetState state.State) (widget.Widget, error) {
	curVal, err := state.Current[ScopeState](widgetState)
	if err != nil {
		return nil, werror.WrapWithContextParams(ctx, err, "failed to deserialize state object")
	}
	return s.Child(ctx, curVal.InFocus)
}

func NewScope(ChildFunc func(ctx context.Context, hasFocus bool) (widget.Widget, error)) widget.Widget {
	return &Scope{
		Child: ChildFunc,
	}
}

func ScopeEnabled(s state.State) bool {
	return true
}
