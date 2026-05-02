package text

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/widgets/utils"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/charmbracelet/lipgloss"
	werror "github.com/palantir/witchcraft-go-error"
)

var _ widget.RenderElement = &Element{}

type Element struct {
	ID      string
	options Options
	text    string
}

func (e *Element) Dispose() {}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) View(ctx context.Context, box constraints.Constraints) (canvas.Canvas, size.Size, error) {
	if err := box.Validate(); err != nil {
		return canvas.Canvas{}, size.Size{}, werror.WrapWithContextParams(ctx, err, "invalid constraints provided to text.Element", werror.SafeParam("constraints", box))
	}

	characterCells := make(canvas.Canvas, 0)
	currentRow := 0
	//gr := uniseg.NewGraphemes(e.text)
	options := e.options
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
			if options.ShouldWrap {
				characterCells = append(characterCells, make(canvas.Row, 0))
				currentRow += 1
			}
		}
		characterCells[currentRow] = append(characterCells[currentRow], canvas.Cell{
			BgColor:         options.BackgroundColor,
			FgColor:         options.ForegroundColor,
			Runes:           []rune{character},
			Italic:          options.Italic,
			Faint:           options.Faint,
			Bold:            options.Bold,
			Underline:       options.Underline,
			UnderlineSpaces: options.UnderlineSpaces,
			Transparent:     options.TransparentWhitespace && character == ' ' && options.BackgroundColor == "",
		})
	}
	characterCells = canvas.Merge(lipgloss.Top, options.Alignment, characterCells)
	return utils.ViolatingConstraintsCanvas(ctx, box, characterCells)
}

// AdoptChild the text element does not adopt any addition children.
// It is effectively a leaf node.
func (e *Element) AdoptChild(_ widget.RenderElement) {}

// ClearChildren the text element does not have any children to clear.
func (e *Element) ClearChildren() {}

// DirectDescendants the text element does not have any direct descendants.
func (e *Element) DirectDescendants(ctx context.Context) ([]widget.Widget, error) { return nil, nil }
