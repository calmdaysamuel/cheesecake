package tree

import (
	"context"
	"errors"
	"reflect"

	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
	werror "github.com/palantir/witchcraft-go-error"
	wparams "github.com/palantir/witchcraft-go-params"
)

const StatefulWidgetType = "stateful-widget"
const RenderWidgetType = "render-widget"
const UnknownWidgetType = "unknown-widget"

type StatefulNode struct {
	Widget widget.StatefulWidget
	State  state.State
	Child  *Node
}

type RenderNode struct {
	Widget   widget.RenderWidget
	Element  widget.RenderElement
	Children []*Node
}

type Node struct {
	Stateful     *StatefulNode
	RenderWidget *RenderNode
}

func (n *Node) Type() string {
	if n.Stateful != nil {
		return StatefulWidgetType
	} else if n.RenderWidget != nil {
		return RenderWidgetType
	}
	return UnknownWidgetType
}

func (n *Node) IsSameType(other *Node) bool {
	if other == nil {
		return false
	}
	if n.Type() != other.Type() {
		return false
	}
	switch n.Type() {
	case StatefulWidgetType:
		return reflect.TypeOf(n.Stateful.Widget) == reflect.TypeOf(other.Stateful.Widget)
	case RenderWidgetType:
		return reflect.TypeOf(n.RenderWidget.Widget) == reflect.TypeOf(other.RenderWidget.Widget)
	}
	return false
}

