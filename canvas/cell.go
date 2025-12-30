package canvas

import "github.com/charmbracelet/lipgloss"

type Cell struct {
	BgColor         lipgloss.Color
	FgColor         lipgloss.Color
	Runes           []rune
	Transparent     bool
	Italic          bool
	Faint           bool
	Bold            bool
	Underline       bool
	UnderlineSpaces bool
}

func (c Cell) View() string {
	return lipgloss.NewStyle().Background(c.BgColor).Foreground(c.FgColor).Faint(c.Faint).Bold(c.Bold).Italic(c.Italic).Underline(c.Underline).UnderlineSpaces(c.UnderlineSpaces).Render(string(c.Runes))
}

func DefaultCell() Cell {
	return Cell{
		Runes:       []rune(" "),
		Transparent: true,
	}
}

func DefaultCellWithBgColor(color string) Cell {
	if color == "" {
		return DefaultCell()
	}
	return Cell{
		BgColor: lipgloss.Color(color),
		Runes:   []rune(" "),
	}
}
