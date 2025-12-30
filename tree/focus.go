package tree

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/widgets/focus"
)

func Focus(ctx context.Context, root *Node) []*focus.Element {
	focusElements := make([]*focus.Element, 0)
	for _, child := range root.Children {
		focusElements = append(focusElements, Focus(ctx, child)...)
	}
	if f, ok := root.E.(*focus.Element); ok {
		focusElements = append(focusElements, f)
	}
	return focusElements
}
