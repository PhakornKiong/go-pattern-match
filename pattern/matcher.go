package pattern

import (
	"reflect"
)

type Patterner interface {
	Match(any) bool
}

func Patteners(patterns ...Patterner) []Patterner {
	return patterns
}

type Handler[T any] func() T

// Matcher is a generic struct that matches a value of type V to a response of type T.
// It has three fields: value, isMatched and response.
// value is the input that needs to be matched.
// isMatched is a boolean that indicates whether a match has been found.
// response is the output that is returned when a match is found.
type Matcher[T any, V any] struct {
	input     V
	isMatched bool
	response  T
}

// NewMatcher is a function that creates a new Matcher instance.
// It takes a value of any type V and returns a pointer to a Matcher instance.
// The returned Matcher instance has its value field set to the input value and isMatched field set to false by default.
func NewMatcher[T any, V any](input V) *Matcher[T, V] {
	return &Matcher[T, V]{input: input}
}

// WithPattern check if pattern matches the entire input
func (m *Matcher[T, V]) WithPattern(pattern Patterner, fn Handler[T]) *Matcher[T, V] {
	if !m.isMatched && pattern.Match(m.input) {
		m.patternMatched(fn)
	}
	return m
}

// WithPatterns check each of the patterns against the each of the input
func (m *Matcher[T, V]) WithPatterns(patterns []Patterner, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	var allMatched = true
	input := reflect.ValueOf(m.input)
	if input.Len() != len(patterns) {
		return m
	}

	for i := 0; i < input.Len(); i++ {
		inputVal := input.Index(i)

		if !patterns[i].Match(inputVal.Interface()) {
			allMatched = false
			break
		}
	}

	if allMatched {
		m.patternMatched(fn)
	}

	return m
}

// WithValues check for deep equality between each of the value  against the each of the input
func (m *Matcher[T, V]) WithValues(value any, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	if reflect.TypeOf(value).Kind() == reflect.Array || reflect.TypeOf(value).Kind() == reflect.Slice {
		patternVal := reflect.ValueOf(value)
		input := reflect.ValueOf(m.input)

		if input.Len() != patternVal.Len() {
			return m
		}

		var allMatched = true
		for i := 0; i < patternVal.Len(); i++ {
			firstVal := patternVal.Index(i)
			inputVal := input.Index(i)

			// Check if firstVal is a Patterner
			if patterner, ok := firstVal.Interface().(Patterner); ok {
				// If it is, run patterner.Match
				if !patterner.Match(inputVal.Interface()) {
					allMatched = false
					break
				}
			} else {
				// If it's not a Patterner, then run reflect.DeepEqual
				if !reflect.DeepEqual(firstVal.Interface(), inputVal.Interface()) {
					allMatched = false
					break
				}
			}
		}

		if allMatched {
			m.patternMatched(fn)
		}

		return m
	}

	return m
}

// WithValue check for deep equality between the value and the input
func (m *Matcher[T, V]) WithValue(pattern V, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	if reflect.DeepEqual(m.input, pattern) {
		m.patternMatched(fn)
	}

	return m
}

// Otherwise is called if no patterns match
func (m *Matcher[T, V]) Otherwise(fn Handler[T]) T {
	if !m.isMatched {
		m.response = fn()
	}
	return m.response
}

func (m *Matcher[T, V]) patternMatched(fn Handler[T]) {
	m.response = fn()
	m.isMatched = true
}
