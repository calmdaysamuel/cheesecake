package tree

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/widget"
	werror "github.com/palantir/witchcraft-go-error"
	wparams "github.com/palantir/witchcraft-go-params"
)

func RootRenderObject(ctx context.Context, root *Node) (widget.RenderElement, error) {
	ctx = wparams.ContextWithSafeParam(ctx, "component", "findRootRenderObject")
	elements := RefreshRenderTree(ctx, root)
	if len(elements) == 0 {
		return nil, werror.ErrorWithContextParams(ctx, "no render-able widget has been provided")
	}
	if len(elements) > 1 {
		return nil, werror.ErrorWithContextParams(ctx, "discovered more than one root render object widget")
	}
	return elements[0], nil
}

func RefreshRenderTree(ctx context.Context, root *Node) []widget.RenderElement {
	ctx = wparams.ContextWithSafeParam(ctx, "component", "refreshRenderTree")
	renderElementChildren := make([]widget.RenderElement, 0)
	switch root.Type() {
	case StatefulWidgetType:
		childRenderObjects := RefreshRenderTree(ctx, root.Stateful.Child)
		renderElementChildren = append(renderElementChildren, childRenderObjects...)
	case RenderWidgetType:
		for _, child := range root.RenderWidget.Children {
			childRenderObjects := RefreshRenderTree(ctx, child)
			renderElementChildren = append(renderElementChildren, childRenderObjects...)
		}
		root.RenderWidget.Element.ClearChildren()
		for _, child := range renderElementChildren {
			root.RenderWidget.Element.AdoptChild(child)
		}
		return []widget.RenderElement{root.RenderWidget.Element}
	}
	return renderElementChildren
}
