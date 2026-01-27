package widget

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
)

type StatefulWidget interface {
	Widget
	CreateState(ctx context.Context) (state.State, error)
	Build(ctx context.Context, widgetContext widgetcontext.Context, widgetState state.State) (Widget, error)
}
