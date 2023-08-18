# Go-Pattern

[![Test](https://github.com/PhakornKiong/go-pattern/actions/workflows/test.yml/badge.svg)](https://github.com/PhakornKiong/go-pattern/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/PhakornKiong/go-pattern/branch/master/graph/badge.svg?token=IL7G963OAF)](https://codecov.io/gh/PhakornKiong/go-pattern)
[![Go Report Card](https://goreportcard.com/badge/github.com/phakornkiong/go-pattern)](https://goreportcard.com/report/github.com/phakornkiong/go-pattern)
[![GoDoc](https://godoc.org/phakornkiong/go-pattern?status.svg)](https://godoc.org/github.com/phakornkiong/go-pattern)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/PhakornKiong/go-pattern/blob/master/LICENSE)

Pattern Matching library for Go

```go
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
```

## Why Pattern Matching?

Pattern matching originated in functional programming languages like ML, Haskell, and Erlang. It provides a powerful technique for control flow based on matching inputs against patterns.

Compared to imperative control flow using conditionals and switch statements, pattern matching enables more concise and declarative code, especially when handling complex conditional logic. It avoids verbose boilerplate code and better conveys intent.

## Key Components

- `Patterner`: This is an interface that requires the implementation of a `Match` function. Any type that implements this interface can be used as a pattern in the matcher.

- `Handler`: This is a function type that returns a generic type `T`. This function is called when a match is found.

- `Matcher`: This is a struct that holds the value to be matched, a flag indicating if a match has been found, and the response to be returned when a match is found.

## Documentation

### `NewMatcher[T any, V any](input V) *Matcher[T, V]`

This function creates a new Matcher instance. The generics `T` and `V` represent any types.

`T` is the type that the handler function will return when a match is found.

`V` is the type of the input value that will be matched against the patterns.

### `.WithPattern(pattern Patterner, fn Handler[T]) *Matcher[T, V]`

It checks if the provided pattern matches the entire input. If a match is found, it calls the provided Handler function and return the response `T`.

### `.WithPatterns(patterns []Patterner, fn Handler[T]) *Matcher[T, V]`

It checks each of the provided patterns against <b>each of the input</b>. If a match is found, it calls the provided Handler function and return the response `T`.

### `.WithValue(value V, fn Handler[T]) *Matcher[T, V]`

It checks for deep equality between the provided value and the entire input. If a match is found, it calls the provided Handler function and return the response `T`.

### `.WithValues(values any, fn Handler[T]) *Matcher[T, V]`

If a `Patterner` is provided as the pattern, it will check if any of the provided values matches the input by calling the `Match` method on each value, else it will do a deep equality check on each value.
If a match is found, it calls the provided Handler function and return the response `T`.

This enables more flexible pattern match where the provided values can be a `patterner` or just `actual value`.

For example, now you can use many of the built-in patterns like `Union`, `Any`, `Not`. See the [Patterns](#patterns) section for more details.

```go
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

match([]int{1, 2, 3, 4}) // "Nope"
match([]int{25, 35, 99, 255}) // "Its a match"
match([]int{1, 5, 6, 7}) // "Otherwise"
```

### `.Otherwise(fn Handler[T]) *Matcher[T, V]`

It is called when no match is found for the input. It calls the provided Handler function and return the response `T`.

## [Patterns](#patterns)

Patterns provide a way to declaratively match values. In general, they all implements the `Patterner` interface which requires a `Match(any) bool` method.

Some common patterns included are:

- [Any Pattern](#any-pattern)
- [Not Pattern](#not-pattern)
- [NotPattern Pattern](#notpattern-pattern)
- [When Pattern](#when-pattern)
- [Union Pattern](#union-pattern)
- [UnionPattern Pattern](#unionpattern-pattern)
- [IntersectionPattern Pattern](#intersectionpattern-pattern)
- [String Pattern](#string-pattern)
- [Int Pattern](#int-pattern)
- [Slice Pattern](#slice-pattern)

TODO patterns

- [ ] Maps
- [ ] Struct

Currently you can use [When Pattern](#when-pattern) to do custom matching logic for these pattern.

### [Any Pattern](#any-pattern)

`pattern.Any()` returns a `Patterner` that matches any value.

```go
func match(input int) string {
	return pattern.NewMatcher[string](input).
			WithValue(
				2,
				func() string { return "Nope" },
			).
			WithPattern(
				pattern.Any(), // Always matches
				func() string { return "Its a match" },
			).
			Otherwise(func() string { return "Nope" })
}

match(5) // "Its a match"
match(6) // "Its a match"
match(7) // "Its a match"
```

### [Not Pattern](#not-pattern)

`pattern.Not(input)` returns a `Patterner` that matches any value other than the input by comparing using deep equality.

```go
func match(input int) string {
	return pattern.NewMatcher[string](input).
			WithValue(
				2,
				func() string { return "2" },
			).
			WithPattern(
				pattern.Not(3), // Always matches if not 3
				func() string { return "Its a match" },
			).
			Otherwise(func() string { return "Otherwise" })
}

match(3) // "Otherwise"
match(6) // "Its a match"
match(7) // "Its a match"
```

### [NotPattern Pattern](#notpattern-pattern)

Similar to `Not` pattern, but accepts only a `Patterner` instead of a value. Used mainly to get inverse of a pattern.

```go
func match(input int) string {
	intPattern := pattern.Int().Between(25, 35)
	return pattern.NewMatcher[string](input).
			WithValue(
				2,
				func() string { return "2" },
			).
			WithPattern(
				// Always matches if not in between 25 & 35
				pattern.NotPattern(intPattern),
				func() string { return "Its a match" },
			).
			Otherwise(func() string { return "Otherwise" })
}

match(2) // "2"
match(30) // "Otherwise"
match(36) // "Its a match"
match(24) // "Its a match"
```

### [When Pattern](#when-pattern)

`When` pattern accepts a predicate, which is a function that takes a value and returns a boolean. This pattern matches when the predicate function returns true for the input value.

```go
func match(input int) string {
	return pattern.NewMatcher[string](input).
			WithValue(
				2,
				func() string { return "2" },
			).
			WithPattern(
				// match if input larger than 100
				pattern.When[int](func(i int) bool { return i > 100 }),
				func() string { return "Its a match" },
			).
			Otherwise(func() string { return "Otherwise" })
}

match(2) // "2"
match(99) // "Otherwise"
match(100) // "Its a match"
match(105) // "Its a match"
```

### [Union Pattern](#union-pattern)

`Union` pattern matches if the input equals any of the provided values by using deep equality check.

```go
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

FoodSorterWithPattern("apple")  // "fruit"
FoodSorterWithPattern("orange") // "fruit"
FoodSorterWithPattern("carrot") // "vegetable"
FoodSorterWithPattern("candy")  // "unknown"
```

### [UnionPattern Pattern](#unionpattern-pattern)

Similar to [Union pattern](#union-pattern) but accepts only `Patterner` instead of values. Useful for extending patterning capabilities.

### [IntersectionPattern Pattern](#intersectionpattern-pattern)

`IntersectionPattern` accepts multiple patterns and matches if the input matches all of them. Useful for extending patterning capabilities.

### [String Pattern](#string-pattern)

`String` pattern matches string values. It provides additional methods to match on string contents:

#### `StartsWith(value string) stringPattern`

Chainable method for matching strings starting with the provided value.

#### `EndsWith(value string) stringPattern`

Chainable method for matching strings ending with the provided value.

#### `Contains(value string) stringPattern`

Chainable method for matching strings containing the provided value.

#### `Regex(value string) stringPattern`

Chainable method for matching strings according to the provided regular expression.

#### `MinLength(value int) stringPattern`

Chainable method for matching strings with a minimum length of the provided value.

#### `MaxLength(value int) stringPattern`

Chainable method for matching strings with a maximum length of the provided value.

Here is an example of how to use these methods:

```go
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

match("hello world") // "pattern 1"
match("goodnight")   // "pattern 2"
match("abc")         // "pattern 3"
match("ab")          // "pattern 4"
match("")            // "pattern 4"

```

### [Slice Pattern](#slice-pattern)

`Slice` pattern matches slice values. It provides additional methods to match on slice contents:

#### `Head(v V) slicePattern[V]`

Chainable method for the first element of the input slice to equal the provided value.

### `HeadPattern(p Patterner) slicePattern[V]`

Chained method for the first element of the input slice to match the provided pattern. It calls the underlying `Patterner`'s `Match` method.

### `Tail(v V) slicePattern[V]`

Chainable method for the last element of the input slice to equal the provided value.

### `TailPattern(p Patterner) slicePattern[V]`

Chained method for the last element of the input slice to match the provided pattern. It calls the underlying `Patterner`'s `Match` method.

### `Contains(v V) slicePattern[V]`

Chainable method for the input slice to contain the provided value. Can be used multiple times to check for multiple values.

### `Contains(p Patterner) slicePattern[V]`

Chainable method for the input slice to contain element that matches the provided pattern. It calls the underlying `Patterner`'s `Match` method. Can be used multiple times to check for multiple patterns.

```go
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

match([]int{1, 2, 3, 100})       // "pattern 1"
match([]int{2, 25, 85, 50})      // "pattern 2"
match([]int{1001, 25, 3, 25001}) // "pattern 3"
```

## Examples

You can find more examples and usage scenarios in the `fxStrategy` and `switchUnion` files in the repository. Here are the direct links:

- [fxStrategy](https://github.com/PhakornKiong/go-pattern/blob/master/example/fxstrategy/main.go)
- [switchUnion](https://github.com/PhakornKiong/go-pattern/blob/master/example/switchunion/main.go)

These files a demonstrate its common use case. You may also refer to the test file for more information

## Inspiration

This library is heavily inspired by [ts-pattern](https://github.com/gvergnaud/ts-pattern) which provides powerful pattern matching capabilities for TypeScript. The goal is to provide a similar experience in Go.
