package widget

type Element interface {
	Identifier() string
	Dispose()
}
