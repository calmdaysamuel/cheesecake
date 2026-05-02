package padding

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	size.Size
	renderObject widget.RenderElement
	ID           string
	child        widget.Widget
	options      Options
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) Dispose() {}

func (e *Element) AdoptChild(child widget.RenderElement) {
	e.renderObject = child
}

func (e *Element) ClearChildren() {
	e.renderObject = nil
}

func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) {
	if e.child == nil {
		return nil, nil
	}
	return []widget.Widget{e.child}, nil
}

func (e *Element) View() canvas.Canvas {
	c := canvas.MergeCenter(e.renderObject.View())
	c = canvas.AddLeft(c, e.options.LeftPadding, canvas.DefaultCellWithBgColor(string(e.options.BackgroundColor)))
	c = canvas.AddRight(c, e.options.RightPadding, canvas.DefaultCellWithBgColor(string(e.options.BackgroundColor)))
	c = canvas.AddBottom(c, e.options.BottomPadding, canvas.DefaultCellWithBgColor(string(e.options.BackgroundColor)))
	c = canvas.AddTop(c, e.options.TopPadding, canvas.DefaultCellWithBgColor(string(e.options.BackgroundColor)))
	return c
}

func (e *Element) SetConstraints(ctx context.Context, c constraints.Constraints) error {
	top, right, bottom, left := e.options.TopPadding, e.options.RightPadding, e.options.BottomPadding, e.options.LeftPadding
	childConstraints := constraints.Constraints{
		MaxHeight: c.MaxHeight - top - bottom,
		MaxWidth:  c.MaxWidth - right - left,
	}
	if err := e.renderObject.SetConstraints(ctx, childConstraints); err != nil {
		return err
	}
	childSize := e.renderObject.GetSize()
	e.Width, e.Height = childSize.Width+right+left, childSize.Height+top+bottom
	return nil
}
