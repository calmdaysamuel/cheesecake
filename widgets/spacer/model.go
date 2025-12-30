package spacer

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderWidget = &Model{}

type Model struct {
	flex int
}

func (m *Model) Element() widget.Element {
	return &Element{
		ID:   random.ID(),
		flex: m.flex,
	}
}

func New(flex int) *Model {
	return &Model{flex: flex}
}
