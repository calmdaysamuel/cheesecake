package application

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/mouseactions"
	"github.com/calmdaysamuel/cheesecake/tree"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/mouseregion"
	tea "github.com/charmbracelet/bubbletea"
	werror "github.com/palantir/witchcraft-go-error"
	wparams "github.com/palantir/witchcraft-go-params"
)

type Options struct {
	FrameRate int
}

type RerenderMsg struct {
	TickTime time.Time
}

var _ tea.Model = &Program{}

type Program struct {
	Root              widget.Widget
	Tree              *tree.Node
	FrameRate         int64
	Ctx               context.Context
	rootRenderObject  widget.RenderElement
	constraints       tea.WindowSizeMsg
	oldMouseManagers  []*mouseactions.Manager
	FocusChain        []*tree.FocusNode
	FocusedNode       *tree.FocusNode
	LastError         error
	ActiveMouseRegion map[string]canvas.CellRelativePosition
}

func (p *Program) Init() tea.Cmd {
	return nil
}

func (p *Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	ctx := p.Ctx
	ctx = wparams.ContextWithSafeParam(ctx, "receivedMessage", fmt.Sprintf("%T", msg))
	switch msg := msg.(type) {
	case RerenderMsg:
	case tea.WindowSizeMsg:
		p.constraints = msg
		_, _ = p.FrameStep()
		return p, nil
	case tea.KeyMsg:
		if msg.String() == "tab" {
			p.FocusedNode, p.LastError = tree.FocusTreeStep(ctx, p.FocusChain, p.FocusedNode, false, tree.FocusMoveForward)
			if p.LastError != nil {
				return p, tea.Quit
			}
		} else if msg.String() == "shift+tab" {
			p.FocusedNode, p.LastError = tree.FocusTreeStep(ctx, p.FocusChain, p.FocusedNode, false, tree.FocusMoveBackward)
			if p.LastError != nil {
				return p, tea.Quit
			}
		} else {
			if p.FocusedNode != nil && p.FocusedNode.KeyEventHandler != nil {
				if err := p.FocusedNode.KeyEventHandler(ctx, msg); err != nil {
					p.LastError = err
					return p, tea.Quit
				}
			}
		}
	case tea.MouseMsg:
		_ = p.HandleMouseEvent(ctx, msg)
	}
	return p, nil
}

func (p *Program) FrameStep() (shouldRerender bool, err error) {
	ctx := p.Ctx
	shouldRerender, err = tree.RefreshWidgetTree(ctx, p.Tree, false)
	if err != nil {
		return shouldRerender, err
	}
	if p.rootRenderObject, err = tree.RootRenderObject(ctx, p.Tree); err != nil {
		return shouldRerender, err
	}

	if err := p.HandleFocus(ctx); err != nil {
		return shouldRerender, err
	}
	return shouldRerender, nil
}

func (p *Program) HandleFocus(ctx context.Context) error {
	p.FocusChain = tree.RefreshFocusTree(ctx, p.Tree, nil)
	if p.FocusedNode == nil {
		var err error
		if p.FocusedNode, err = tree.FocusTreeStep(ctx, p.FocusChain, nil, true, tree.FocusMoveForward); err != nil {
			return p.LastError
		}
	}
	return nil
}

func (p *Program) View() string {
	c, _ := p.CurrentAppCanvas(p.Ctx)
	return c.View()
}

func (p *Program) HandleMouseEvent(ctx context.Context, msg tea.MouseMsg) error {
	c, err := p.CurrentAppCanvas(ctx)
	if err != nil {
		return err
	}
	row, column := msg.X, msg.Y
	if column >= len(c) || row >= len(c[0]) || column < 0 || row < 0 {
		return nil
	}

	if err := p.handleMouseHovering(ctx, c, column, row); err != nil {
		return err
	}
	switch msg.Action {
	case tea.MouseActionPress:
		p.ActiveMouseRegion = c[column][row].RelativePositioning
		if err := mouseregion.HandleKeyDown(ctx, p.ActiveMouseRegion, msg); err != nil {
			return err
		}
	case tea.MouseActionRelease:
		if err := mouseregion.HandleKeyUp(ctx, c[column][row].RelativePositioning, msg); err != nil {
			return err
		}
		p.ActiveMouseRegion = nil
	}
	return nil
}

func (p *Program) CurrentAppCanvas(ctx context.Context) (canvas.Canvas, error) {
	if p.rootRenderObject == nil {
		return canvas.MergeTopLeft(canvas.New(p.constraints.Width, p.constraints.Height)), nil
	}
	c, _, err := p.rootRenderObject.View(ctx, constraints.Constraints{
		MaxHeight: p.constraints.Height,
		MaxWidth:  p.constraints.Width,
	})
	if err != nil {
		return canvas.Canvas{}, err
	}
	return canvas.MergeTopLeft(canvas.New(p.constraints.Width, p.constraints.Height), c), nil
}

func (p *Program) handleMouseHovering(ctx context.Context, c canvas.Canvas, column int, row int) error {
	// Mark the current cell as hovering
	hoverstates := make([]string, 0)
	for s := range c[column][row].RelativePositioning {
		hoverstates = append(hoverstates, s)
	}
	if err := mouseregion.HandleMouseHoverState(ctx, c[column][row].RelativePositioning, true); err != nil {
		return err
	}
	// Only the cell being hovered on can be in a hover state
	for i := range c {
		for j, cell := range c[i] {
			if i == column && j == row {
				continue
			}
			notHovering := make(map[string]canvas.CellRelativePosition)
			for s, position := range cell.RelativePositioning {
				if !slices.Contains(hoverstates, s) {
					notHovering[s] = position
				}
			}
			if err := mouseregion.HandleMouseHoverState(ctx, notHovering, false); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewProgram(ctx context.Context, root widget.Widget) (*Program, error) {
	rootNode, err := tree.NodeFromWidget(ctx, root)
	if err != nil {
		return nil, werror.WrapWithContextParams(ctx, err, "failed to create a new program.")
	}
	return &Program{
		Ctx:       ctx,
		Root:      root,
		Tree:      rootNode,
		FrameRate: 30,
	}, nil
}
