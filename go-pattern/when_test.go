package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPattern(t *testing.T) {
	t.Run("whenPattern input positive case", func(t *testing.T) {
		assert := assert.New(t)
		input := "test"
		w := When[string](func(s string) bool { return s == "test" })

		output := w.Match(input)

		assert.True(output)
	})

	t.Run("whenPattern input negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "test"
		w := When[string](func(s string) bool { return s == "wrong" })

		output := w.Match(input)

		assert.False(output)
	})
}
