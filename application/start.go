package application

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/widget"
	tea "github.com/charmbracelet/bubbletea"
)

func Start(w widget.Widget) error {
	t := tea.NewProgram(NewProgram(context.Background(), w), tea.WithAltScreen(), tea.WithoutCatchPanics())
	if _, err := t.Run(); err != nil {
		return err
	}
	return nil
}
