package focus

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
	tea "github.com/charmbracelet/bubbletea"
	werror "github.com/palantir/witchcraft-go-error"
)

var _ widget.StatefulWidget = &Node{}

type Option func(*Options)
type KeyEventHandler func(ctx context.Context, msg tea.KeyMsg) error
type Options struct {
	KeyEventHandler KeyEventHandler
}

type Node struct {
	Child   func(ctx context.Context, hasFocus bool, hasPrimaryFocus bool) (widget.Widget, error)
	Options Options
}

func (n *Node) CreateState(ctx context.Context) (state.State, error) {
	return state.New(NodeState{})
}

func (n *Node) Build(ctx context.Context, widgetContext widgetcontext.Context, widgetState state.State) (widget.Widget, error) {
	curVal, err := state.Current[NodeState](widgetState)
	if err != nil {
		return nil, werror.WrapWithContextParams(ctx, err, "failed to deserialize state object")
	}
	return n.Child(ctx, curVal.InFocus, curVal.HasPrimaryFocus)
}

type NodeState struct {
	InFocus         bool
	HasPrimaryFocus bool
}

func NewNode(ChildFunc func(ctx context.Context, hasFocus bool, hasPrimaryFocus bool) (widget.Widget, error), options ...Option) widget.Widget {
	n := &Node{
		Child: ChildFunc,
	}
	for _, option := range options {
		option(&n.Options)
	}
	return n
}
