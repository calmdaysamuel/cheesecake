package widget

import (
	"cheesecake/canvas"
	"cheesecake/constraints"
	"cheesecake/size"
)

type RenderElement interface {
	Element
	AdoptChild(child RenderElement)
	ClearChildren()
	DirectDescendants() []Widget
	View() canvas.Canvas
	SetConstraints(constraints constraints.Constraints)
	GetSize() size.Size
}
