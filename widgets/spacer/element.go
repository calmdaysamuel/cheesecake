package spacer

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}
var _ widget.Flexible = &Element{}

type Element struct {
	ID      string
	spacing canvas.Canvas
	flex    int
	size.Size
}

func (e *Element) Flex() (int, int) {
	return e.flex, e.flex
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) Dispose() {
}

func (e *Element) AdoptChild(child widget.RenderElement) {}

func (e *Element) ClearChildren() {}

func (e *Element) DirectDescendants() []widget.Widget { return nil }

func (e *Element) View() canvas.Canvas {
	return e.spacing
}

func (e *Element) SetConstraints(constraints constraints.Constraints) {
	e.spacing = canvas.New(constraints.MaxWidth, constraints.MaxHeight)
	e.Width, e.Height = canvas.Size(e.spacing)
}
