package textfield

import (
	"context"
	"strings"

	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
	"github.com/calmdaysamuel/cheesecake/widgets/focus"
	"github.com/calmdaysamuel/cheesecake/widgets/layoutbuilder"
	"github.com/calmdaysamuel/cheesecake/widgets/row"
	"github.com/calmdaysamuel/cheesecake/widgets/stack"
	"github.com/calmdaysamuel/cheesecake/widgets/text"
	tea "github.com/charmbracelet/bubbletea"
)

var _ widget.StatefulWidget = &Model{}

type Option func(*Options)

type Options struct {
	Placeholder string
}
type State struct {
	Value          string
	CursorPosition int
	WindowStart    int
	WindowEnd      int
}
type Model struct {
	Options Options
}

func (m *Model) CreateState(ctx context.Context) (s state.State, err error) {
	return state.New(State{})
}

func (m *Model) Build(ctx context.Context, widgetContext widgetcontext.Context, widgetState state.State) (widget.Widget, error) {
	val, err := state.Current[State](widgetState)
	if err != nil {
		return nil, err
	}
	placeholder := m.Options.Placeholder
	if len(val.Value) > 0 {
		placeholder = ""
	}
	cursorPosition := val.CursorPosition
	userVal := val.Value
	if len(userVal) == cursorPosition {
		userVal += " "
	}
	cursorVal := strings.Repeat(" ", len(userVal[:cursorPosition]))
	return layoutbuilder.New(func(box constraints.Constraints) (widget.Widget, error) {
		box.MaxWidth = max(0, box.MaxWidth-1)
		windowStart, windowEnd, err := computeVisibleText(box, widgetState, max(len(userVal), len(placeholder)), cursorPosition)
		if err != nil {
			return nil, err
		}
		return focus.NewNode(func(ctx context.Context, hasFocus bool, hasPrimaryFocus bool) (widget.Widget, error) {
			widgets := []widget.Widget{
				text.New(GetSubstringInRange(placeholder, windowStart, windowEnd)),
				text.New(GetSubstringInRange(userVal, windowStart, windowEnd)),
			}
			if hasPrimaryFocus {
				blinkingCursor := GetOrDefault(userVal, cursorPosition, " ")
				if placeholder != "" {
					blinkingCursor = GetOrDefault(placeholder, cursorPosition, " ")
				}
				widgets = append(widgets, row.New([]widget.Widget{
					text.New(GetSubstringInRange(cursorVal, windowStart, windowEnd)),
					text.New(blinkingCursor, func(options *text.Options) {
						options.BackgroundColor = "5"
					}),
				}))
			}
			return stack.New(widgets), nil
		}, func(options *focus.Options) {
			options.KeyEventHandler = handleKeyEvent(widgetState)
		}), nil
	}), nil
}

func GetOrDefault(val string, position int, s string) string {
	if position < len(val) {
		return val[position : position+1]
	}
	return s
}

func computeVisibleText(
	box constraints.Constraints,
	widgetState state.State,
	textLength,
	cursorPosition int,
) (int, int, error) {
	val, err := state.Current[State](widgetState)
	if err != nil {
		return 0, 0, err
	}
	windowStart := val.WindowStart
	windowEnd := val.WindowEnd
	if windowEnd-windowStart != box.MaxWidth {
		windowStart = 0
		windowEnd = box.MaxWidth
	}
	if textLength > box.MaxWidth {
		if cursorPosition > windowEnd {
			diff := cursorPosition - windowEnd
			windowStart += diff
			windowEnd += diff
		} else if cursorPosition < windowStart {
			diff := windowStart - cursorPosition
			windowStart -= diff
			windowEnd -= diff
		}
	}

	if val.WindowStart != windowStart || val.WindowEnd != windowEnd {
		val.WindowStart = windowStart
		val.WindowEnd = windowEnd
		if err := state.Update(widgetState, val); err != nil {
			return 0, 0, err
		}
	}
	return windowStart, windowEnd, nil
}

func handleKeyEvent(widgetState state.State) func(ctx context.Context, msg tea.KeyMsg) error {
	return func(ctx context.Context, msg tea.KeyMsg) error {
		val, err := state.Current[State](widgetState)
		if err != nil {
			return err
		}
		currentValue := val.Value
		cursorPosition := val.CursorPosition
		switch msg := interface{}(msg).(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyRunes:
				currentValue = InsertAtOrAppend(currentValue, cursorPosition, string(msg.Runes))
				cursorPosition += len(msg.Runes)
			case tea.KeyLeft:
				if msg.Alt == true {
					cursorPosition = 0
				} else {
					cursorPosition--
				}
			case tea.KeyRight:
				if msg.Alt == true {
					cursorPosition = len(currentValue)
				} else {
					cursorPosition++
				}
			case tea.KeyBackspace:
				currentValue = RemoveStringAt(currentValue, cursorPosition, 1)
				cursorPosition--
			case tea.KeySpace:
				currentValue = InsertAtOrAppend(currentValue, cursorPosition, " ")
				cursorPosition += len(" ")
			}
		}
		cursorPosition = max(cursorPosition, 0)
		cursorPosition = min(cursorPosition, len(currentValue))
		val.Value = currentValue
		val.CursorPosition = cursorPosition
		return state.Update(widgetState, val)
	}
}

func InsertAtOrAppend(str string, at int, content string) string {
	if at >= len(str) {
		return str + content
	}
	return str[:at] + content + str[at:]
}

func RemoveStringAt(s string, position, length int) string {
	if position < 0 || position >= len(s) || length <= 0 {
		return s
	}
	end := position + length
	if end > len(s) {
		end = len(s)
	}
	return s[:position] + s[end:]
}

func GetSubstringInRange(s string, start, end int) string {
	if start > len(s) {
		return ""
	}
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}

func New(options ...Option) widget.Widget {
	m := &Model{}
	for _, option := range options {
		option(&m.Options)
	}
	return m
}
