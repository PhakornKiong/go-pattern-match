package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Run("contains int positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().Contains(1).Contains(3)

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("contains int negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().Contains(1).Contains(5)

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("contains invalid input type", func(t *testing.T) {
		assert := assert.New(t)

		s := Slice[int]().Contains(1).Contains(5)

		output := s.Match("invalid input")

		assert.False(output)
	})

	t.Run("contains struct positive case", func(t *testing.T) {
		assert := assert.New(t)

		type custom struct {
			X int
			Y string
		}

		input := []custom{{1, "one"}, {2, "two"}, {3, "three"}}

		s := Slice[custom]().Contains(custom{2, "two"}).Contains(custom{3, "three"})

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("contains struct negative case", func(t *testing.T) {
		assert := assert.New(t)

		type custom struct {
			X int
			Y string
		}

		input := []custom{{1, "one"}, {2, "two"}, {3, "three"}}

		s := Slice[custom]().Contains(custom{2, "???"}).Contains(custom{3, "three"})

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("containsPattern int positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{100, 101, 102}

		s := Slice[int]().ContainsPattern(Int().Gt(101))

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("containsPattern int positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{100, 101, 102}

		s := Slice[int]().ContainsPattern(Not(Int().Gt(102)))

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("containsPattern int negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{100, 101, 102}

		s := Slice[int]().ContainsPattern(Int().Gt(102))

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("containsPattern invalid input type", func(t *testing.T) {
		assert := assert.New(t)

		s := Slice[int]().ContainsPattern(Int().Gt(102))

		output := s.Match("invalid input")

		assert.False(output)
	})

	t.Run("headElement int positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().Head(1)

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("headElement int negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().Head(2)

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("headPattern int positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().HeadPattern(Int().Lt(2))

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("headPattern int negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().HeadPattern(Int().Lt(1))

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("tailElement int positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().Tail(3)

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("tailElement int negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().Tail(2)

		output := s.Match(input)

		assert.False(output)
	})

	t.Run("TailPattern int positive case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().TailPattern(Int().Gt(2))

		output := s.Match(input)

		assert.True(output)
	})

	t.Run("TailPattern int negative case", func(t *testing.T) {
		assert := assert.New(t)

		input := []int{1, 2, 3}

		s := Slice[int]().TailPattern(Int().Gt(3))

		output := s.Match(input)

		assert.False(output)
	})

}
