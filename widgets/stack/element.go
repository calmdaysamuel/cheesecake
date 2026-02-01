package stack

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	parentWidget         *Model
	renderObjectChildren []widget.RenderElement
	constraints.Constraints
	size.Size
	ID string
	*mouseactions.Manager
}

func (e *Element) Widget() widget.Widget {
	return e.parentWidget
}

func (e *Element) Dispose() {
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) AdoptChild(child widget.RenderElement) {
	e.renderObjectChildren = append(e.renderObjectChildren, child)
}

func (e *Element) ClearChildren() {
	e.renderObjectChildren = nil
}

func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) {
	return e.parentWidget.Children, nil
}

func (e *Element) View() canvas.Canvas {
	childrenViews := []canvas.Canvas{canvas.NewWithCell(e.Width, e.Height, canvas.DefaultCellWithBgColor(string(e.parentWidget.BgColor)))}
	for _, child := range e.renderObjectChildren {
		childrenViews = append(childrenViews, child.View())
	}
	return canvas.AddMouseActionManager(canvas.MergeWithMouseManager(e.parentWidget.VerticalAlignment, e.parentWidget.HorizontalAlignment, false, childrenViews...), e.Manager)
}

func (e *Element) SetConstraints(ctx context.Context, constraints constraints.Constraints) error {
	e.Constraints = constraints
	for _, child := range e.renderObjectChildren {
		if err := child.SetConstraints(ctx, constraints); err != nil {
			return err
		}
	}

	e.Width, e.Height = 0, 0
	for _, child := range e.renderObjectChildren {
		childSize := child.GetSize()
		e.Width = max(childSize.Width, e.Width)
		e.Height = max(childSize.Height, e.Height)
	}
	return nil
}
