package pattern

import (
	"reflect"
)

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

func (m *Matcher[T, V]) With(pattern interface{}, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	switch p := pattern.(type) {
	case whenPattern[V]:
		if p.Match(m.value) {
			m.patternMatched(fn)
		}
	case unionPattern[V]:
		if p.Match(m.value) {
			m.patternMatched(fn)
		}
	case intersectionPattern[V]:
		if p.Match(m.value) {
			m.patternMatched(fn)
		}
	default:
		if reflect.DeepEqual(m.value, pattern) {
			m.patternMatched(fn)
		}
	}

	return m
}

func (m *Matcher[T, V]) WithString(fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	if reflect.TypeOf(m.value).Kind() == reflect.String {
		m.patternMatched(fn)
	}
	return m
}

func (m *Matcher[T, V]) Otherwise(fn Handler[T]) T {
	if !m.isMatched {
		m.response = fn()
	}
	return m.response
}
