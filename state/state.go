package state

import (
	"errors"

	"github.com/calmdaysamuel/cheesecake/random"
)

type State interface {
	_unimplementable()
	ID() string
}

type state struct {
	current interface{}
	dirty   bool
	id      string
}

func (s *state) ID() string {
	return s.id
}

func (s *state) _unimplementable() {}

// New creates a new state object for your widget. The default state cannot be nil.
func New[T any](val T) (State, error) {
	return &state{
		current: val,
		dirty:   true,
		id:      random.ID(),
	}, nil
}

// Current returns the current value stored in this state object
func Current[T any](s State) (val T, err error) {
	typedState, ok := s.(*state)
	if !ok {
		var noop T
		return noop, errors.New("state.CreateState: state object not recognized")
	}
	typedValue, ok := typedState.current.(T)
	if !ok {
		var noop T
		return noop, errors.New("state.CreateState: state object does not match request type")
	}
	return typedValue, nil
}

// Update updates teh value stored in the state object and marks the state object as dirty.
// Does not accept nil objects
func Update[T any](s State, newValue T) (err error) {
	typedState, ok := s.(*state)
	if !ok {
		return errors.New("state.CreateState: state object not recognized")
	}
	_, ok = typedState.current.(T)
	if !ok {
		return errors.New("state.CreateState: state object does not match request type")
	}
	typedState.current = newValue
	typedState.dirty = true
	return nil
}

// IsDirty returns true if the state object is considered dirty
func IsDirty(s State) (isDirty bool, err error) {
	typedState, ok := s.(*state)
	if !ok {
		return false, errors.New("state.CreateState: state object not recognized")
	}
	return typedState.dirty, nil
}

// Clean marks the state object as clean.
func Clean(s State) error {
	typedState, ok := s.(*state)
	if !ok {
		return errors.New("state.CreateState: state object not recognized")
	}
	typedState.dirty = true
	return nil
}

// Init initializes resources required by this state object
func Init(s State) error {
	return nil
}

// Dispose cleans up resources used by the state object
func Dispose(s State) error {
	return nil
}
