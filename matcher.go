package pattern

import (
	"reflect"
)

type Handler[T any] func() T

type Matcher[T any] struct {
	value    interface{}
	matched  bool
	response T
}

type NotPattern struct {
	pattern interface{}
}

func NewMatcher[T any](value any) *Matcher[T] {
	return &Matcher[T]{value: value}
}

func (m *Matcher[T]) With(pattern interface{}, fn Handler[T]) *Matcher[T] {
	if m.matched {
		return m
	}

	switch p := pattern.(type) {
	case NotPattern:
		if !reflect.DeepEqual(m.value, p.pattern) {
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

func (m *Matcher[T]) WithString(fn Handler[T]) *Matcher[T] {
	if m.matched {
		return m
	}

	if _, ok := m.value.(string); ok {
		m.response = fn()
		m.matched = true
	}
	return m
}

func (m *Matcher[T]) Otherwise(fn Handler[T]) T {
	if !m.matched {
		m.response = fn()
	}
	return m.response
}

func Not(pattern interface{}) NotPattern {
	return NotPattern{pattern: pattern}
}
