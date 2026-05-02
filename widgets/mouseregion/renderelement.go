package mouseregion

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}
var _ widget.Flexible = &Element{}

type Element struct {
	ID          string
	renderChild widget.RenderElement
	Child       widget.Widget
	size.Size
	mouseRegionState state.State
	options          Options
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) AdoptChild(child widget.RenderElement) {
	e.renderChild = child
}

func (e *Element) ClearChildren() {
	e.renderChild = nil
}

func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) {
	return []widget.Widget{e.Child}, nil
}

func (e *Element) View(ctx context.Context, box constraints.Constraints) (canvas.Canvas, size.Size, error) {
	childCanvas, childSize, err := e.renderChild.View(ctx, box)
	if err != nil {
		return nil, size.Size{}, err
	}
	childCanvas = canvas.UpdateRelativePosition(childCanvas, e.mouseRegionState, canvas.MouseOptions{
		OnKeyDown:          e.options.OnKeyDown,
		OnKeyUp:            e.options.OnKeyUp,
		OnHoverStateChange: e.options.OnHoverStateChange,
		OnDragChange:       e.options.OnDragChange,
	})
	return childCanvas, childSize, nil
}

func (e *Element) Flex() (horizontal, vertical int) {
	return 1, 1
}
