package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	t.Run("Union input positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := Union(input, "hello", "there")

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Union input negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := Union("hello", "no", "match")

		output := w.Match(input)

		assert.False(output)
	})
}

func TestUnionPattern(t *testing.T) {
	t.Run("UnionPattern with stringPattern positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		stringP := String().Contains("test")
		stringN := String().Contains("hey")

		w := UnionPattern(stringP, stringN)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("UnionPattern with stringPattern negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		stringN := String().Contains("hey")

		w := UnionPattern(stringN, stringN)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("UnionPattern with union positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		unionP := Union(input, "no", "no")
		unionN := Union("no", "no", "no")

		w := UnionPattern(unionP, unionN)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("UnionPattern with union negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		unionN := Union("no", "no", "no")

		w := UnionPattern(unionN, unionN)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("UnionPattern with intersection positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		intersectionP := Intersection(input, input, input)
		intersectionN := Intersection("no", input, input)

		w := UnionPattern(intersectionP, intersectionN)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("UnionPattern with intersection negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		intersectionN := Intersection("no", input, input)

		w := UnionPattern(intersectionN, intersectionN)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("UnionPattern with intersectionPattern positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		intersectionP := Intersection(input, input, input)
		intersectionN := Intersection("no", "test", "test")

		intersectionPatternP := IntersectionPattern(intersectionP, intersectionP)
		intersectionPatternN := IntersectionPattern(intersectionP, intersectionN)

		w := UnionPattern(intersectionPatternP, intersectionPatternN)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("UnionPattern with intersectionPattern negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		intersectionN := Intersection("no", input, input)
		intersectionPatternN := IntersectionPattern(intersectionN, intersectionN)

		w := UnionPattern(intersectionPatternN, intersectionPatternN)

		output := w.Match(input)

		assert.False(output)
	})
}
