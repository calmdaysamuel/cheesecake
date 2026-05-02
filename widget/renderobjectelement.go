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
	View(ctx context.Context, box constraints.Constraints) (canvas.Canvas, size.Size, error)
}
