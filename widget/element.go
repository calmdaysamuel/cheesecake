package widget

type Element interface {
	Identifier() string
}

type State interface {
	Element
	Init()
	Dispose()
}
