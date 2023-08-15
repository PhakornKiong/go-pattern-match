package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern/pattern"
)

func match(input []int) string {
	return pattern.NewMatcher[string](input).
		WithValues(
			[]any{1, 2, 3, 4},
			func() string { return "Nope" },
		).
		WithValues(
			[]any{
				pattern.Any(),
				pattern.Not(36),
				pattern.Union[int](99, 98),
				255,
			},
			func() string { return "Its a match" },
		).
		Otherwise(func() string { return "Otherwise" })
}

func main() {
	fmt.Println(match([]int{1, 2, 3, 4}))      // "Nope"
	fmt.Println(match([]int{25, 35, 99, 255})) // "Its a match"
	fmt.Println(match([]int{1, 5, 6, 7}))      // "Otherwise"
}
