package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot(t *testing.T) {
	t.Run("not input positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := Not("wrong")

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("not input negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := Not(input)

		output := w.Match(input)

		assert.False(output)
	})
}

func TestNotPattern(t *testing.T) {
	t.Run("notPattern input positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		stringP := String().MinLength(2)
		w := NotPattern(stringP)

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("notPattern input negative case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		stringP := String().MinLength(5)
		w := NotPattern(stringP)

		output := w.Match(input)

		assert.True(output)
	})
}
