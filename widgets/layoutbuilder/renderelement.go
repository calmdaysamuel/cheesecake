package layoutbuilder

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}
var _ widget.Flexible = &Element{}

type Element struct {
	ID                  string
	renderChild         widget.RenderElement
	ChildFunc           func(constraints.Constraints) widget.Widget
	LastConstraints     constraints.Constraints
	ConstraintsListener func(constraints.Constraints)
	size.Size
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

func (e *Element) DirectDescendants() []widget.Widget {
	return []widget.Widget{e.ChildFunc(e.LastConstraints)}
}

func (e *Element) View() canvas.Canvas {
	if e.renderChild != nil {
		return e.renderChild.View()
	}
	return canvas.New(0, 0)
}

func (e *Element) SetConstraints(c constraints.Constraints) {
	e.ConstraintsListener(c)
	e.renderChild.SetConstraints(c)
	e.Size = e.renderChild.GetSize()
}

func (e *Element) Flex() (horizontal, vertical int) {
	return 1, 1
}
