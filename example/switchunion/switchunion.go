package switchunion

import (
	"github.com/phakornkiong/go-pattern"
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
		With(
			pattern.Union("apple", "strawberry", "orange"),
			func() string { return "fruit" },
		).
		With(
			pattern.Union("carrot", "pok-choy", "cabbage"),
			func() string { return "vegetable" },
		).
		Otherwise(func() string { return "unknown" })

	return output
}

func RunComparisonFoodMatcher() {
	FoodSorterWithPattern("apple")  // "fruit"
	FoodSorterWithPattern("carrot") // "vegetable"
	FoodSorterWithPattern("candy")  // "unknown"
}
