package widgetcontext

type Context interface {
	_unimplementable()
}

type context struct{}

func (c *context) _unimplementable() {}

func New() Context {
	return &context{}
}
