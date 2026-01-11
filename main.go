package main

import (
	"github.com/calmdaysamuel/cheesecake/application"
	"github.com/calmdaysamuel/cheesecake/widgets/text"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	_ = application.Start(&text.InteractiveModel{
		Text:         "Hello Samuel",
		DefaultStyle: lipgloss.NewStyle().Background(lipgloss.Color("5")),
		HoverStyle:   lipgloss.NewStyle().Background(lipgloss.Color("6")),
		ActiveStyle:  lipgloss.NewStyle().Background(lipgloss.Color("7")),
	})
}
