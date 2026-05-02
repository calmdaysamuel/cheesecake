package mouseregion

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/position"
	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
	tea "github.com/charmbracelet/bubbletea"
)

var _ widget.StatefulWidget = &Model{}

type Option func(*Options)

type Options struct {
	OnKeyDown          func(ctx context.Context, button tea.MouseButton, X, Y int) error
	OnKeyUp            func(ctx context.Context, button tea.MouseButton, X, Y int) error
	OnHoverStateChange func(ctx context.Context, b bool) error
	OnDragChange       func(ctx context.Context, from, to position.Position) error
}

type State struct {
	MouseButton                  tea.MouseButton
	MouseDownPositionStart       position.Position
	IsHovering                   bool
	MouseDownPositionStartGlobal position.Position
}
type Model struct {
	Child   widget.Widget
	Options Options
}

func (m *Model) CreateState(ctx context.Context) (s state.State, err error) {
	return state.New(State{})
}

func (m *Model) Build(ctx context.Context, widgetContext widgetcontext.Context, widgetState state.State) (widget.Widget, error) {
	return &RenderModel{
		State:   widgetState,
		Child:   m.Child,
		Options: m.Options,
	}, nil
}

func New(child widget.Widget, options ...Option) *Model {
	m := &Model{
		Child: child,
	}
	for _, option := range options {
		option(&m.Options)
	}
	return m
}

func HandleKeyDown(ctx context.Context, positions map[string]canvas.CellRelativePosition, msg tea.MouseMsg) error {
	if msg.Button == tea.MouseButtonNone {
		return nil
	}
	for _, p := range positions {
		curVal, err := state.Current[State](p.State)
		if err != nil {
			return err
		}
		curVal.MouseButton = msg.Button
		curVal.MouseDownPositionStart = position.Position{X: p.X, Y: p.Y}
		curVal.MouseDownPositionStartGlobal = position.Position{X: msg.X - 1, Y: msg.Y - 1}
		if p.Region.OnKeyDown != nil {
			if err := p.Region.OnKeyDown(ctx, msg.Button, p.X, p.Y); err != nil {
				return err
			}
		}
		if err := state.Update(p.State, curVal); err != nil {
			return err
		}
	}
	return nil
}

func HandleKeyUp(ctx context.Context, positions map[string]canvas.CellRelativePosition, msg tea.MouseMsg) error {
	if msg.Button == tea.MouseButtonNone {
		return nil
	}
	for _, p := range positions {
		curVal, err := state.Current[State](p.State)
		if err != nil {
			return err
		}
		downButton := curVal.MouseButton
		curVal.MouseButton = 0
		curVal.MouseDownPositionStart = position.Position{}
		curVal.MouseDownPositionStartGlobal = position.Position{}
		if p.Region.OnKeyUp != nil || downButton != msg.Button {
			if err := p.Region.OnKeyUp(ctx, msg.Button, p.X, p.Y); err != nil {
				return err
			}
		}
		if err := state.Update(p.State, curVal); err != nil {
			return err
		}
	}
	return nil
}

func HandleMouseMotion(ctx context.Context, positions map[string]canvas.CellRelativePosition) error {
	for _, p := range positions {
		curVal, err := state.Current[State](p.State)
		if err != nil {
			return err
		}
		if !curVal.IsHovering {
			curVal.IsHovering = true
			if p.Region.OnHoverStateChange != nil {
				if err := p.Region.OnHoverStateChange(ctx, true); err != nil {
					return err
				}
			}
		}

		//if curVal.MouseButton != 0 {
		//	curVal.CurrentMousePosition = position.Position{X: p.X, Y: p.Y}
		//	if p.Region.OnDragChange != nil {
		//		if err := p.Region.OnDragChange(ctx, curVal.MouseDownPositionStart, curVal.CurrentMousePosition); err != nil {
		//			return err
		//		}
		//	}
		//}

		if err := state.Update(p.State, curVal); err != nil {
			return err
		}
	}
	return nil
}

func HandleMouseHoverState(ctx context.Context, positions map[string]canvas.CellRelativePosition, isHovering bool) error {
	for _, p := range positions {
		curVal, err := state.Current[State](p.State)
		if err != nil {
			return err
		}
		if curVal.IsHovering == isHovering {
			continue
		}
		curVal.IsHovering = isHovering
		if p.Region.OnHoverStateChange != nil {
			if err := p.Region.OnHoverStateChange(ctx, curVal.IsHovering); err != nil {
				return err
			}
		}
		if err := state.Update(p.State, curVal); err != nil {
			return err
		}
	}
	return nil
}
