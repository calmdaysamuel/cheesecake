package widget

import "github.com/calmdaysamuel/cheesecake/random"

type StatefulElement interface {
	Element
	GetState() State
}

type State interface {
	Clean()
	Dirty() bool
}

var _ State = &StatefulElementImpl[any]{}
var _ StatefulElement = &StatefulElementImpl[any]{}

type StatefulElementImpl[T any] struct {
	dirty   bool
	current T
	id      string
}

func (s *StatefulElementImpl[T]) Identifier() string {
	return s.id
}

func (s *StatefulElementImpl[T]) Dispose() {}

func (s *StatefulElementImpl[T]) GetState() State {
	return s
}

func (s *StatefulElementImpl[T]) Clean() {
	s.dirty = false
}

func (s *StatefulElementImpl[T]) Dirty() bool {
	return s.dirty
}

func (s *StatefulElementImpl[T]) SetState(stateChangeFunc func(oldState T) T) {
	s.current = stateChangeFunc(s.current)
	s.dirty = true
}

func (s *StatefulElementImpl[T]) Current() T {
	return s.current
}

func NewStatefulElement[T any](defaultState T) *StatefulElementImpl[T] {
	return &StatefulElementImpl[T]{current: defaultState, id: random.ID()}
}
