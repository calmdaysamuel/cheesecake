package text

import (
	"strings"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	parentWidget         *Model
	renderObjectChildren []widget.RenderElement
	constraints.Constraints
	size.Size
	ID     string
	canvas canvas.Canvas
	*mouseactions.Manager
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
	charIdx := 0
	for _, s := range strings.Split(e.parentWidget.Text, "\n") {
		cells := make([]canvas.Cell, 0)
		for _, char := range s {
			cell := canvas.Cell{
				Runes: []rune{char},
			}
			style := e.parentWidget.style
			if e.parentWidget.styleFunc != nil {
				cellStyle := e.parentWidget.styleFunc(charIdx, char)
				if cellStyle != nil {
					style = *cellStyle
				}
			}
			if lip, ok := style.GetBackground().(lipgloss.Color); ok {
				cell.BgColor = lip
			}
			if lip, ok := style.GetForeground().(lipgloss.Color); ok {
				cell.FgColor = lip
			}
			cell.Bold = style.GetBold()
			cell.Faint = style.GetFaint()
			cell.Italic = style.GetItalic()
			cell.Underline = style.GetUnderline()
			cell.UnderlineSpaces = style.GetUnderlineSpaces()
			cells = append(cells, cell)
			charIdx++
		}

		c := canvas.Partition(cells, e.MaxWidth)
		c = canvas.MergeTopLeft(c)
		totalCanvas = append(totalCanvas, c...)
	}

	e.canvas = canvas.AddMouseActionManager(totalCanvas, e.Manager)
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
