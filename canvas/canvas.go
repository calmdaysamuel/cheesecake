package canvas

import (
	"github.com/charmbracelet/lipgloss"
	"slices"
	"strings"
)

type Canvas []Row

func Size(canvas Canvas) (width, height int) {
	height = len(canvas)
	if height > 0 {
		width = len(canvas[0])
	}
	return width, height
}

func MergeCenter(canvases ...Canvas) Canvas {
	return Merge(lipgloss.Center, lipgloss.Center, canvases...)
}

func MergeTopLeft(canvases ...Canvas) Canvas {
	return Merge(lipgloss.Top, lipgloss.Left, canvases...)
}

func MergeBottomLeft(canvases ...Canvas) Canvas {
	return Merge(lipgloss.Bottom, lipgloss.Left, canvases...)
}
func MergeTopRight(canvases ...Canvas) Canvas {
	return Merge(lipgloss.Top, lipgloss.Right, canvases...)
}

func MergeBottomRight(canvases ...Canvas) Canvas {
	return Merge(lipgloss.Bottom, lipgloss.Right, canvases...)
}

func Truncate(canvas Canvas, width, height int) Canvas {
	width = max(width, 0)
	height = max(height, 0)
	if len(canvas) > height {
		canvas = canvas[:height]
	}
	for i, row := range canvas {
		if len(row) > width {
			canvas[i] = row[:width]
		}
	}
	return canvas
}

// Merge the provided canvas assuming the first layer is at the bottom and the last layer is at the top
func Merge(verticalPosition, horizontalPosition lipgloss.Position, canvases ...Canvas) Canvas {
	maxWidth := 0
	maxHeight := 0
	for _, canvas := range canvases {
		maxHeight = max(maxHeight, len(canvas))
		for _, row := range canvas {
			maxWidth = max(maxWidth, len(row))
		}
	}
	for i, canvas := range canvases {
		c := PlaceHorizontal(canvas, horizontalPosition, maxWidth)
		c = PlaceVertical(c, verticalPosition, maxWidth, maxHeight)
		canvases[i] = c
	}

	if len(canvases) == 1 {
		return canvases[0]
	}

	mergedCanvas := New(maxWidth, maxHeight)
	if maxWidth == 0 || maxHeight == 0 {
		return mergedCanvas
	}
	for _, canvas := range canvases {
		for j, row := range canvas {
			for k, cell := range row {
				if cell.Transparent {
					continue
				}
				cellCopy := cell
				if cellCopy.BgColor == "" {
					cellCopy.BgColor = mergedCanvas[j][k].BgColor
				}
				mergedCanvas[j][k] = cellCopy
			}
		}
	}
	return mergedCanvas
}

func JoinVertical(horizontalPosition lipgloss.Position, canvases ...Canvas) Canvas {
	maxWidth := 0
	totalHeight := 0
	for _, canvas := range canvases {
		totalHeight += len(canvas)
		for _, row := range canvas {
			maxWidth = max(maxWidth, len(row))
		}
	}
	for i, canvas := range canvases {
		canvases[i] = PlaceHorizontal(canvas, horizontalPosition, maxWidth)
	}

	mergedCanvas := New(maxWidth, totalHeight)
	j := 0
	for _, canvas := range canvases {
		for _, row := range canvas {
			for k, cell := range row {
				if cell.Transparent {
					continue
				}
				mergedCanvas[j][k] = cell
			}
			j++
		}
	}
	return mergedCanvas
}

func JoinHorizontal(verticalPosition lipgloss.Position, canvases ...Canvas) Canvas {
	totalWidth := 0
	maxHeight := 0
	for _, canvas := range canvases {
		maxHeight = max(maxHeight, len(canvas))
		maxWidth := 0
		for _, row := range canvas {
			maxWidth = max(maxWidth, len(row))
		}
		totalWidth += maxWidth
	}
	for i, canvas := range canvases {
		maxWidth := 0
		for _, row := range canvas {
			maxWidth = max(maxWidth, len(row))
		}
		canvases[i] = PlaceVertical(canvas, verticalPosition, maxWidth, maxHeight)
	}

	mergedCanvas := New(totalWidth, maxHeight)
	offset := 0
	for _, canvas := range canvases {
		maxWidth := 0
		for j, row := range canvas {
			for k, cell := range row {
				maxWidth = max(len(row), maxWidth)
				if cell.Transparent {
					continue
				}
				mergedCanvas[j][k+offset] = cell
			}
		}
		offset += maxWidth
	}
	return mergedCanvas
}

