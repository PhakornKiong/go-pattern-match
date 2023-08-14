package pattern

import (
	"fmt"
	"reflect"
	"regexp"
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
		WithValue(MyStruct{25, 35}, func() OtherStruct { return expected }).
		WithValue(MyStruct{25, 35}, func() OtherStruct { return OtherStruct{} }).
		Otherwise(func() OtherStruct { return OtherStruct{} })

	if output != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, output)
	}
}

func TestMatcher(t *testing.T) {
	input := 2
	expected := "number: two"

	output := NewMatcher[string](input).
		WithValue(2, func() string { return expected }).
		Otherwise(func() string { return "something else" })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestMatcher2(t *testing.T) {
	input := 100
	expected := "something else"

	output := NewMatcher[string](input).
		WithValue(2, func() string { return "number: two" }).
		Otherwise(func() string { return expected })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestMatcherWithPatterns(t *testing.T) {
	unexpected := "did not match"
	expected := "matched"

	t.Run("string patterns positive case", func(t *testing.T) {
		assert := assert.New(t)

		firstP := String().Contains("you")
		secondP := String().MinLength(10)
		thirdP := String().EndsWith("us")

		input := []string{"you", "this is more than 10 character", "world ends with us"}
		output := NewMatcher[string](input).
			WithPatterns(
				Patteners(firstP, secondP, thirdP),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("string patterns negative case", func(t *testing.T) {
		assert := assert.New(t)

		firstP := String().Contains("you")
		secondP := String().MinLength(256)
		thirdP := String().EndsWith("us")

		input := []string{"you", "this is more than 10 character", "world ends with us"}
		output := NewMatcher[string](input).
			WithPatterns(
				Patteners(firstP, secondP, thirdP),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("mismatch input length negative case", func(t *testing.T) {
		assert := assert.New(t)

		firstP := String().Contains("you")
		secondP := String().MinLength(10)
		thirdP := String().EndsWith("us")

		input := []string{"you", "this is more than 10 character"}
		output := NewMatcher[string](input).
			WithPatterns(
				Patteners(firstP, secondP, thirdP),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("mismatch pattern length negative case", func(t *testing.T) {
		assert := assert.New(t)

		firstP := String().Contains("you")
		secondP := String().MinLength(10)

		input := []string{"you", "this is more than 10 character", "world ends with us"}
		output := NewMatcher[string](input).
			WithPatterns(
				Patteners(firstP, secondP),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})
}

func TestMatcherWithValues(t *testing.T) {
	unexpected := "did not match"
	expected := "matched"

	t.Run("int values positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{25, 35, 99}
		output := NewMatcher[string](input).
			WithValues(
				input,
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("int values negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{25, 35, 99}
		output := NewMatcher[string](input).
			WithValues(
				[]int{25, 45, 99},
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("mismatch input length negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{25, 99}
		output := NewMatcher[string](input).
			WithValues(
				[]int{25, 45, 99},
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("mismatch values length negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{25, 35, 99}
		output := NewMatcher[string](input).
			WithValues(
				[]int{25, 99},
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})
}

func TestMatcherWithNotPattern(t *testing.T) {
	unexpected := "did not match"
	expected := "matched"

	t.Run("int input positive case ", func(t *testing.T) {
		assert := assert.New(t)

		input := 2
		output := NewMatcher[string, int](input).
			WithPattern(Not(3), func() string { return expected }).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})
	t.Run("int input negative case ", func(t *testing.T) {
		assert := assert.New(t)

		input := 2
		output := NewMatcher[string, int](input).
			WithPattern(Not(2), func() string { return unexpected }).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("string input positive case ", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello"
		output := NewMatcher[string, string](input).
			WithPattern(Not("world"), func() string { return expected }).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})
	t.Run("string input negative case ", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello"
		output := NewMatcher[string, string](input).
			WithPattern(Not("hello"), func() string { return unexpected }).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})
	t.Run("struct input positive case ", func(t *testing.T) {
		assert := assert.New(t)

		type MyStruct struct {
			x int
		}

		input := MyStruct{5}
		output := NewMatcher[string, MyStruct](input).
			WithPattern(Not(MyStruct{6}), func() string { return expected }).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})
	t.Run("struct input negative case ", func(t *testing.T) {
		assert := assert.New(t)

		type MyStruct struct {
			x int
		}

		input := MyStruct{5}
		output := NewMatcher[string, MyStruct](input).
			WithPattern(Not(MyStruct{5}), func() string { return unexpected }).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})
}

func TestMatcherWithWhenPattern(t *testing.T) {
	input := 5
	expected := "greater than three"
	output := NewMatcher[string, int](input).
		WithPattern(When[int](func(i int) bool { return i > 3 }),
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
		WithPattern(When[string](func(i string) bool { return input == i }),
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
		WithPattern(When[predicate](func(i predicate) bool { return i.x == 33 }),
			func() string {
				return expected
			}).
		Otherwise(func() string { return "did not match" })

	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestMatcherWithUnion(t *testing.T) {
	unexpected := "did not match"
	expected := "matched"
	t.Run("int intput positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := 356
		output := NewMatcher[string, int](input).
			WithPattern(
				Union[int](25, 356, 123),
				func() string { return expected },
			).
			Otherwise(func() string { return "did not match" })

		assert.Equal(expected, output)
	})

	t.Run("int intput negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := 356
		output := NewMatcher[string, int](input).
			WithPattern(
				Union[int](2, 1, 3),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("string input positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "test union"
		output := NewMatcher[string, string](input).
			WithPattern(Union[string]("five", "six", input), func() string { return expected }).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("string input negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "test union"
		expected := "matched"
		output := NewMatcher[string, string](input).
			WithPattern(Union[string]("five", "six", "nine"), func() string { return unexpected }).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("stringPattern input positive case", func(t *testing.T) {
		assert := assert.New(t)
		fmt.Printf("%+v", reflect.TypeOf(UnionPattern(String().EndsWith("union"))))

		input := "test union"
		expected := "matched"
		output := NewMatcher[string, string](input).
			WithPattern(UnionPattern(String().EndsWith("union")),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("intersectionPattern input positive case", func(t *testing.T) {
		assert := assert.New(t)
		fmt.Printf("%+v", reflect.TypeOf(UnionPattern(String().EndsWith("union"))))

		input := "test union"
		expected := "matched"

		i := IntersectionPattern(
			String().EndsWith("union"),
			String().MinLength(5),
			String().MaxLength(10),
		)

		output := NewMatcher[string, string](input).
			WithPattern(i,
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("intersectionPattern input negative case", func(t *testing.T) {
		assert := assert.New(t)
		fmt.Printf("%+v", reflect.TypeOf(UnionPattern(String().EndsWith("union"))))

		input := "test union"
		expected := "matched"

		i := IntersectionPattern(
			String().EndsWith("union"),
			String().MinLength(5),
			String().MaxLength(9),
		)

		output := NewMatcher[string, string](input).
			WithPattern(i,
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("unionPattern input positive case", func(t *testing.T) {
		assert := assert.New(t)
		fmt.Printf("%+v", reflect.TypeOf(UnionPattern(String().EndsWith("union"))))

		input := "test union"
		expected := "matched"

		u1 := UnionPattern(String().EndsWith("union"), String().StartsWith("test"))
		u2 := UnionPattern(String().EndsWith("union"), String().StartsWith("test"))

		i := IntersectionPattern(u1, u2)

		output := NewMatcher[string, string](input).
			WithPattern(i,
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

}

func TestMatcherWithIntersection(t *testing.T) {
	unexpected := "did not match"
	t.Run("int intput positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := 356
		expected := "matched"
		output := NewMatcher[string, int](input).
			WithPattern(
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
			WithPattern(
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
			WithPattern(
				Intersection[string](input, input, input),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

}

func TestStringPattern(t *testing.T) {
	unexpected := "did not match"
	expected := "matched"

	t.Run("int input negative case ", func(t *testing.T) {
		assert := assert.New(t)

		input := 356
		output := NewMatcher[string, int](input).
			WithPattern(
				String(),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("string input positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String(),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("StartsWith positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().StartsWith("hello"),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("StartsWith negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().StartsWith("world"),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("EndsWith positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().EndsWith("world"),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("EndsWith negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().EndsWith("hello"),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("MinLength positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().MinLength(5),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("MinLength negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().MinLength(10),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("Regex positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().Regex(regexp.MustCompile("hello")),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("Regex negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().Regex(regexp.MustCompile("universe$")),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})

	t.Run("Contains positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().Contains("world"),
				func() string { return expected },
			).
			Otherwise(func() string { return unexpected })

		assert.Equal(expected, output)
	})

	t.Run("Contains negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "hello world"
		output := NewMatcher[string, string](input).
			WithPattern(
				String().Contains("universe"),
				func() string { return unexpected },
			).
			Otherwise(func() string { return expected })

		assert.Equal(expected, output)
	})
}
