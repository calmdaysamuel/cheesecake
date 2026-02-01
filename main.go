package main

import (
	"context"

	"github.com/calmdaysamuel/cheesecake/application"
	"github.com/calmdaysamuel/cheesecake/widgets/textfield"
)

func main() {
	_ = application.Start(context.Background(),
		textfield.New(func(options *textfield.Options) {
			options.Placeholder = "Type something..."
		}),
	)
}

func boolToIntString(b bool) string {
	if b {
		return "5"
	}
	return "4"
}
