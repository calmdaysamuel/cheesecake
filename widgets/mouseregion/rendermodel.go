package mouseregion

import (
	"github.com/calmdaysamuel/cheesecake/random"
	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widget"
)

var _ widget.RenderWidget = &RenderModel{}

type RenderModel struct {
	Child   widget.Widget
	State   state.State
	Options Options
}

func (r *RenderModel) Element() widget.RenderElement {
	return &Element{
		ID:               random.ID(),
		Child:            r.Child,
		mouseRegionState: r.State,
		options:          r.Options,
	}
}
