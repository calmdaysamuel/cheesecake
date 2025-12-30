package preferred

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
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
	horizontal, vertical := 1, 1
	if e.parent.PreferredHeight > 0 {
		vertical = 0
	}

	if e.parent.PreferredWidth > 0 {
		horizontal = 0
	}
	return horizontal, vertical
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
	return []widget.Widget{e.parent.Child}
}

func (e *Element) View() canvas.Canvas {
	return canvas.MergeTopLeft(canvas.New(e.Width, e.Height), e.renderObject.View())
}

func (e *Element) SetConstraints(constraints constraints.Constraints) {
	if e.parent.PreferredWidth > 0 {
		constraints.MaxWidth = min(e.parent.PreferredWidth, constraints.MaxWidth)
	}

	if e.parent.PreferredHeight > 0 {
		constraints.MaxHeight = min(e.parent.PreferredHeight, constraints.MaxHeight)
	}
	e.Constraints = constraints
	e.renderObject.SetConstraints(e.Constraints)

	e.Size = size.Size{
		Height: e.MaxHeight,
		Width:  e.MaxWidth,
	}
}
