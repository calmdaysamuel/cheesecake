package application

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/tree"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/focus"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"slices"
	"time"
)

type FrameTick struct {
	TickTime time.Time
}

var _ tea.Model = &Program{}

type Program struct {
	Root             widget.Widget
	Tree             *tree.Node
	FrameRate        int64
	time             int64
	Ctx              context.Context
	rootRenderObject widget.RenderElement
	focusChain       []*focus.Element
	constraints      tea.WindowSizeMsg
	frameTime        time.Duration
}

func (p *Program) Init() tea.Cmd {
	_, cmd := p.FrameStep()
	return cmd
}

func (p *Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FrameTick:
		return p.FrameStep()
	case tea.WindowSizeMsg:
		p.constraints = msg
	case tea.KeyMsg:
		if msg.String() == "tab" {
			p.ForwardFocus()
		} else if msg.String() == "shift+tab" {
			p.BackwardFocus()

		}
	}
	return p, nil
}

func (p *Program) FrameStep() (tea.Model, tea.Cmd) {
	startTime := time.Now()
	defer func() { p.frameTime = time.Since(startTime) }()
	tree.Initialize(p.Ctx, p.Tree)
	p.rootRenderObject = tree.RootRenderObject(p.Tree)
	p.HandleFocus()
	p.rootRenderObject.SetConstraints(constraints.Tight(p.constraints.Width, p.constraints.Height))
	return p, TickAtFrameRate(p.FrameRate)
}

func (p *Program) HandleFocus() {
	p.focusChain = tree.Focus(p.Ctx, p.Tree)
	for _, element := range p.focusChain {
		if element.InFocus() {
			return
		}
	}
	if len(p.focusChain) > 0 {
		p.focusChain[0].GainLocus()
	}
}

func (p *Program) ForwardFocus() {
	p.focusChain = tree.Focus(p.Ctx, p.Tree)
	if len(p.focusChain) <= 0 {
		return
	}
	found := slices.IndexFunc(p.focusChain, func(element *focus.Element) bool {
		return element.InFocus()
	})
	if found != -1 {
		p.focusChain[found].LoseFocus()
	}
	found++
	if found < len(p.focusChain) {
		p.focusChain[found].GainLocus()
	} else {
		p.focusChain[0].GainLocus()
	}
}

func (p *Program) BackwardFocus() {
	p.focusChain = tree.Focus(p.Ctx, p.Tree)
	if len(p.focusChain) <= 0 {
		return
	}
	found := slices.IndexFunc(p.focusChain, func(element *focus.Element) bool {
		return element.InFocus()
	})
	if found != -1 {
		p.focusChain[found].LoseFocus()
	}
	found--
	if found >= 0 {
		p.focusChain[found].GainLocus()
	} else {
		p.focusChain[len(p.focusChain)-1].GainLocus()
	}
}

func (p *Program) View() string {
	if p.rootRenderObject == nil {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			Render("application is running but there is nothing to render")
	}
	return p.rootRenderObject.View().View()
}

func TickAtFrameRate(frameRate int64) tea.Cmd {
	return tea.Tick(time.Second/time.Duration(frameRate), func(time.Time) tea.Msg {
		return FrameTick{TickTime: time.Now()}
	})
}

func NewProgram(ctx context.Context, root widget.Widget) *Program {
	return &Program{
		Ctx:  ctx,
		Root: root,
		Tree: &tree.Node{
			W: root,
		},
		FrameRate: 30,
	}
}
