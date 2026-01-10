package widget

import (
	"context"
)

type StatefulWidget interface {
	Widget
	Element() State
	Build(ctx context.Context, element State) Widget
}
