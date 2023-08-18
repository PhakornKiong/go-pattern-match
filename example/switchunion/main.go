package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern-match/pattern"
)

func FoodSorter(input string) (output string) {
	switch input {
	case "apple", "strawberry", "orange":
		output = "fruit"
	case "carrot", "pok-choy", "cabbage":
		output = "vegetable"
	default:
		output = "unknown"
	}

	return output
}

func FoodSorterWithPattern(input string) (output string) {

	output = pattern.NewMatcher[string](input).
		WithPattern(
			pattern.Union("apple", "strawberry", "orange"),
			func() string { return "fruit" },
		).
		WithPattern(
			pattern.Union("carrot", "pok-choy", "cabbage"),
			func() string { return "vegetable" },
		).
		Otherwise(func() string { return "unknown" })

	return output
}

func main() {
	fmt.Println(FoodSorterWithPattern("apple"))  // "fruit"
	fmt.Println(FoodSorterWithPattern("orange")) // "fruit"
	fmt.Println(FoodSorterWithPattern("carrot")) // "vegetable"
	fmt.Println(FoodSorterWithPattern("candy"))  // "unknown"
}
