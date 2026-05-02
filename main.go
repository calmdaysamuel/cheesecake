package main

import (
	"github.com/calmdaysamuel/cheesecake/application"
	"github.com/calmdaysamuel/cheesecake/crossaxis"
	"github.com/calmdaysamuel/cheesecake/widget"
	"github.com/calmdaysamuel/cheesecake/widgets/column"
	"github.com/calmdaysamuel/cheesecake/widgets/flex"
	"github.com/calmdaysamuel/cheesecake/widgets/row"
	"github.com/calmdaysamuel/cheesecake/widgets/text"
	"strings"
)

func main() {
	_ = application.Start(
		row.New([]widget.Widget{
			text.New(strings.Repeat("Hello World Hello World\n", 10), func(options *text.Options) {
				options.ShouldWrap = false
			}),

			text.New("How far", func(options *text.Options) {
				options.ShouldWrap = false
			}),
			column.New([]widget.Widget{
				text.New("Hello", func(options *text.Options) {
					options.ShouldWrap = false
				}),
				text.New("World", func(options *text.Options) {
					options.ShouldWrap = false
				}),
			}),
		}, func(options *flex.Options) {
			options.CrossAxisAlignment = crossaxis.Start
			options.Spacing = 2
		}),
	)
}
