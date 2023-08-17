package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("int-int map positive case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().KeyVal(1, 55).KeyVal(2, 44)

		output := w.Match(inputMap)

		assert.True(output)
	})

	t.Run("int-string map negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]string{
			1: "john",
			2: "wick",
			3: "doe",
		}

		w := Map[int, string]().KeyVal(1, "john").KeyVal(2, "doe")

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("int-string map with missing key negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]string{
			1: "john",
			2: "wick",
			3: "doe",
		}

		w := Map[int, string]().KeyVal(5, "john").KeyVal(2, "doe")

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("not map input negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := "some string pretending to be map"

		w := Map[int, string]().KeyVal(5, "john").KeyVal(2, "doe")

		output := w.Match(input)

		assert.False(output)
	})

	t.Run("struct positive case", func(t *testing.T) {
		assert := assert.New(t)

		type custom struct {
			X string
			Y string
		}

		inputMap := map[int]custom{
			1: {"1", "2"},
			2: {"3", "4"},
			3: {"5", "6"},
		}

		w := Map[int, custom]().KeyVal(1, custom{"1", "2"}).KeyVal(3, custom{"5", "6"})

		output := w.Match(inputMap)

		assert.True(output)
	})

	t.Run("struct negative case", func(t *testing.T) {
		assert := assert.New(t)

		type custom struct {
			X string
			Y string
		}

		inputMap := map[int]custom{
			1: {"1", "2"},
			2: {"3", "4"},
			3: {"5", "6"},
		}

		w := Map[int, custom]().KeyVal(1, custom{"1", "2"}).KeyVal(3, custom{"5", "4"})

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("map[string]string positive case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]map[string]string{
			1: {"one": "one value"},
			2: {"one": "1", "two": "2"},
			3: {"three": "3"},
		}

		w := Map[int, map[string]string]().KeyVal(2, map[string]string{"one": "1", "two": "2"}).KeyVal(3, map[string]string{"three": "3"})

		output := w.Match(inputMap)

		assert.True(output)
	})

	t.Run("map[string]string negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]map[string]string{
			1: {"one": "one value"},
			2: {"one": "1", "two": "2"},
			3: {"three": "3"},
		}

		w := Map[int, map[string]string]().KeyVal(2, map[string]string{"hehe": "1", "two": "2"}).KeyVal(3, map[string]string{"three": "3"})

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("Key int-int map positive case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().Key(1).Key(2).Key(3)

		output := w.Match(inputMap)

		assert.True(output)
	})

	t.Run("Key int-int map negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().Key(4).Key(2).Key(3)

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("Val int-int map positive case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().Val(55).Val(44).Val(66)

		output := w.Match(inputMap)

		assert.True(output)
	})

	t.Run("Val int-int map negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().Val(2).Val(44).Val(66)

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("Val map[string]string positive case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]map[string]string{
			1: {"one": "one value"},
			2: {"one": "1", "two": "2"},
			3: {"three": "3"},
		}

		w := Map[int, map[string]string]().Val(map[string]string{"one": "1", "two": "2"}).Val(map[string]string{"three": "3"})

		output := w.Match(inputMap)

		assert.True(output)
	})

	t.Run("Val map[string]string negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]map[string]string{
			1: {"one": "one value"},
			2: {"one": "1", "two": "2"},
			3: {"three": "3"},
		}

		w := Map[int, map[string]string]().Val(map[string]string{"one": "1", "two": "2"}).Val(map[string]string{"three": "44"})

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("KeyValPattern int-int map positive case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().KeyValPatterns(1, Int().Gt(54)).KeyValPatterns(2, Int().Lt(45))

		output := w.Match(inputMap)

		assert.True(output)
	})

	t.Run("KeyValPattern int-int map negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().KeyValPatterns(1, Int().Gt(54)).KeyValPatterns(2, Int().Lt(44))

		output := w.Match(inputMap)

		assert.False(output)
	})

	t.Run("KeyValPattern key does not exists negative case", func(t *testing.T) {
		assert := assert.New(t)

		inputMap := map[int]int{
			1: 55,
			2: 44,
			3: 66,
		}

		w := Map[int, int]().KeyValPatterns(4, Int().Gt(54)).KeyValPatterns(2, Int().Lt(44))

		output := w.Match(inputMap)

		assert.False(output)
	})

}
