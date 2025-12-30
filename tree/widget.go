package tree

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/widget"
	"reflect"
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
			initialized := false
			if current.E == nil {
				initialized = true
				current.E = w.Element()
			}
			state := current.E.(widget.StatefulElement).GetState()
			if state.Dirty() || initialized {
				oldChildren := current.Children
				newChildren := []widget.Widget{w.Build(ctx, state)}
				current.Children = nil
				for i, child := range newChildren {
					current.Children = append(current.Children, &Node{
						W: child,
					})
					if i < len(oldChildren) {
						if _, isStateful := oldChildren[i].W.(widget.StatefulElement); isStateful && reflect.TypeOf(oldChildren[i].W) == reflect.TypeOf(current.Children[i].W) {
							current.Children[i].E = oldChildren[i].E
							current.Children[i].Children = oldChildren[i].Children
						} else if oldChildren[i].E != nil {
							oldChildren[i].E.Dispose()
						}
					}
				}
				state.Clean()
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
					if reflect.TypeOf(oldChildren[i].W) == reflect.TypeOf(current.Children[i].W) {
						current.Children[i].E = oldChildren[i].E
						current.Children[i].Children = oldChildren[i].Children
					} else if oldChildren[i].E != nil {
						oldChildren[i].E.Dispose()
					}
				}
			}
		}
		widgetQueue = append(widgetQueue, current.Children...)
	}
}