func RefreshWidgetTree(ctx context.Context, root *Node, subtreeDirty bool) (shouldRerender bool, err error) {
	if root == nil {
		return false, nil
	}
	switch root.Type() {
	case StatefulWidgetType:
		ctx = wparams.ContextWithSafeParam(ctx, "nodeType", StatefulWidgetType)
		sw := root.Stateful
		isDirty, err := state.IsDirty(sw.State)
		if err != nil {
			return false, werror.WrapWithContextParams(ctx, err, "failed to check if stateful widget is dirty")
		}
		ctx = wparams.ContextWithSafeParam(ctx, "isDirty", isDirty)
		if isDirty {
			subtreeDirty = true
		}
		switch {
		case sw.Child == nil || isDirty || subtreeDirty:
			oldChild := sw.Child
			newChild, err := sw.Widget.Build(ctx, nil, sw.State)
			if err != nil {
				return false, werror.WrapWithContextParams(ctx, err, "failed to build stateful widget")
			}
			switch n := newChild.(type) {
			case widget.StatefulWidget:
				var preferredChild *Node
				if oldChild != nil {
					if oldChild.Stateful != nil {
						preferredChild = oldChild.Stateful.Child
					} else if oldChild.RenderWidget != nil {
						if len(oldChild.RenderWidget.Children) == 1 {
							preferredChild = oldChild.RenderWidget.Children[0]
						}
					}
				}
				sw.Child = &Node{
					Stateful: &StatefulNode{
						Widget: n,
						Child:  preferredChild,
					},
				}
				if sw.Child.IsSameType(oldChild) {
					sw.Child.Stateful.State = oldChild.Stateful.State
				} else {
					sw.Child.Stateful.State, err = sw.Child.Stateful.Widget.CreateState(ctx)
					if err != nil {
						return false, werror.WrapWithContextParams(ctx, err, "failed to create state")
					}
				}
			case widget.RenderWidget:
				var preferredChildren []*Node
				if oldChild != nil {
					if oldChild.Stateful != nil && oldChild.Stateful.Child != nil {
						preferredChildren = []*Node{oldChild.Stateful.Child}
					} else if oldChild.RenderWidget != nil {
						preferredChildren = oldChild.RenderWidget.Children
					}
				}
				sw.Child = &Node{
					RenderWidget: &RenderNode{
						Widget:   n,
						Children: preferredChildren,
					},
				}
			default:
				return false, errors.New("tree.RefreshWidgetTree: child widget type not recognized")
			}
			shouldRerenderChild, err := RefreshWidgetTree(ctx, sw.Child, subtreeDirty)
			if err != nil {
				return false, werror.ErrorWithContextParams(ctx, "failed to refresh widget subtree")
			}
			subtreeDirty = shouldRerenderChild || subtreeDirty
			if err := state.Clean(sw.State); err != nil {
				return false, werror.ErrorWithContextParams(ctx, "failed to clean state")
			}
		default:
			shouldRerenderChild, err := RefreshWidgetTree(ctx, sw.Child, subtreeDirty)
			if err != nil {
				return false, werror.ErrorWithContextParams(ctx, "failed to refresh widget subtree")
			}
			return shouldRerenderChild || subtreeDirty, nil
		}
	case RenderWidgetType:
		ctx = wparams.ContextWithSafeParam(ctx, "nodeType", RenderWidgetType)
		root.RenderWidget.Element = root.RenderWidget.Widget.Element()
		oldChildren := root.RenderWidget.Children
		root.RenderWidget.Children = nil
		descendants, err := root.RenderWidget.Element.DirectDescendants(ctx)
		if err != nil {
			return false, werror.WrapWithContextParams(ctx, err, "failed to compute descendant for widget")
		}
		for i, child := range descendants {
			switch child := child.(type) {
			case widget.StatefulWidget:
				node := &Node{
					Stateful: &StatefulNode{
						Widget: child,
					},
				}
				if i < len(oldChildren) {
					if node.IsSameType(oldChildren[i]) {
						node.Stateful.Child = oldChildren[i].Stateful.Child
						node.Stateful.State = oldChildren[i].Stateful.State
					}
				}
				if node.Stateful.State == nil {
					var err error
					node.Stateful.State, err = node.Stateful.Widget.CreateState(ctx)
					if err != nil {
						return false, err
					}
				}
				root.RenderWidget.Children = append(root.RenderWidget.Children, node)
				shouldRerenderChild, err := RefreshWidgetTree(ctx, node, subtreeDirty)
				subtreeDirty = shouldRerenderChild || subtreeDirty
				if err != nil {
					return false, werror.ErrorWithContextParams(ctx, "failed to refresh widget subtree")
				}
			case widget.RenderWidget:
				node := &Node{
					RenderWidget: &RenderNode{
						Widget: child,
					},
				}
				if i < len(oldChildren) {
					if oldChildren[i].RenderWidget != nil {
						node.RenderWidget.Children = oldChildren[i].RenderWidget.Children
					} else if oldChildren[i].Stateful != nil {
						node.RenderWidget.Children = []*Node{oldChildren[i].Stateful.Child}
					}
				}
				root.RenderWidget.Children = append(root.RenderWidget.Children, node)
				shouldRerenderChild, err := RefreshWidgetTree(ctx, node, subtreeDirty)
				subtreeDirty = shouldRerenderChild || subtreeDirty
				if err != nil {
					return false, werror.ErrorWithContextParams(ctx, "failed to refresh widget subtree")
				}
			default:
				return false, werror.ErrorWithContextParams(ctx, "child widget type of recognized")
			}
		}
	default:
		ctx = wparams.ContextWithSafeParam(ctx, "nodeType", UnknownWidgetType)
		return false, werror.ErrorWithContextParams(ctx, "unknown widget type provided to the widget tree refresher")
	}
	return subtreeDirty, nil
}

func NodeFromWidget(ctx context.Context, root widget.Widget) (*Node, error) {
	switch c := root.(type) {
	case widget.StatefulWidget:
		s, err := c.CreateState(ctx)
		if err != nil {
			return nil, werror.WrapWithContextParams(ctx, err, "failed to initialize first application state")
		}
		return &Node{
			Stateful: &StatefulNode{
				Widget: c,
				State:  s,
			},
		}, nil
	case widget.RenderWidget:
		return &Node{
			RenderWidget: &RenderNode{
				Widget:  c,
				Element: c.Element(),
			},
		}, nil
	}
	return nil, werror.ErrorWithContextParams(ctx, "failed to determine type of widget. Cannot initialize application.")
}
