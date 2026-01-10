package widget

type RenderWidget interface {
	Element() RenderElement
	Widget
}

type Object interface {
	View() string
}
