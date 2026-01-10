package textfield

import (
	"time"

	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.State = &Element{}

type Element struct {
	ID             string
	Text           string
	LastTypeTime   time.Time
	CursorPosition int
	InFocus        bool
}

func (e *Element) Identifier() string {
	return e.ID
}

func (e *Element) Init() {}

func (e *Element) Dispose() {}
