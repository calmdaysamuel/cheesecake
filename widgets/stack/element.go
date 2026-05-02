package stack

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	ID       string
	Options  Options
	Children []widget.Widget

	renderObjectChildren []widget.RenderElement
	constraints.Constraints
	size.Size
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

func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) {
	return e.Children, nil
}

func (e *Element) View() canvas.Canvas {
	childrenViews := []canvas.Canvas{canvas.NewWithCell(e.Width, e.Height, canvas.DefaultCellWithBgColor(string(e.Options.BackgroundColor)))}
	for _, child := range e.renderObjectChildren {
		childrenViews = append(childrenViews, child.View())
	}
	return canvas.MergeWithMouseManager(lipgloss.Position(e.Options.MainAxisAlignment), lipgloss.Position(e.Options.CrossAxisAlignment), false, childrenViews...)
}

func (e *Element) SetConstraints(ctx context.Context, constraints constraints.Constraints) error {
	e.Constraints = constraints
	for _, child := range e.renderObjectChildren {
		if err := child.SetConstraints(ctx, constraints); err != nil {
			return err
		}
	}

	e.Width, e.Height = 0, 0
	for _, child := range e.renderObjectChildren {
		childSize := child.GetSize()
		e.Width = max(childSize.Width, e.Width)
		e.Height = max(childSize.Height, e.Height)
	}
	return nil
}
