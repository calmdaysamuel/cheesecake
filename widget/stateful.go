package widget

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/state"
	"github.com/calmdaysamuel/cheesecake/widgetcontext"
)

type StatefulWidget interface {
	Widget
	CreateState(ctx context.Context) (s state.State, err error)
	Build(ctx context.Context, widgetContext widgetcontext.Context, widgetState state.State) (Widget, error)
}
