package layoutbuilder

import (
	"cheesecake/constraints"
	"cheesecake/size"
	"cheesecake/widget"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	constraints.Constraints
	parentWidget *Model
	child        widget.Widget
	renderObject widget.RenderElement
	ID           string
}

func (e *Element) Dispose() {
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) GetSize() size.Size {
	if e.renderObject != nil {
		return e.renderObject.GetSize()
	}
	return size.Size{}
}

func (e *Element) Widget() widget.Widget {
	return e.parentWidget
}

func (e *Element) AdoptChild(child widget.RenderElement) {
	e.renderObject = child
}

func (e *Element) ClearChildren() {
	e.renderObject = nil
}

func (e *Element) DirectDescendants() []widget.Widget {
	if e.child != nil {
		return []widget.Widget{e.child}
	}
	return nil
}

func (e *Element) View() string {
	if e.renderObject == nil {
		return ""
	}
	return e.renderObject.View()
}

func (e *Element) SetConstraints(constraints constraints.Constraints) {
	old := e.Constraints
	e.Constraints = constraints
	if e.renderObject != nil {
		e.renderObject.SetConstraints(constraints)
	}
	if old == e.Constraints {
		return
	}
	e.child = e.parentWidget.ChildFunc(constraints)
}
