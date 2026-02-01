package layoutbuilder

import (
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	LastConstraints     constraints.Constraints
	ConstraintsListener func(c constraints.Constraints) error
	Child               widget.Widget
}

func (m *Model) Element() widget.RenderElement {
	return &Element{
		ID:                  random.ID(),
		Child:               m.Child,
		LastConstraints:     m.LastConstraints,
		ConstraintsListener: m.ConstraintsListener,
	}
}
