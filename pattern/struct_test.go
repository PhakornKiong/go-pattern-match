package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {

	t.Run("struct invalid input type", func(t *testing.T) {
		assert := assert.New(t)

		s := Struct()

		output := s.Match("not struct")

		assert.False(output)
	})

	t.Run("fieldValue positive case", func(t *testing.T) {
		type embed struct {
			H string
		}

		type custom struct {
			X int
			Y string
			Z []int
			E embed
		}

		input := custom{0, "hey world", []int{1, 2, 3}, embed{"embedded struct"}}
		assert := assert.New(t)

		s := Struct().
			FieldValue("X", 0).
			FieldValue("Z", []int{1, 2, 3}).
			FieldValue("E", embed{"embedded struct"})

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("fieldValue negative case", func(t *testing.T) {
		type custom struct {
			X int
			Y string
			Z []int
		}

		input := custom{25, "hey world", []int{1, 2, 3}}
		assert := assert.New(t)

		s := Struct().FieldValue("X", 25).FieldValue("Z", []int{1, 5, 3})
		output := s.Match(input)

		assert.False(output)
	})

	t.Run("fieldValue cannot check unexported field", func(t *testing.T) {
		type custom struct {
			x int
		}

		input := custom{25}
		assert := assert.New(t)

		s := Struct().FieldValue("x", 25)
		output := s.Match(input)

		assert.False(output)
	})

	t.Run("fieldValue field does not exists", func(t *testing.T) {
		type custom struct {
			X int
			Y string
			Z []int
		}

		input := custom{25, "hey world", []int{1, 2, 3}}
		assert := assert.New(t)

		s := Struct().FieldValue("some", 25)
		output := s.Match(input)

		assert.False(output)
	})

	t.Run("fieldPattern positive case", func(t *testing.T) {
		type custom struct {
			X int
			Y string
			Z []int
		}

		input := custom{0, "hey world", []int{1, 2, 3}}
		assert := assert.New(t)

		s := Struct().
			FieldPattern("X", Int().Gt(0)).
			FieldPattern("Y", String().Contains("world")).
			FieldPattern("Z", Slice[int]().Head(1).TailPattern(Int().Gte(3)))

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("fieldPattern negative case", func(t *testing.T) {
		type custom struct {
			X int
			Y string
			Z []int
		}

		input := custom{0, "hey world", []int{1, 2, 3}}
		assert := assert.New(t)

		s := Struct().
			FieldPattern("X", Int().Gt(1))

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("fieldPattern field does not exists", func(t *testing.T) {
		type custom struct {
			X int
			Y string
			Z []int
		}

		input := custom{0, "hey world", []int{1, 2, 3}}
		assert := assert.New(t)

		s := Struct().
			FieldPattern("some", Int().Gt(1))

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("fieldPattern cannot check unexported field", func(t *testing.T) {
		type custom struct {
			x int
		}

		input := custom{25}
		assert := assert.New(t)

		s := Struct().FieldPattern("x", Int())
		output := s.Match(input)

		assert.False(output)
	})
}
