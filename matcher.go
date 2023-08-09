package pattern

import (
	"reflect"
)

type Patterner[V any] interface {
	Match(V) bool
}

type AnyPatterner interface {
	Match(any) bool
}

type Handler[T any] func() T

type Matcher[T any, V any] struct {
	value     V
	isMatched bool
	response  T
}

func NewMatcher[T any, V any](value V) *Matcher[T, V] {
	return &Matcher[T, V]{value: value}
}

func (m *Matcher[T, V]) patternMatched(fn Handler[T]) {
	m.response = fn()
	m.isMatched = true
}

func (m *Matcher[T, V]) With(pattern any, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	switch p := pattern.(type) {
	// Matches with the AnyPatterner interface
	// For example, notPattern and stringPattern
	case AnyPatterner:
		if p.Match(m.value) {
			m.patternMatched(fn)
		}
	case V:
		if reflect.DeepEqual(m.value, p) {
			m.patternMatched(fn)
		}
	}

	return m
}

func (m *Matcher[T, V]) Otherwise(fn Handler[T]) T {
	if !m.isMatched {
		m.response = fn()
	}
	return m.response
}
