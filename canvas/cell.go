package canvas

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/position"
	"github.com/calmdaysamuel/cheesecake/state"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MouseOptions struct {
	OnKeyDown          func(ctx context.Context, button tea.MouseButton, X, Y int) error
	OnHoverStateChange func(ctx context.Context, b bool) error
	OnDragChange       func(ctx context.Context, from, to position.Position) error
	OnKeyUp            func(ctx context.Context, button tea.MouseButton, X int, Y int) error
}

type CellRelativePosition struct {
	X, Y             int
	GlobalX, GlobalY int
	State            state.State
	Region           MouseOptions
}

type Cell struct {
	BgColor             lipgloss.Color
	FgColor             lipgloss.Color
	Runes               []rune
	Transparent         bool
	Italic              bool
	Faint               bool
	Bold                bool
	Underline           bool
	UnderlineSpaces     bool
	RelativePositioning map[string]CellRelativePosition
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
