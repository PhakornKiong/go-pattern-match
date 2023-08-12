package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	t.Run("Int with int input", func(t *testing.T) {
		assert := assert.New(t)
		input := 35
		w := Int()

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int with string input", func(t *testing.T) {
		assert := assert.New(t)
		input := "35"
		w := Int()

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("Int Lt positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := 35
		w := Int().Lt(45)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int Lt negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := 35
		w := Int().Lt(35)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("Int Lte positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := 45
		w := Int().Lte(45)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int Lte negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := 26
		w := Int().Lte(25)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("Int Gt positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := 45
		w := Int().Gt(35)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int Gt negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := 25
		w := Int().Gt(35)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("Int Gte positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := 45
		w := Int().Gte(45)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int Gte negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := 25
		w := Int().Gte(26)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("Int Between positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := 30
		w := Int().Between(25, 35)

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int Between negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := 36
		w := Int().Between(25, 35)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("Int Positive positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := 5
		w := Int().Positive()

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int Positive negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := -5
		w := Int().Positive()

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("Int Negative positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := -5
		w := Int().Negative()

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("Int Negative negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := 5
		w := Int().Negative()

		output := w.Match(input)

		assert.False(output)
	})

}
