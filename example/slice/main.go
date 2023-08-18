package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern/pattern"
)

func match(input []int) string {
	pattern1 := pattern.Slice[int]().
		Head(1).
		Tail(100)

	subPattern2 := pattern.Int().Between(75, 100)

	pattern2 := pattern.Slice[int]().
		Contains(25).
		Contains(50).
		ContainsPattern(subPattern2)

	subHeadPattern3 := pattern.Int().Gt(1000)
	subTailattern3 := pattern.Int().Gt(2500)
	pattern3 := pattern.Slice[int]().
		HeadPattern(subHeadPattern3).
		TailPattern(subTailattern3)

	return pattern.NewMatcher[string](input).
		WithPattern(
			pattern1,
			func() string { return "pattern 1" },
		).
		WithPattern(
			pattern2,
			func() string { return "pattern 2" },
		).
		WithPattern(
			pattern3,
			func() string { return "pattern 3" },
		).
		Otherwise(func() string { return "No pattern matched" })
}

func main() {
	fmt.Println(match([]int{1, 2, 3, 100}))       // "pattern 1"
	fmt.Println(match([]int{2, 25, 85, 50}))      // "pattern 2"
	fmt.Println(match([]int{1001, 25, 3, 25001})) // "pattern 3"
}
