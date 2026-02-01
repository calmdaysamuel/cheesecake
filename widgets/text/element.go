package text

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
	werror "github.com/palantir/witchcraft-go-error"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	size.Size
	ID     string
	canvas canvas.Canvas
	*mouseactions.Manager
	options Options
	text    string
}

func (e *Element) Dispose() {}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) SetConstraints(ctx context.Context, box constraints.Constraints) error {
	if err := box.Validate(); err != nil {
		return werror.WrapWithContextParams(ctx, err, "invalid constraints provided to text.Element", werror.SafeParam("constraints", box))
	}
	if box.IsZero() {
		e.canvas = canvas.Canvas{}
		e.Width, e.Height = 0, 0
		return nil
	}
	characterCells := make(canvas.Canvas, 0)
	currentRow := 0
	for _, character := range e.text {
		if character == '\n' {
			currentRow += 1
			characterCells = append(characterCells, make(canvas.Row, 0))
			continue
		}
		if len(characterCells) != currentRow+1 {
			characterCells = append(characterCells, make(canvas.Row, 0))
		}
		if len(characterCells[currentRow]) >= box.MaxWidth {
			if e.options.ShouldWrap {
				characterCells = append(characterCells, make(canvas.Row, 0))
				currentRow += 1
			} else {
				continue
			}
		}
		characterCells[currentRow] = append(characterCells[currentRow], canvas.Cell{
			BgColor:         e.options.BackgroundColor,
			FgColor:         e.options.ForegroundColor,
			Runes:           []rune{character},
			Italic:          e.options.Italic,
			Faint:           e.options.Faint,
			Bold:            e.options.Bold,
			Underline:       e.options.Underline,
			UnderlineSpaces: e.options.UnderlineSpaces,
			Transparent:     character == ' ' && e.options.BackgroundColor == "",
		})
	}

	characterCells = canvas.Merge(lipgloss.Top, e.options.Alignment, characterCells)
	e.canvas = canvas.AddMouseActionManager(characterCells, e.Manager)
	e.Width, e.Height = canvas.Size(e.canvas)
	return nil
}

func (e *Element) View() canvas.Canvas {

	return e.canvas
}

// AdoptChild the text element does not adopt any addition children.
// It is effectively a leaf node.
func (e *Element) AdoptChild(_ widget.RenderElement) {}

// ClearChildren the text element does not have any children to clear.
func (e *Element) ClearChildren() {}

// DirectDescendants the text element does not have any direct descendants.
func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) { return nil, nil }
