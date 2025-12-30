package tree

import (
	"cheesecake/widget"
)

func RootRenderObject(root *Node) widget.RenderElement {
	elements := RenderObjectTree(root)
	if len(elements) == 0 {
		return nil
	}
	if len(elements) > 1 {
		panic("strange: there must never more than one root render object")
	}
	return elements[0]
}

func RenderObjectTree(root *Node) []widget.RenderElement {

	renderElementChildren := make([]widget.RenderElement, 0)
	for _, child := range root.Children {
		if rc := RenderObjectTree(child); rc != nil {
			renderElementChildren = append(renderElementChildren, rc...)
		}
	}
	if renderObject, ok := root.E.(widget.RenderElement); ok {
		renderObject.ClearChildren()
		for _, child := range renderElementChildren {
			renderObject.AdoptChild(child)
		}
		return []widget.RenderElement{renderObject}
	}
	return renderElementChildren
}
