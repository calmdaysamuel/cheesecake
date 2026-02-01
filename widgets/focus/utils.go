package focus

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/state"
	werror "github.com/palantir/witchcraft-go-error"
)

func GainFocusForScope(ctx context.Context, s state.State) error {
	if s == nil {
		return nil
	}
	return werror.WrapWithContextParams(ctx, state.Update(s, ScopeState{InFocus: true}), "failed to gain focus for focus.Scope")
}

func LoseFocusForScope(ctx context.Context, s state.State) error {
	if s == nil {
		return nil
	}
	return werror.WrapWithContextParams(ctx, state.Update(s, ScopeState{InFocus: false}), "failed to lose focus for focus.Scope")
}

func GainFocusForNode(ctx context.Context, s state.State) error {
	if s == nil {
		return nil
	}
	return werror.WrapWithContextParams(ctx, state.Update(s, NodeState{InFocus: true}), "failed to gain focus for focus.Node")
}

func GainPrimaryFocusForNode(ctx context.Context, s state.State) error {
	if s == nil {
		return nil
	}
	return werror.WrapWithContextParams(ctx, state.Update(s, NodeState{InFocus: true, HasPrimaryFocus: true}), "failed to gain focus for focus.Node")
}

func LoseFocusForNode(ctx context.Context, s state.State) error {
	if s == nil {
		return nil
	}
	return werror.WrapWithContextParams(ctx, state.Update(s, NodeState{InFocus: false, HasPrimaryFocus: false}), "failed to lose focus for focus.Node")
}
