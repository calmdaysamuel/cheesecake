package padding

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	size.Size
	constraints.Constraints
	parent       *Model
	renderObject widget.RenderElement
	ID           string
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
	top, right, bottom, left, _ := GetPadding(e.parent.Padding...)
	c := canvas.MergeCenter(e.renderObject.View())
	c = canvas.AddLeft(c, left, canvas.DefaultCellWithBgColor(string(e.parent.BgColor)))
	c = canvas.AddRight(c, right, canvas.DefaultCellWithBgColor(string(e.parent.BgColor)))
	c = canvas.AddBottom(c, bottom, canvas.DefaultCellWithBgColor(string(e.parent.BgColor)))
	c = canvas.AddTop(c, top, canvas.DefaultCellWithBgColor(string(e.parent.BgColor)))
	return c
}

func (e *Element) SetConstraints(c constraints.Constraints) {
	e.Constraints = c
	top, right, bottom, left, _ := GetPadding(e.parent.Padding...)
	childConstraints := constraints.Constraints{
		MaxHeight: e.MaxHeight - top - bottom,
		MaxWidth:  e.MaxWidth - right - left,
	}
	e.renderObject.SetConstraints(childConstraints)
	childSize := e.renderObject.GetSize()
	e.Width, e.Height = childSize.Width+right+left, childSize.Height+top+bottom
}

func GetPadding(i ...int) (top, right, bottom, left int, ok bool) {
	switch len(i) {
	case 1:
		top = i[0]
		bottom = i[0]
		left = i[0]
		right = i[0]
		ok = true
	case 2: //nolint:mnd
		top = i[0]
		bottom = i[0]
		left = i[1]
		right = i[1]
		ok = true
	case 3: //nolint:mnd
		top = i[0]
		left = i[1]
		right = i[1]
		bottom = i[2]
		ok = true
	case 4: //nolint:mnd
		top = i[0]
		right = i[1]
		bottom = i[2]
		left = i[3]
		ok = true
	}
	return top, right, bottom, left, ok
}
