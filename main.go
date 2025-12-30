package main

import (
	"cheesecake/application"
	"cheesecake/widget"
	"cheesecake/widgets/column"
	"context"
	"fmt"
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
