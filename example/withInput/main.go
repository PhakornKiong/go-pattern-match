package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern/pattern"
)

func match(input string) string {
	unionPattern := pattern.Union("mango", "papaya")
	return pattern.NewMatcher[string](input).
		WithValue(
			"durian",
			func() string { return "I like durian" },
		).
		WithPattern(
			// matches if input is any of "mango" or "papaya"
			unionPattern,
			func() string { return fmt.Sprintf("I like %s", input) },
		).
		Otherwise(func() string { return "I dont like fruits" })
}

func main() {
	fmt.Println(match("durian")) // "I like durian"
	fmt.Println(match("apple"))  // "I dont like fruits"
	fmt.Println(match("mango"))  // "I like mango"
	fmt.Println(match("papaya")) // "I like papaya"
}
