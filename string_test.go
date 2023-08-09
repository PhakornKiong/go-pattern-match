package pattern

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("String with string input", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := Union(input, "hello", "there")

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String with number", func(t *testing.T) {
		assert := assert.New(t)
		input := 25
		w := String()

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("String MaxLength positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().MaxLength(4)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String MaxLength negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().MaxLength(3)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("String MinLength positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().MinLength(4)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String MinLength negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().MinLength(5)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("String StartsWith positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().StartsWith("te")

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String StartsWith negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().StartsWith("no")

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("String EndsWith positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().EndsWith("st")

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String EndsWith negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().EndsWith("no")

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("String Includes positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().Includes("es")

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String Includes negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := String().Includes("no")

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("String Regex positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		regex := regexp.MustCompile("^te.*")
		w := String().Regex(regex)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String Regex negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		regex := regexp.MustCompile("^no.*")
		w := String().Regex(regex)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("String Combine positive case 01", func(t *testing.T) {
		assert := assert.New(t)
		input := "test combine positive case"
		regex := regexp.MustCompile("^te.*")
		w := String().
			Regex(regex).
			MinLength(4).
			MaxLength(26).
			Includes("est").
			EndsWith("e case").
			StartsWith("test")

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("String Combine negative case 01", func(t *testing.T) {
		assert := assert.New(t)
		input := "test combine negative case"
		regex := regexp.MustCompile("^te.*")
		w := String().
			Regex(regex).
			MinLength(4).
			MaxLength(25). // Fail
			Includes("est").
			EndsWith("e case").
			StartsWith("test")

		output := w.Match(input)

		assert.False(output)
	})

}
