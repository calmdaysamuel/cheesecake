package widget

type RenderWidget interface {
	Widget
}

type Object interface {
	View() string
}
