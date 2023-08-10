package pattern

import (
	"fmt"
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
		// fmt.Println("ttttttthere here")
		// fmt.Println(fmt.Sprintf("Value: %+v, Type: %s", p, reflect.TypeOf(p)))
		if p.Match(m.value) {
			m.patternMatched(fn)
		}
	case []AnyPatterner:
		// fmt.Println("I AM HERE")
		var allMatched = true
		patternVal := reflect.ValueOf(p)
		value := reflect.ValueOf(m.value)

		if value.Len() != patternVal.Len() {
			// fmt.Println("Break here")
			break
		}

		for i := 0; i < value.Len(); i++ {
			val := value.Index(i)

			// fmt.Println(fmt.Sprintf("Value: %+v, Type: %v", val))
			// fmt.Printf("Content of p[%d]: %+v, Type: %s\n", i, p[i], reflect.TypeOf(p[i]))

			// fmt.Printf("Content of p[%d]: %+v,\n", i, val.Interface())
			// fmt.Printf("Content of p[%d]: %+v,\n", i, !p[i].Match(val.Interface()))

			if !p[i].Match(val.Interface()) {
				allMatched = false
				break
			}
		}

		if allMatched {
			m.patternMatched(fn)
		}
	case V:
		// Handle special case where slice/array is passed in
		// Compare each ith element against the ith pattern
		if reflect.TypeOf(p).Kind() == reflect.Array || reflect.TypeOf(p).Kind() == reflect.Slice {
			// fmt.Println("I AM HERE")
			patternVal := reflect.ValueOf(p)
			value := reflect.ValueOf(m.value)
			// fmt.Println(fmt.Sprintf("%+v type", reflect.TypeOf(p).Kind()))
			// fmt.Println(fmt.Sprintf("%+v and %+v", p, m.value))
			// fmt.Println(fmt.Sprintf("%+v", patternVal.Len()))

			if value.Len() != patternVal.Len() {
				fmt.Println("Break here")
				break
			}

			var allMatched = true
			for i := 0; i < patternVal.Len(); i++ {

				firstVal := patternVal.Index(i)
				secondVal := value.Index(i)

				if firstVal.Type().Implements(reflect.TypeOf((*AnyPatterner)(nil)).Elem()) && secondVal.Type().Implements(reflect.TypeOf((*AnyPatterner)(nil)).Elem()) {
					if firstVal.Interface().(AnyPatterner).Match(secondVal.Interface()) {
						fmt.Println("Match found between", firstVal, "and", secondVal)
					} else {
						fmt.Println("No match found between", firstVal, "and", secondVal)
					}
					continue
				}

				if !reflect.DeepEqual(firstVal.Interface(), secondVal.Interface()) {
					allMatched = false
					break
				}
			}

			if allMatched {
				m.patternMatched(fn)
			}

		}

	// if reflect.DeepEqual(m.value, p) {
	// 	m.patternMatched(fn)
	// }

	default:
		// fmt.Println(reflect.TypeOf(m.value))
		// fmt.Println(reflect.TypeOf(pattern))

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
