package main

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/application"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/focus"
	"github.com/calmdaysamuel/cheesecake/widgets/row"
	"github.com/calmdaysamuel/cheesecake/widgets/text"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	_ = application.Start(context.Background(),
		focus.NewScope(func(ctx context.Context, hasFocus bool) (widget.Widget, error) {
			return row.New([]widget.Widget{
				focus.PrimaryFocusNode(func(ctx context.Context, hasFocus bool) (widget.Widget, error) {
					return text.New("hello world\n780", text.WithTextStyle(lipgloss.NewStyle().Background(lipgloss.Color(boolToIntString(hasFocus))))), nil
				}),
				focus.PrimaryFocusNode(func(ctx context.Context, hasFocus bool) (widget.Widget, error) {
					return text.New("hello world\n780", text.WithTextStyle(lipgloss.NewStyle().Background(lipgloss.Color(boolToIntString(hasFocus))))), nil
				}),
			}), nil
		}),
	)
}

func boolToIntString(b bool) string {
	if b {
		return "5"
	}
	return "4"
}
