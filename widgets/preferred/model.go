package preferred

import (
	"cheesecake/random"
	"cheesecake/widget"
)

var _ widget.RenderWidget = &Model{}

type Option func(*Model)

type Model struct {
	Child           widget.Widget
	PreferredHeight int
	PreferredWidth  int
}

func (m *Model) Element() widget.Element {
	return &Element{
		parent: m,
		ID:     random.ID(),
	}
}

func Height(child widget.Widget, height int) *Model {
	return &Model{
		Child:           child,
		PreferredHeight: height,
	}
}

func Width(child widget.Widget, width int) *Model {
	return &Model{
		Child:          child,
		PreferredWidth: width,
	}
}
