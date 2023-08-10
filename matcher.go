package pattern

import (
	"reflect"
)

type AnyPatterner interface {
	Match(any) bool
}

type Handler[T any] func() T

// Matcher is a generic struct that matches a value of type V to a response of type T.
// It has three fields: value, isMatched and response.
// value is the input that needs to be matched.
// isMatched is a boolean that indicates whether a match has been found.
// response is the output that is returned when a match is found.
type Matcher[T any, V any] struct {
	value     V
	isMatched bool
	response  T
}

// NewMatcher is a function that creates a new Matcher instance.
// It takes a value of any type V and returns a pointer to a Matcher instance.
// The returned Matcher instance has its value field set to the input value and isMatched field set to false by default.
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
		if matchesPattern(m.value, p) {
			m.patternMatched(fn)
		}
		// if reflect.DeepEqual(m.value, p) {
		// 	m.patternMatched(fn)
		// }
	}

	return m
}

func (m *Matcher[T, V]) Otherwise(fn Handler[T]) T {
	if !m.isMatched {
		m.response = fn()
	}
	return m.response
}

func matchesPattern(input, pattern interface{}) bool {
	inputVal := reflect.ValueOf(input)
	patternVal := reflect.ValueOf(pattern)

	if inputVal.Type() != patternVal.Type() {
		return false
	}

	for i := 0; i < patternVal.NumField(); i++ {
		if !patternVal.Field(i).IsZero() {
			inputField := inputVal.Field(i)
			patternField := patternVal.Field(i)

			if patternField.Kind() == reflect.Struct {
				if !matchesPattern(inputField.Interface(), patternField.Interface()) {
					return false
				}
			} else {
				if inputField.Interface() != patternField.Interface() {
					return false
				}
			}
		}
	}

	return true
}
