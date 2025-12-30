package widget

import (
	"context"
)

type StatefulWidget interface {
	Widget
	Build(ctx context.Context, element State) Widget
}
