package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyPattern(t *testing.T) {
	t.Run("anyPattern always return true", func(t *testing.T) {
		assert := assert.New(t)
		a := Any()

		output := a.Match("")
		assert.False(output)
	})

}
