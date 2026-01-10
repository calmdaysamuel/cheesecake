package tree

import (
	"context"
	"reflect"

	"github.com/calmdaysamuel/cheesecake/widget"
)

type Node struct {
	W        widget.Widget
	E        widget.Element
	Children []*Node
}

func Initialize(ctx context.Context, root *Node) {
	if root == nil {
		return
	}
	widgetQueue := []*Node{root}
	for len(widgetQueue) > 0 {
		current := widgetQueue[0]
		if len(widgetQueue) > 0 {
			widgetQueue = widgetQueue[1:]
		}

		switch w := current.W.(type) {
		case widget.StatefulWidget:
			if current.E == nil {
				e := w.Element()
				e.Init()
				current.E = e
			}
			oldChildren := current.Children
			newChildren := []widget.Widget{w.Build(ctx, current.E.(widget.State))}
			current.Children = nil
			for i, child := range newChildren {
				current.Children = append(current.Children, &Node{
					W: child,
				})
				if i < len(oldChildren) {
					current.Children[i].Children = oldChildren[i].Children
					if _, isStateful := oldChildren[i].W.(widget.StatefulWidget); isStateful && reflect.TypeOf(oldChildren[i].W) == reflect.TypeOf(current.Children[i].W) {
						current.Children[i].E = oldChildren[i].E
					} else if oldChildren[i].E != nil {
						if oc, isStateful := oldChildren[i].E.(widget.State); isStateful {
							oc.Dispose()
						}
					}
				}
			}
		case widget.RenderWidget:
			if current.E == nil {
				current.E = w.Element()
			}
			oldChildren := current.Children
			current.Children = nil
			for i, child := range current.E.(widget.RenderElement).DirectDescendants() {
				current.Children = append(current.Children, &Node{
					W: child,
				})
				if i < len(oldChildren) {
					current.Children[i].Children = oldChildren[i].Children
					if _, isStateful := oldChildren[i].W.(widget.StatefulWidget); isStateful && reflect.TypeOf(oldChildren[i].W) == reflect.TypeOf(current.Children[i].W) {
						current.Children[i].E = oldChildren[i].E
					} else if oldChildren[i].E != nil {
						if oc, isStateful := oldChildren[i].E.(widget.State); isStateful {
							oc.Dispose()
						}
					}
				}
			}
		}
		widgetQueue = append(widgetQueue, current.Children...)
	}
}
