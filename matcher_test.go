package pattern

import (
	"math/big"
	"testing"
)

func TestMatcherStruct(t *testing.T) {

	type MyStruct struct {
		x int
		y int
	}

	type OtherStruct struct {
		z int
	}

	expected := OtherStruct{25}

	input := MyStruct{25, 35}
	output := NewMatcher[OtherStruct](input).
		With(MyStruct{25, 35}, func() OtherStruct { return expected }).
		With(MyStruct{x: 25}, func() OtherStruct { return OtherStruct{} }).
		Otherwise(func() OtherStruct { return OtherStruct{} })

	if output != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, output)
	}
}

func TestMatcher(t *testing.T) {
	input := 2
	output := NewMatcher[string](input).
		With(2, func() string { return "number: two" }).
		With(true, func() string { return "boolean: true" }).
		With("hello", func() string { return "string: hello" }).
		With(nil, func() string { return "null" }).
		With(big.NewInt(20), func() string { return "bigint: 20n" }).
		Otherwise(func() string { return "something else" })

	expected := "number: two"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestMatcher2(t *testing.T) {
	input := 100
	output := NewMatcher[string](input).
		With(2, func() string { return "number: two" }).
		With(true, func() string { return "boolean: true" }).
		With("hello", func() string { return "string: hello" }).
		With(nil, func() string { return "null" }).
		With(big.NewInt(20), func() string { return "bigint: 20n" }).
		Otherwise(func() string { return "something else" })

	expected := "something else"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestWithString(t *testing.T) {
	input := "hello world"
	output := NewMatcher[string](input).
		WithString(func() string { return "number: two" }).
		Otherwise(func() string { return "something else" })

	expected := "number: two"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestNotPattern(t *testing.T) {
	input := 2
	output := NewMatcher[string](input).
		With(NotPattern{2}, func() string { return "not number: two" }).
		Otherwise(func() string { return "something else" })

	expected := "something else"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}
