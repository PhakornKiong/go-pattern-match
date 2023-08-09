package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	t.Run("Intersection input positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := Intersection(input, input, input)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Intersection input negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := Intersection("hello", "no", "match")

		output := w.Match(input)

		assert.False(output)
	})
}

func TestIntersectionPattern(t *testing.T) {

	t.Run("IntersectionPattern with stringPattern positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		stringP := String().Includes(input)

		w := IntersectionPattern(stringP, stringP)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("IntersectionPattern with stringPattern negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		stringP := String().Includes(input)
		stringN := String().Includes("hey")

		w := IntersectionPattern(stringP, stringN)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("IntersectionPattern with intersection positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		intersectionP := Intersection(input, input, input)

		w := IntersectionPattern(intersectionP, intersectionP)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("IntersectionPattern with intersection negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		intersectionP := Intersection(input, input, input)
		intersectionN := Intersection("no", input, input)

		w := IntersectionPattern(intersectionP, intersectionN)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("IntersectionPattern with union positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		unionP := Union(input, "no", "no")

		w := IntersectionPattern(unionP, unionP)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("IntersectionPattern with union negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		unionP := Union(input, "no", "no")
		unionN := Union("no", "no", "no")

		w := IntersectionPattern(unionP, unionN)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("IntersectionPattern with UnionPattern positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		unionP := Union(input, "no", "no")
		unionN := Union("no", "no", "no")

		unionPatternP := UnionPattern(unionP, unionN)

		w := IntersectionPattern(unionPatternP, unionPatternP)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("IntersectionPattern with UnionPattern negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		unionP := Union(input, "no", "no")
		unionN := Union("no", "no", "no")

		unionPatternP := UnionPattern(unionP, unionN)
		unionPatternN := UnionPattern(unionN, unionN)

		w := IntersectionPattern(unionPatternP, unionPatternN)

		output := w.Match(input)

		assert.False(output)
	})

}