func PlaceHorizontal(canvas Canvas, position lipgloss.Position, width int) Canvas {
	if position == lipgloss.Left {
		for i := 0; i < len(canvas); i++ {
			row := canvas[i]
			for j := 0; j < max(width-len(row), 0); j++ {
				row = append(row, DefaultCell())
			}
		}
	} else if position == lipgloss.Right {
		for i := 0; i < len(canvas); i++ {
			row := make(Row, 0)
			for j := 0; j < max(width-len(canvas[i]), 0); j++ {
				row = append(row, DefaultCell())
			}
			row = append(row, canvas[i]...)
			canvas[i] = row
		}
	} else if position == lipgloss.Center {
		for i := 0; i < len(canvas); i++ {
			row := make(Row, 0)
			padding := width - len(canvas[i])
			paddingLeft := padding / 2
			for j := 0; j < max(paddingLeft, 0); j++ {
				row = append(row, DefaultCell())
			}
			row = append(row, canvas[i]...)
			for j := 0; j < max(padding-paddingLeft, 0); j++ {
				row = append(row, DefaultCell())
			}
			canvas[i] = row
		}
	}
	return canvas
}
func AddTop(canvas Canvas, extraHeight int, cell Cell) Canvas {
	width := 0
	if len(canvas) > 0 {
		width = len(canvas[0])
	}
	newCanvas := make(Canvas, 0)
	for i := 0; i < extraHeight; i++ {
		row := []Cell{cell}
		newCanvas = append(newCanvas, slices.Repeat(row, width))
	}
	return append(newCanvas, canvas...)
}

func AddBottom(canvas Canvas, extraHeight int, cell Cell) Canvas {
	width := 0
	if len(canvas) > 0 {
		width = len(canvas[0])
	}
	for i := 0; i < extraHeight; i++ {
		row := []Cell{cell}
		canvas = append(canvas, slices.Repeat(row, width))
	}
	return canvas
}

func AddLeft(canvas Canvas, extraWidth int, cell Cell) Canvas {
	for i := 0; i < len(canvas); i++ {
		row := make(Row, 0)
		for j := 0; j < extraWidth; j++ {
			row = append(row, cell)
		}
		row = append(row, canvas[i]...)
		canvas[i] = row
	}
	return canvas
}

func AddRight(canvas Canvas, extraWidth int, cell Cell) Canvas {
	for i := 0; i < len(canvas); i++ {
		for j := 0; j < extraWidth; j++ {
			canvas[i] = append(canvas[i], cell)
		}
	}
	return canvas
}

func New(width, height int) Canvas {
	var canvas Canvas
	for i := 0; i < height; i++ {
		row := make(Row, width)
		for j := range row {
			row[j] = DefaultCell()
		}
		canvas = append(canvas, row)
	}
	return canvas
}

func PlaceVertical(canvas Canvas, position lipgloss.Position, width, height int) Canvas {
	if position == lipgloss.Top {
		for i := 0; i < max(0, height-len(canvas)); i++ {
			row := []Cell{DefaultCell()}
			canvas = append(canvas, slices.Repeat(row, width))
		}
		return canvas
	} else if position == lipgloss.Bottom {
		row := []Cell{DefaultCell()}
		row = slices.Repeat(row, width)
		var newCanvas Canvas
		for i := 0; i < max(0, height-len(canvas)); i++ {
			newCanvas = append(newCanvas, row)
		}
		newCanvas = append(newCanvas, canvas...)
		return newCanvas
	} else if position == lipgloss.Center {
		row := []Cell{DefaultCell()}
		row = slices.Repeat(row, width)
		var newCanvas Canvas

		padding := height - len(canvas)
		paddingTop := padding / 2
		for i := 0; i < max(0, paddingTop); i++ {
			newCanvas = append(newCanvas, row)
		}
		newCanvas = append(newCanvas, canvas...)

		for i := 0; i < max(0, padding-paddingTop); i++ {
			newCanvas = append(newCanvas, row)
		}
		return newCanvas
	}
	return canvas
}

func NewWithCell(width, height int, cell Cell) Canvas {
	var canvas Canvas
	for i := 0; i < height; i++ {
		row := make(Row, width)
		for j := range row {
			row[j] = cell
		}
		canvas = append(canvas, row)
	}
	return canvas
}

func (c Canvas) View() string {
	lines := make([]string, 0)
	for _, row := range c {
		lines = append(lines, row.View())
	}
	return strings.Join(lines, "\n")
}
