package tree

import (
	"context"
	"slices"

	"github.com/Goldziher/go-utils/sliceutils"
	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widgets/focus"
	werror "github.com/palantir/witchcraft-go-error"
)

type FocusMoveDirection = int

const (
	FocusMoveForward  FocusMoveDirection = 1
	FocusMoveBackward FocusMoveDirection = -1
)

type FocusNode struct {
	FocusNode       state.State
	Chain           []state.State
	ParentScope     state.State
	KeyEventHandler focus.KeyEventHandler
	ScopeEnabled    bool
}

func RefreshFocusTree(ctx context.Context, root *Node, chain []state.State) []*FocusNode {
	if root == nil {
		return nil
	}
	focusNodes := make([]*FocusNode, 0)
	switch root.Type() {
	case StatefulWidgetType:
		switch f := root.Stateful.Widget.(type) {
		case *focus.Scope:
			childNodes := RefreshFocusTree(ctx, root.Stateful.Child, chain)
			for _, node := range childNodes {
				if node.ParentScope == nil {
					node.ParentScope = root.Stateful.State
					node.ScopeEnabled = focus.ScopeEnabled(root.Stateful.State)
				}
			}
			focusNodes = append(focusNodes, childNodes...)
		case *focus.Node:
			childNodes := RefreshFocusTree(ctx, root.Stateful.Child, append(chain, root.Stateful.State))
			focusNodes = append(focusNodes, childNodes...)
			focusNodes = append(focusNodes, &FocusNode{FocusNode: root.Stateful.State, Chain: chain, KeyEventHandler: f.Options.KeyEventHandler})
		default:
			childNodes := RefreshFocusTree(ctx, root.Stateful.Child, chain)
			focusNodes = append(focusNodes, childNodes...)
		}
	case RenderWidgetType:
		for _, child := range root.RenderWidget.Children {
			childNodes := RefreshFocusTree(ctx, child, chain)
			focusNodes = append(focusNodes, childNodes...)
		}
	}
	return focusNodes
}

func FocusTreeStep(ctx context.Context, nodes []*FocusNode, previous *FocusNode, ignoreScope bool, direction FocusMoveDirection) (newFocusedNode *FocusNode, rErr error) {
	defer func() {
		if previous != nil {
			if rErr = werror.WrapWithContextParams(ctx, focus.LoseFocusForScope(ctx, previous.ParentScope), "failed to lose focus for previous focus scope"); rErr != nil {
				return
			}
			if rErr = werror.WrapWithContextParams(ctx, focus.LoseFocusForNode(ctx, previous.FocusNode), "failed to lose focus for previous focus node"); rErr != nil {
				return
			}
			for _, nodeInChain := range previous.Chain {
				if rErr = werror.WrapWithContextParams(ctx, focus.LoseFocusForNode(ctx, nodeInChain), "failed to lose focus for node in previous focus chain"); rErr != nil {
					return
				}
			}
		}
		if newFocusedNode != nil {
			if rErr = werror.WrapWithContextParams(ctx, focus.GainFocusForScope(ctx, newFocusedNode.ParentScope), "failed to gain focus for new focus scope"); rErr != nil {
				return
			}
			if rErr = werror.WrapWithContextParams(ctx, focus.GainPrimaryFocusForNode(ctx, newFocusedNode.FocusNode), "failed to gain primary focus for current focus node"); rErr != nil {
				return
			}
			for _, s := range newFocusedNode.Chain {
				if rErr = werror.WrapWithContextParams(ctx, focus.GainFocusForNode(ctx, s), "failed to gain focus for node in current focus chain"); rErr != nil {
					return
				}
			}
		}
	}()
	if !ignoreScope {
		nodes = sliceutils.Filter(nodes, func(value *FocusNode, index int, slice []*FocusNode) bool {
			if previous == nil || previous.ParentScope == nil || !previous.ScopeEnabled {
				return true
			}
			return value.ParentScope != nil && (value.ParentScope.ID() == previous.ParentScope.ID())
		})
	}
	if len(nodes) == 0 {
		return nil, nil
	}
	previousIndex := slices.IndexFunc(nodes, func(node *FocusNode) bool {
		if previous == nil {
			return false
		}
		return node != nil && node.FocusNode.ID() == previous.FocusNode.ID()
	})
	if previousIndex == -1 {
		return nodes[0], nil
	}
	nextIndex := previousIndex + direction
	if nextIndex >= len(nodes) {
		nextIndex = 0
	} else if nextIndex < 0 {
		nextIndex = len(nodes) - 1
	}
	return nodes[nextIndex], nil
}
