package textfield

import (
	"time"

	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.State = &State{}

type State struct {
	ID             string
	Text           string
	LastTypeTime   time.Time
	CursorPosition int
	SelectionStart int
	InFocus        bool
}

func (e *State) Identifier() string {
	return e.ID
}

func (e *State) Init() {}

func (e *State) Dispose() {}
