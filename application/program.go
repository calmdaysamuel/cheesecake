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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	Root             widget.Widget
	Tree             *tree.Node
	FrameRate        int64
	Ctx              context.Context
	rootRenderObject widget.RenderElement
	constraints      tea.WindowSizeMsg
	oldMouseManagers []*mouseactions.Manager
	FocusChain       []*tree.FocusNode
	FocusedNode      *tree.FocusNode
	LastError        error
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
		//p.HandleMouseEvent(msg)
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
	return shouldRerender, p.rootRenderObject.SetConstraints(ctx, constraints.Tight(p.constraints.Width, p.constraints.Height))
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
	if p.rootRenderObject == nil {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			Render("application is running but there is nothing to render")
	}
	return canvas.MergeTopLeft(canvas.New(p.constraints.Width, p.constraints.Height), p.rootRenderObject.View()).View()
}

func (p *Program) HandleMouseEvent(msg tea.MouseMsg) {
	if p.rootRenderObject != nil {
		var c = p.rootRenderObject.View()
		row, column := msg.X, msg.Y
		if column >= len(c) || row >= len(c[0]) {
			for _, manager := range p.oldMouseManagers {
				manager.Reset()
			}
			p.oldMouseManagers = nil
			return
		}

		cell := c[column][row]
		currentActionManagers := make([]*mouseactions.Manager, 0)
		for _, manager := range cell.ActionManagers {
			if manager == nil {
				continue
			}
			currentActionManagers = append(currentActionManagers, manager)
		}

		for _, manager := range p.oldMouseManagers {
			if slices.Contains(currentActionManagers, manager) {
				continue
			}
			manager.Reset()
		}
		p.oldMouseManagers = currentActionManagers
		for _, manager := range currentActionManagers {
			manager.OnHoverStateChange(true)
		}
		switch msg.Action {
		case tea.MouseActionPress:
			for _, manager := range currentActionManagers {
				manager.OnMouseDown(msg.Button)
			}
		case tea.MouseActionRelease:
			for _, manager := range currentActionManagers {
				manager.OnMouseUp(msg.Button)
			}
		}
	}

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
