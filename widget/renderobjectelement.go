package widget

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
)

type RenderElement interface {
	Element
	AdoptChild(child RenderElement)
	ClearChildren()
	DirectDescendants(ctx context.Context) ([]Widget, error)
	View() canvas.Canvas
	SetConstraints(ctx context.Context, box constraints.Constraints) error
	GetSize() size.Size
}
