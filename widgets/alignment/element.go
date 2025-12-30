package alignment

import (
	"cheesecake/canvas"
	"cheesecake/constraints"
	"cheesecake/size"
	"cheesecake/widget"
)

var _ widget.RenderElement = &Element{}
var _ widget.Flexible = &Element{}

type Element struct {
	size.Size
	constraints.Constraints
	parent       *Model
	renderObject widget.RenderElement
	ID           string
}

func (e *Element) Flex() (int, int) {
	return 1, 1
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

func (e *Element) DirectDescendants() []widget.Widget {
	if e.parent.Child == nil {
		return nil
	}
	return []widget.Widget{e.parent.Child}
}

func (e *Element) View() canvas.Canvas {
	background := canvas.NewWithCell(e.Width, e.Height, canvas.DefaultCellWithBgColor(string(e.parent.BgColor)))
	if e.renderObject == nil {
		return background
	}
	c := canvas.Merge(e.parent.VerticalAlignment, e.parent.HorizontalAlignment, background, e.renderObject.View())
	return c
}

func (e *Element) SetConstraints(c constraints.Constraints) {
	e.Constraints = c
	if e.renderObject != nil {
		e.renderObject.SetConstraints(e.Constraints)
	}
	e.Width, e.Height = c.MaxWidth, c.MaxHeight
}
