package focus

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
	werror "github.com/palantir/witchcraft-go-error"
)

var _ widget.StatefulWidget = &Node{}

type Node struct {
	Child func(ctx context.Context, hasFocus bool) (widget.Widget, error)
}

func (n *Node) CreateState(ctx context.Context) (state.State, error) {
	return state.New(NodeState{})
}

func (n *Node) Build(ctx context.Context, widgetContext widgetcontext.Context, widgetState state.State) (widget.Widget, error) {
	curVal, err := state.Current[NodeState](widgetState)
	if err != nil {
		return nil, werror.WrapWithContextParams(ctx, err, "failed to deserialize state object")
	}
	return n.Child(ctx, curVal.InFocus)
}

type NodeState struct {
	InFocus bool
}

func PrimaryFocusNode(ChildFunc func(ctx context.Context, hasFocus bool) (widget.Widget, error)) widget.Widget {
	return &Node{
		Child: ChildFunc,
	}
}

func GainFocus(s state.State) error {
	if err := state.Update(s, NodeState{InFocus: true}); err != nil {
		return state.Update(s, ScopeState{InFocus: true})
	}
	return nil
}

func LoseFocus(s state.State) error {
	if err := state.Update(s, NodeState{InFocus: false}); err != nil {
		return state.Update(s, ScopeState{InFocus: false})
	}
	return nil
}
