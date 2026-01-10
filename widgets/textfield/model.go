package textfield

import (
	"context"
	"time"

	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/alignment"
	"github.com/calmdaysamuel/cheesecake/widgets/focus"
	"github.com/calmdaysamuel/cheesecake/widgets/preferred"
	"github.com/calmdaysamuel/cheesecake/widgets/text"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ widget.StatefulWidget = &Model{}

type Option func(*Model)

type Theme struct {
	BackgroundColor  lipgloss.Color
	PlaceholderStyle lipgloss.Style
	PrimaryTextStyle lipgloss.Style
	CursorStyle      lipgloss.Style
}
type Model struct {
	VisibleLines    int
	Placeholder     string
	OnValueChange   func(currentValue string)
	InFocusTheme    Theme
	OutOfFocusTheme Theme
}

func (m *Model) Element() widget.State {
	return &Element{
		ID: random.ID(),
	}
}

func (m *Model) Build(ctx context.Context, element widget.State) widget.Widget {
	state := element.(*Element)
	displayText := m.Placeholder
	theme := m.OutOfFocusTheme
	if state.InFocus {
		theme = m.InFocusTheme
	}
	displayTextStyle := theme.PlaceholderStyle
	if state.Text != "" {
		displayText = state.Text
		displayTextStyle = theme.PrimaryTextStyle
		if state.InFocus {
			if len(state.Text) == state.CursorPosition {
				displayText += " "
			}
		}
	}
	return focus.New(
		preferred.Height(
			alignment.TopLeft(
				text.New(displayText, text.WithTextStyle(displayTextStyle),
					text.WithTextStyleFunc(func(idx int, char rune) *lipgloss.Style {
						if state.InFocus {
							if time.Now().UnixMilli()%1000 < 500 || state.LastTypeTime.Add(time.Second/2).After(time.Now()) {
								if idx == state.CursorPosition {
									c := theme.CursorStyle
									return &c
								}
							}
						}
						return nil
					})),
				alignment.WithBackgroundColor(theme.BackgroundColor),
			),
			m.VisibleLines,
		),
		focus.WithOnFocusGain(func() {
			state.InFocus = true
		}),
		focus.WithOnFocusLoss(func() {
			state.InFocus = false
		}),
		focus.WithOnKeyPress(func(msg tea.KeyMsg) {
			beforeCursor := state.Text[:state.CursorPosition]
			afterCursor := state.Text[state.CursorPosition:]
			kp := msg.String()
			switch {
			case kp == "left":
				state.CursorPosition = state.CursorPosition - 1
				state.CursorPosition = max(state.CursorPosition, 0)
			case kp == "right":
				state.CursorPosition = state.CursorPosition + 1
				state.CursorPosition = min(state.CursorPosition, len(state.Text))
			case kp == "backspace":
				if len(beforeCursor) == 0 {
					return
				}
				beforeCursor = beforeCursor[:len(beforeCursor)-1]
				state.CursorPosition = state.CursorPosition - 1
			case len(kp) == 1:
				beforeCursor += msg.String()
				state.CursorPosition = state.CursorPosition + 1
			}
			state.LastTypeTime = time.Now()
			state.Text = beforeCursor + afterCursor
			if m.OnValueChange != nil {
				m.OnValueChange(state.Text)
			}
		}),
	)
}

func New(options ...Option) *Model {
	m := &Model{
		VisibleLines: 1,
		Placeholder:  "Type something...",
	}
	for _, option := range options {
		option(m)
	}
	return m
}
func WithPlaceholderText(placeholder string) Option {
	return func(model *Model) {
		model.Placeholder = placeholder
	}
}

func WithInFocusTheme(c Theme) Option {
	return func(model *Model) {
		model.InFocusTheme = c
	}
}

func WithOutOfFocusTheme(c Theme) Option {
	return func(model *Model) {
		model.OutOfFocusTheme = c
	}
}

func WithOnValueChange(onValueChange func(currentValue string)) Option {
	return func(model *Model) {
		model.OnValueChange = onValueChange
	}
}
