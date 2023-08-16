package main

import (
	"fmt"
	"regexp"

	"github.com/phakornkiong/go-pattern/pattern"
)

func match(input string) string {
	pattern1 := pattern.String().
		StartsWith("hello").
		EndsWith("world").
		MaxLength(11)

	pattern2 := pattern.String().
		Contains("dni").
		Regex(regexp.MustCompile("night$"))

	pattern3 := pattern.String().
		MinLength(3)

	pattern4 := pattern.String()

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
		WithPattern(
			pattern4,
			func() string { return "pattern 4" },
		).
		Otherwise(func() string { return "This is impossible" })
}

func main() {
	fmt.Println(match("hello world")) // "pattern 1"
	fmt.Println(match("goodnight"))   // "pattern 2"
	fmt.Println(match("abc"))         // "pattern 3"
	fmt.Println(match("ab"))          // "pattern 4"
	fmt.Println(match(""))            // "pattern 4"

}
