package nullable

import (
	"github.com/solsw/generichelper"
)

// Nullable represents any type that may have no value.
type Nullable[T any] struct {
	val *T
}

// NewNull creates a new [Nullable] without value.
func NewNull[T any]() *Nullable[T] {
	return &Nullable[T]{val: nil}
}

// New creates a new [Nullable] with value 'v'.
func New[T any](v T) *Nullable[T] {
	return &Nullable[T]{val: &v}
}

// nullable_Set is used for testing
func nullable_Set[T any](n *Nullable[T], v T) {
	n.val = &v
}

// Set sets the n's value to 'v'.
func (n *Nullable[T]) Set(v T) {
	nullable_Set(n, v)
}

// Null sets n as having no value.
func (n *Nullable[T]) Null() {
	n.val = nil
}

// nullable_Get is used for testing
func nullable_Get[T any](n *Nullable[T]) (T, bool) {
	if !n.Has() {
		return generichelper.ZeroValue[T](), false
	}
	return *n.val, true
}

// Get returns (n's value, true) if n has a value and (T's [zero value], false) if n has no value.
//
// [zero value]: https://go.dev/ref/spec#The_zero_value
func (n *Nullable[T]) Get() (T, bool) {
	return nullable_Get(n)
}

// Has reports whether n has a value.
func (n *Nullable[T]) Has() bool {
	return n.val != nil
}
