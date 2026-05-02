package layoutbuilder

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
	werror "github.com/palantir/witchcraft-go-error"
)

var _ widget.RenderElement = &Element{}
var _ widget.Flexible = &Element{}

type Element struct {
	ID                  string
	renderChild         widget.RenderElement
	Child               widget.Widget
	LastConstraints     constraints.Constraints
	ConstraintsListener func(c constraints.Constraints) error
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

func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) {
	return []widget.Widget{e.Child}, nil
}

func (e *Element) View() canvas.Canvas {
	if e.renderChild != nil {
		return e.renderChild.View()
	}
	return canvas.New(0, 0)
}

func (e *Element) SetConstraints(ctx context.Context, c constraints.Constraints) error {
	_ = e.ConstraintsListener(c)
	if err := e.renderChild.SetConstraints(ctx, c); err != nil {
		return werror.WrapWithContextParams(ctx, err, "child in layout-builder element fail to set constraint")
	}
	e.Size = e.renderChild.GetSize()
	return nil
}

func (e *Element) Flex() (horizontal, vertical int) {
	return 1, 1
}
