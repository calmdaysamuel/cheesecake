package border

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
	size.Size
	constraints.Constraints
	renderObject widget.RenderElement
	ID           string
	child        widget.Widget
	options      Options
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

func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) {
	if e.child == nil {
		return nil, nil
	}
	return []widget.Widget{e.child}, nil
}

func (e *Element) View() canvas.Canvas {
	top, right, bottom, left := e.options.EnableTopBorder, e.options.EnableRightBorder, e.options.EnableBottomBorder, e.options.EnableLeftBorder
	c := canvas.New(e.Width, e.Height)
	if len(c) == 0 || len(c[0]) == 0 {
		return c
	}
	if top {
		for i := range c[0] {
			cell := canvas.Cell{
				Runes:   []rune(e.options.BorderStyle.Top),
				FgColor: e.options.BorderColor,
				BgColor: e.options.BackgroundColor,
			}
			if i == 0 {
				cell.Runes = []rune(e.options.BorderStyle.TopLeft)
			} else if i == len(c[0])-1 {
				cell.Runes = []rune(e.options.BorderStyle.TopRight)
			}
			c[0][i] = cell
		}
	}

	if bottom {
		for i := range c[len(c)-1] {
			cell := canvas.Cell{
				Runes:   []rune(e.options.BorderStyle.Bottom),
				FgColor: e.options.BorderColor,
				BgColor: e.options.BackgroundColor,
			}
			if i == 0 {
				cell.Runes = []rune(e.options.BorderStyle.BottomLeft)
			} else if i == len(c[0])-1 {
				cell.Runes = []rune(e.options.BorderStyle.BottomRight)
			}
			c[len(c)-1][i] = cell
		}
	}

	if left {
		for i := 1; i < len(c)-1; i++ {
			cell := canvas.Cell{
				Runes:   []rune(e.options.BorderStyle.Left),
				FgColor: e.options.BorderColor,
				BgColor: e.options.BackgroundColor,
			}
			c[i][0] = cell
		}
	}

	if right {
		for i := 1; i < len(c)-1; i++ {
			cell := canvas.Cell{
				Runes:   []rune(e.options.BorderStyle.Right),
				FgColor: e.options.BorderColor,
				BgColor: e.options.BackgroundColor,
			}
			c[i][len(c[i])-1] = cell
		}
	}
	return canvas.MergeCenter(c, e.renderObject.View())
}

func (e *Element) SetConstraints(ctx context.Context, c constraints.Constraints) error {
	childConstraints := constraints.Constraints{
		MaxHeight: c.MaxHeight,
		MaxWidth:  c.MaxWidth,
	}
	top, right, bottom, left := e.options.EnableTopBorder, e.options.EnableRightBorder, e.options.EnableBottomBorder, e.options.EnableLeftBorder
	if top {
		childConstraints.MaxHeight -= len(e.options.BorderStyle.Top)
	}
	if bottom {
		childConstraints.MaxHeight -= len(e.options.BorderStyle.Bottom)
	}

	if right {
		childConstraints.MaxWidth -= len(e.options.BorderStyle.Right)
	}
	if left {
		childConstraints.MaxWidth -= len(e.options.BorderStyle.Left)
	}
	if err := e.renderObject.SetConstraints(ctx, childConstraints); err != nil {
		return err
	}

	childSize := e.renderObject.GetSize()
	e.Width, e.Height = childSize.Width, childSize.Height

	if top {
		e.Height += lipgloss.Height(e.options.BorderStyle.Top)
	}
	if bottom {
		e.Height += lipgloss.Height(e.options.BorderStyle.Bottom)
	}

	if right {
		e.Width += lipgloss.Width(e.options.BorderStyle.Right)
	}
	if left {
		e.Width += lipgloss.Width(e.options.BorderStyle.Left)
	}
	return nil
}
