package layoutbuilder

import (
	"cheesecake/constraints"
	"cheesecake/random"
	"cheesecake/widget"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	ChildFunc func(constraints constraints.Constraints) widget.Widget
}

func (m *Model) Element() widget.Element {
	return &Element{parentWidget: m,
		ID: random.ID(),
	}
}

func New(childFunc func(constraints constraints.Constraints) widget.Widget) *Model {
	return &Model{
		ChildFunc: childFunc,
	}
}
