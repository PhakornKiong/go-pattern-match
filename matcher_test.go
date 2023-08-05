package pattern

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
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
	output := NewMatcher[string, int](input).
		With(Not[int](2), func() string { return "not number: two" }).
		Otherwise(func() string { return "something else" })

	expected := "something else"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestWhenPattern(t *testing.T) {
	input := 5
	expected := "greater than three"
	output := NewMatcher[string, int](input).
		With(When[int](func(i int) bool { return i > 3 }),
			func() string {
				return expected
			}).
		Otherwise(func() string { return "not greater than three" })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestWhenPatternString(t *testing.T) {
	input := "hey there"
	expected := "string matched"
	output := NewMatcher[string, string](input).
		With(When[string](func(i string) bool { return input == i }),
			func() string {
				return expected
			}).
		Otherwise(func() string { return "not greater than three" })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestWhenPatternWithStruct(t *testing.T) {
	type predicate struct {
		x int
	}

	input := predicate{33}
	expected := "string matched"
	output := NewMatcher[string, predicate](input).
		With(When[predicate](func(i predicate) bool { return i.x == 33 }),
			func() string {
				return expected
			}).
		Otherwise(func() string { return "did not match" })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestUnionPatternString(t *testing.T) {

	input := "test union"
	expected := "matched"
	output := NewMatcher[string, string](input).
		With(Union[string]("five", "six", "test union"), func() string { return expected }).
		Otherwise(func() string { return "did not match" })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestUnionPatternInt(t *testing.T) {

	input := 356
	expected := "matched"
	output := NewMatcher[string, int](input).
		With(
			Union[int](25, 356, 123),
			func() string { return expected },
		).
		Otherwise(func() string { return "did not match" })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestIntersection(t *testing.T) {
	unexpected := "did not match"
	t.Run("int intput positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := 356
		expected := "matched"
		output := NewMatcher[string, int](input).
			With(
				Intersection[int](356, 356, 356),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("int intput negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := 356
		expected := "matched"
		output := NewMatcher[string, int](input).
			With(
				Intersection[int](356, 1, 356),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("string input positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		expected := "matched"
		output := NewMatcher[string, string](input).
			With(
				Intersection[string](input, input, input),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

}
