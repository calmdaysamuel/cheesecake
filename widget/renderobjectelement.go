package widget

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
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
