package main

import (
	"context"
	"fmt"
	"github.com/calmdaysamuel/cheesecake/application"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/column"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	t := tea.NewProgram(application.NewProgram(context.Background(), column.New(
		[]widget.Widget{})), tea.WithAltScreen(), tea.WithoutCatchPanics())
	if _, err := t.Run(); err != nil {
		fmt.Println("Argh:", err)
		panic(err)
	}
}
