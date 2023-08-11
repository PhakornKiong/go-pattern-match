package pattern

import (
	"reflect"
)

type Pattener interface {
	Match(any) bool
}

func Patteners(patterns ...Pattener) []Pattener {
	return patterns
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

func (m *Matcher[T, V]) WithPattern(pattern Pattener, fn Handler[T]) *Matcher[T, V] {
	if !m.isMatched && pattern.Match(m.value) {
		m.patternMatched(fn)
	}
	return m
}

func (m *Matcher[T, V]) WithPatterns(patterns []Pattener, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	var allMatched = true

	value := reflect.ValueOf(m.value)

	if value.Len() != len(patterns) {
		return m
	}

	for i := 0; i < value.Len(); i++ {
		val := value.Index(i)

		if !patterns[i].Match(val.Interface()) {
			allMatched = false
			break
		}
	}

	if allMatched {
		m.patternMatched(fn)
	}

	return m
}

func (m *Matcher[T, V]) WithValues(pattern V, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	if reflect.TypeOf(pattern).Kind() == reflect.Array || reflect.TypeOf(pattern).Kind() == reflect.Slice {
		// fmt.Println("I AM HERE")
		patternVal := reflect.ValueOf(pattern)
		value := reflect.ValueOf(m.value)

		if value.Len() != patternVal.Len() {
			return m
		}

		var allMatched = true
		for i := 0; i < patternVal.Len(); i++ {

			firstVal := patternVal.Index(i)
			secondVal := value.Index(i)

			if !reflect.DeepEqual(firstVal.Interface(), secondVal.Interface()) {
				allMatched = false
				break
			}
		}

		if allMatched {
			m.patternMatched(fn)
		}

		return m
	}

	return m
}

func (m *Matcher[T, V]) WithValue(pattern V, fn Handler[T]) *Matcher[T, V] {
	if m.isMatched {
		return m
	}

	if reflect.DeepEqual(m.value, pattern) {
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
