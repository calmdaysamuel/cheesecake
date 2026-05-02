package flex

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/widgets/utils"
	werror "github.com/palantir/witchcraft-go-error"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	ID                   string
	Options              Options
	Children             []widget.Widget
	renderObjectChildren []widget.RenderElement

	FlexProvider        func(widget.Flexible) int
	SizeProvider        func(size.Size) int
	ConstraintsProvider func(constraints.Constraints) int
	ChildJoiner         func(...canvas.Canvas) canvas.Canvas
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

func (e *Element) View(ctx context.Context, box constraints.Constraints) (canvas.Canvas, size.Size, error) {
	totalFlex := 0.0
	for _, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			flexValue := e.FlexProvider(flex)
			if flexValue <= 0 {
				return canvas.Canvas{}, size.Size{}, werror.ErrorWithContextParams(ctx, "zero or negative flex value for flexible widget")
			}
			totalFlex += max(float64(flexValue), 1.0)
		}
	}

	remainingSize := e.ConstraintsProvider(box)
	childrenViews := make([]canvas.Canvas, len(e.renderObjectChildren))
	for i, child := range e.renderObjectChildren {
		if _, ok := child.(widget.Flexible); ok {
			continue
		}
		childView, childSize, err := child.View(ctx, constraints.Max)
		if err != nil {
			return canvas.Canvas{}, size.Size{}, err
		}
		remainingSize -= e.SizeProvider(childSize)
		childrenViews[i] = childView
	}

	for i, child := range e.renderObjectChildren {
		if flex, ok := child.(widget.Flexible); ok {
			tempConstraints := box
			tempConstraints.MaxWidth = int(float64(e.FlexProvider(flex)) / totalFlex * float64(remainingSize))

			childView, _, err := child.View(ctx, tempConstraints)
			if err != nil {
				return canvas.Canvas{}, size.Size{}, err
			}
			childrenViews[i] = childView
		}
	}
	sumCanvas := e.ChildJoiner(childrenViews...)
	size := canvas.Size(sumCanvas)
	canvasWithBackground := canvas.MergeTopLeft(canvas.NewWithCell(
		size.Width,
		size.Height,
		canvas.DefaultCellWithBgColor(
			string(e.Options.BackgroundColor))), sumCanvas)
	return utils.ViolatingConstraintsCanvas(ctx, box, canvasWithBackground)
}
