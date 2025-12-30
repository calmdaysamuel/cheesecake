package text

import (
	"cheesecake/canvas"
	"cheesecake/constraints"
	"cheesecake/size"
	"cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	parentWidget         *Model
	renderObjectChildren []widget.RenderElement
	constraints.Constraints
	size.Size
	renderText string
	ID         string
	canvas     canvas.Canvas
}

func (e *Element) Dispose() {}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) SetConstraints(constraints constraints.Constraints) {
	e.Constraints = constraints
	e.MaxWidth = max(e.MaxWidth, 0)
	e.MaxHeight = max(e.MaxHeight, 0)
	totalCanvas := make(canvas.Canvas, 0)
	for _, s := range strings.Split(e.parentWidget.Text, "\n") {
		cells := make([]canvas.Cell, 0)
		for _, char := range s {
			cell := canvas.Cell{
				Runes: []rune{char},
			}
			if lip, ok := e.parentWidget.Style.GetBackground().(lipgloss.Color); ok {
				cell.BgColor = lip
			}
			if lip, ok := e.parentWidget.Style.GetForeground().(lipgloss.Color); ok {
				cell.FgColor = lip
			}
			cell.Bold = e.parentWidget.Style.GetBold()
			cell.Faint = e.parentWidget.Style.GetFaint()
			cell.Italic = e.parentWidget.Style.GetItalic()
			cell.Underline = e.parentWidget.Style.GetUnderline()
			cell.UnderlineSpaces = e.parentWidget.Style.GetUnderlineSpaces()
			cells = append(cells, cell)

		}

		c := canvas.Partition(cells, e.MaxWidth)
		c = canvas.MergeTopLeft(c)
		totalCanvas = append(totalCanvas, c...)
	}
	e.canvas = totalCanvas
	e.Width, e.Height = canvas.Size(e.canvas)
}

func (e *Element) View() canvas.Canvas {
	return e.canvas
}

func (e *Element) Widget() widget.Widget {
	return e.parentWidget
}

func (e *Element) AdoptChild(child widget.RenderElement) {
	e.renderObjectChildren = append(e.renderObjectChildren, child)
}

func (e *Element) ClearChildren() {
	e.renderObjectChildren = nil
}

func (e *Element) DirectDescendants() []widget.Widget {
	return nil
}
