package layoutbuilder

import (
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	ChildFunc           func(constraints constraints.Constraints) widget.Widget
	LastConstraints     constraints.Constraints
	ConstraintsListener func(constraints.Constraints)
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:                  random.ID(),
		ChildFunc:           m.ChildFunc,
		LastConstraints:     m.LastConstraints,
		ConstraintsListener: m.ConstraintsListener,
	}
}
