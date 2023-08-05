package pattern

import (
	"reflect"
)

type Handler[T any] func() T
type Predicate[V any] func(V) bool

type Matcher[T any, V any] struct {
	value    V
	matched  bool
	response T
}

type NotPattern struct {
	pattern interface{}
}

type WhenPattern[V any] struct {
	predicate Predicate[V]
}

func NewMatcher[T any, V any](value V) *Matcher[T, V] {
	return &Matcher[T, V]{value: value}
}

func (m *Matcher[T, V]) With(pattern interface{}, fn Handler[T]) *Matcher[T, V] {
	if m.matched {
		return m
	}

	switch p := pattern.(type) {
	case NotPattern:
		if !reflect.DeepEqual(m.value, p.pattern) {
			m.response = fn()
			m.matched = true
		}
	case WhenPattern[V]:
		if p.predicate(m.value) {
			m.response = fn()
			m.matched = true
		}
	default:
		if reflect.DeepEqual(m.value, pattern) {
			m.response = fn()
			m.matched = true
		}
	}

	return m
}

func (m *Matcher[T, V]) WithString(fn Handler[T]) *Matcher[T, V] {
	if m.matched {
		return m
	}

	if reflect.TypeOf(m.value).Kind() == reflect.String {
		m.response = fn()
		m.matched = true
	}
	return m
}

func (m *Matcher[T, V]) Otherwise(fn Handler[T]) T {
	if !m.matched {
		m.response = fn()
	}
	return m.response
}

func Not(pattern interface{}) NotPattern {
	return NotPattern{pattern: pattern}
}
