package column

import (
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/mouseactions"
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
	*mouseactions.Manager
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
	return e.parentWidget.children
}

func (e *Element) View() canvas.Canvas {
	childrenViews := make([]canvas.Canvas, 0)
	for _, child := range e.renderObjectChildren {
		childrenViews = append(childrenViews, child.View())
	}
	return canvas.AddMouseActionManager(canvas.Truncate(canvas.MergeTopLeft(canvas.NewWithCell(e.Width, e.Height, canvas.DefaultCellWithBgColor(string(e.parentWidget.BgColor))), canvas.JoinVertical(e.parentWidget.mainAxisAlignment, childrenViews...)), e.MaxWidth, e.MaxHeight), e.Manager)
}

func (e *Element) SetConstraints(constraints constraints.Constraints) {
	e.Constraints = constraints
	totalFlex := 0.0
	for _, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			_, vertical := flex.Flex()
			if vertical <= 0 {
				continue
			}
			totalFlex += max(float64(vertical), 1.0)
		}
	}

	remainingHeight := constraints.MaxHeight
	for _, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			_, vertical := flex.Flex()
			if vertical > 0 {
				continue
			}
		}
		child.SetConstraints(constraints)
		remainingHeight -= child.GetSize().Height
	}

	for _, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			_, vertical := flex.Flex()
			if vertical > 0 {
				cnst := constraints
				cnst.MaxHeight = int((float64(vertical) / totalFlex) * float64(remainingHeight))
				child.SetConstraints(cnst)
			}
		}
	}

	e.Width, e.Height = 0, 0
	for _, child := range e.renderObjectChildren {
		childSize := child.GetSize()
		e.Width = max(childSize.Width, e.Width)
		e.Height += childSize.Height
	}
}
