package spacer

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}
var _ widget.Flexible = &Element{}

type Element struct {
	ID      string
	Options Options

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

func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) { return nil, nil }

func (e *Element) View() canvas.Canvas {
	return e.spacing
}

func (e *Element) SetConstraints(ctx context.Context, constraints constraints.Constraints) error {
	e.spacing = canvas.NewWithCell(constraints.MaxWidth, constraints.MaxHeight, canvas.DefaultCellWithBgColor(string(e.Options.BackgroundColor)))
	e.Width, e.Height = canvas.Size(e.spacing)
	return nil
}
