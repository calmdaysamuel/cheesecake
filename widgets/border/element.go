package border

import (
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
	top, right, bottom, left, _ := GetBorders(e.parent.Sides...)
	c := canvas.New(e.Width, e.Height)
	if len(c) == 0 || len(c[0]) == 0 {
		return c
	}
	if top {
		for i := range c[0] {
			cell := canvas.Cell{
				Runes:           []rune(e.parent.Border.Top),
				Italic:          e.parent.Style.GetItalic(),
				Faint:           e.parent.Style.GetFaint(),
				Bold:            e.parent.Style.GetBold(),
				Underline:       e.parent.Style.GetUnderline(),
				UnderlineSpaces: e.parent.Style.GetUnderlineSpaces(),
			}
			if color, ok := e.parent.Style.GetForeground().(lipgloss.Color); ok {
				cell.FgColor = color
			}
			if color, ok := e.parent.Style.GetBackground().(lipgloss.Color); ok {
				cell.BgColor = color
			}
			if i > 2 && i <= 2+len(e.parent.Label) {
				cell.Runes = []rune{rune(e.parent.Label[i-3])}
			}
			if i == 0 {
				cell.Runes = []rune(e.parent.Border.TopLeft)
			} else if i == len(c[0])-1 {
				cell.Runes = []rune(e.parent.Border.TopRight)
			}
			if (!left && i == 0) || (!right && i == len(c[0])-1) {
				cell.Runes = []rune(e.parent.Border.Top)
			}
			c[0][i] = cell
		}
	}

	if bottom {
		for i := range c[len(c)-1] {
			cell := canvas.Cell{
				Runes:           []rune(e.parent.Border.Bottom),
				Italic:          e.parent.Style.GetItalic(),
				Faint:           e.parent.Style.GetFaint(),
				Bold:            e.parent.Style.GetBold(),
				Underline:       e.parent.Style.GetUnderline(),
				UnderlineSpaces: e.parent.Style.GetUnderlineSpaces(),
			}
			if color, ok := e.parent.Style.GetForeground().(lipgloss.Color); ok {
				cell.FgColor = color
			}
			if color, ok := e.parent.Style.GetBackground().(lipgloss.Color); ok {
				cell.BgColor = color
			}
			if i == 0 {
				cell.Runes = []rune(e.parent.Border.BottomLeft)
			} else if i == len(c[len(c)-1])-1 {
				cell.Runes = []rune(e.parent.Border.BottomRight)
			}
			if (!left && i == 0) || (!right && i == len(c[len(c)-1])-1) {
				cell.Runes = []rune(e.parent.Border.Bottom)
			}
			c[len(c)-1][i] = cell
		}
	}

	if left {
		for i := 1; i < len(c)-1; i++ {
			cell := canvas.Cell{
				Runes:           []rune(e.parent.Border.Left),
				Italic:          e.parent.Style.GetItalic(),
				Faint:           e.parent.Style.GetFaint(),
				Bold:            e.parent.Style.GetBold(),
				Underline:       e.parent.Style.GetUnderline(),
				UnderlineSpaces: e.parent.Style.GetUnderlineSpaces(),
			}
			if color, ok := e.parent.Style.GetForeground().(lipgloss.Color); ok {
				cell.FgColor = color
			}
			if color, ok := e.parent.Style.GetBackground().(lipgloss.Color); ok {
				cell.BgColor = color
			}
			c[i][0] = cell
		}
	}

	if right {
		for i := 1; i < len(c)-1; i++ {
			cell := canvas.Cell{
				Runes:           []rune(e.parent.Border.Right),
				Italic:          e.parent.Style.GetItalic(),
				Faint:           e.parent.Style.GetFaint(),
				Bold:            e.parent.Style.GetBold(),
				Underline:       e.parent.Style.GetUnderline(),
				UnderlineSpaces: e.parent.Style.GetUnderlineSpaces(),
			}
			if color, ok := e.parent.Style.GetForeground().(lipgloss.Color); ok {
				cell.FgColor = color
			}
			if color, ok := e.parent.Style.GetBackground().(lipgloss.Color); ok {
				cell.BgColor = color
			}
			c[i][len(c[i])-1] = cell
		}
	}
	return canvas.MergeCenter(c, e.renderObject.View())
}

func (e *Element) SetConstraints(c constraints.Constraints) {
	e.Constraints = c
	childConstraints := constraints.Constraints{
		MaxHeight: e.MaxHeight,
		MaxWidth:  e.MaxWidth,
	}
	top, right, bottom, left, _ := GetBorders(e.parent.Sides...)
	if top {
		childConstraints.MaxHeight -= 1
	}
	if bottom {
		childConstraints.MaxHeight -= 1
	}

	if right {
		childConstraints.MaxWidth -= 1
	}
	if left {
		childConstraints.MaxWidth -= 1
	}
	e.renderObject.SetConstraints(childConstraints)
	childSize := e.renderObject.GetSize()
	e.Width, e.Height = childSize.Width, childSize.Height

	if top {
		e.Height += 1
	}
	if bottom {
		e.Height += 1
	}

	if right {
		e.Width += 1
	}
	if left {
		e.Width += 1
	}
	e.Width = max(e.Width, 0)
	e.Height = max(e.Height, 0)
}

func GetBorders(i ...bool) (top, right, bottom, left bool, ok bool) {
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
