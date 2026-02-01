package state

import (
	"errors"
	"fmt"

	"github.com/calmdaysamuel/cheesecake/random"
)

type State interface {
	_unimplementable()
	ID() string
	Type() string
}

type state struct {
	current   interface{}
	dirty     bool
	id        string
	stateType string
}

func (s *state) ID() string {
	return s.id
}

func (s *state) Type() string {
	return s.stateType
}

func (s *state) _unimplementable() {}

// New creates a new state object for your widget. The default state cannot be nil.
func New[T any](val T) (State, error) {
	return &state{
		current:   val,
		dirty:     true,
		id:        random.ID(),
		stateType: fmt.Sprintf("%T", val),
	}, nil
}

// Current returns the current value stored in this state object
func Current[T any](s State) (val T, err error) {
	typedState, ok := s.(*state)
	if !ok {
		var noop T
		return noop, errors.New("state.Current: state object not recognized")
	}
	typedValue, ok := typedState.current.(T)
	if !ok {
		var noop T
		return noop, errors.New("state.Current: state object does not match request type")
	}
	return typedValue, nil
}

// Update updates teh value stored in the state object and marks the state object as dirty.
// Does not accept nil objects
func Update[T any](s State, newValue T) (err error) {
	typedState, ok := s.(*state)
	if !ok {
		return errors.New("state.Update: state object not recognized")
	}
	_, ok = typedState.current.(T)
	if !ok {
		return errors.New("state.Update: state object does not match request type")
	}
	typedState.current = newValue
	typedState.dirty = true
	return nil
}

// IsDirty returns true if the state object is considered dirty
func IsDirty(s State) (isDirty bool, err error) {
	if s == nil {
		return false, errors.New("provided state object is nil")
	}
	typedState, ok := s.(*state)
	if !ok {
		return false, errors.New("state.IsDirty: state object not recognized")
	}
	return typedState.dirty, nil
}

// Clean marks the state object as clean.
func Clean(s State) error {
	typedState, ok := s.(*state)
	if !ok {
		return errors.New("state.Clean: state object not recognized")
	}
	typedState.dirty = false
	return nil
}

// MarkDirty marks the state object as dirty.
func MarkDirty(s State) error {
	typedState, ok := s.(*state)
	if !ok {
		return errors.New("state.MarkDirty: state object not recognized")
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
