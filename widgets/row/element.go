package row

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	parentWidget         *Model
	renderObjectChildren []widget.RenderElement
	constraints.Constraints
	size.Size
	ID string
}

func (e *Element) Widget() widget.Widget {
	return e.parentWidget
}

func (e *Element) Dispose() {
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) AdoptChild(child widget.RenderElement) {
	e.renderObjectChildren = append(e.renderObjectChildren, child)
}

func (e *Element) ClearChildren() {
	e.renderObjectChildren = nil
}

func (e *Element) DirectDescendants() []widget.Widget {
	return e.parentWidget.Children
}

func (e *Element) View() canvas.Canvas {
	childrenViews := make([]canvas.Canvas, 0)
	for _, child := range e.renderObjectChildren {
		childrenViews = append(childrenViews, child.View())
	}
	return canvas.Truncate(canvas.JoinHorizontal(e.parentWidget.MainAxisAlignment, childrenViews...), e.MaxWidth, e.Height)
}

func (e *Element) SetConstraints(constraints constraints.Constraints) {
	e.Constraints = constraints
	totalFlex := 0.0
	for _, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			horizontal, _ := flex.Flex()
			if horizontal <= 0 {
				continue
			}
			totalFlex += max(float64(horizontal), 1.0)
		}
	}

	remainingWidth := constraints.MaxWidth
	for _, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			horizontal, _ := flex.Flex()
			if horizontal > 0 {
				continue
			}
		}
		child.SetConstraints(constraints)
		remainingWidth -= child.GetSize().Width
	}

	for _, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			horizontal, _ := flex.Flex()
			if horizontal > 0 {
				cnst := constraints
				cnst.MaxWidth = int(float64(horizontal) / totalFlex * float64(remainingWidth))
				child.SetConstraints(cnst)
			}
		}
	}

	e.Width, e.Height = 0, 0
	for _, child := range e.renderObjectChildren {
		childSize := child.GetSize()
		e.Height = max(childSize.Height, e.Height)
		e.Width += childSize.Width
	}
}
