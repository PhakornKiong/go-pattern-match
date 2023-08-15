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
			[]any{pattern.Any(), pattern.Not(36), pattern.Union[int](99, 98), 255},
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

- `Pattener`: This is an interface that requires the implementation of a `Match` function. Any type that implements this interface can be used as a pattern in the matcher.

- `Handler`: This is a function type that returns a generic type `T`. This function is called when a match is found.

- `Matcher`: This is a struct that holds the value to be matched, a flag indicating if a match has been found, and the response to be returned when a match is found.

## Documentation

### `NewMatcher[T any, V any](input V) *Matcher[T, V]`

This function creates a new Matcher instance. The generics `T` and `V` represent any types.

`T` is the type that the handler function will return when a match is found.

`V` is the type of the input value that will be matched against the patterns.

### `.WithPattern(pattern Pattener, fn Handler[T]) *Matcher[T, V]`

It checks if the provided pattern matches the entire input. If a match is found, it calls the provided Handler function and return the response `T`.

### `.WithPatterns(patterns []Pattener, fn Handler[T]) *Matcher[T, V]`

It checks each of the provided patterns against <b>each of the input</b>. If a match is found, it calls the provided Handler function and return the response `T`.

### `.WithValue(value V, fn Handler[T]) *Matcher[T, V]`

It checks for deep equality between the provided value and the entire input. If a match is found, it calls the provided Handler function and return the response `T`.

### `.WithValues(values any, fn Handler[T]) *Matcher[T, V]`

If a `Patterner` is provided as the pattern, it will check if any of the provided values matches the input by calling the `Match` method on each value, else it will do a deep equality check on each value.
If a match is found, it calls the provided Handler function and return the response `T`.

This enables more flexible pattern match where the provided values can be a `patterner` or just `actual value`.

For example, now you can use many of the built-in patterns like `Union`, `Any`, `Not`...

```go
func match(input []int) string {
	return pattern.NewMatcher[string](input).
			WithValues(
				[]any{1, 2, 3, 4},
				func() string { return "Nope" },
			).
			WithValues(
				[]any{pattern.Any(), pattern.Not(36), pattern.Union[int](99, 98), 255},
				func() string { return "Its a match" },
			).
			Otherwise(func() string { return "Otherwise" })
}

match([]int{1, 2, 3, 4}) // "Nope"
match([]int{25, 35, 99, 255}) // "Its a match"
match([]int{1, 5, 6, 7}) // "Otherwise"
```

## `.Otherwise(fn Handler[T]) *Matcher[T, V]`

It is called when no match is found for the input. It calls the provided Handler function and return the response `T`.

## Patterns

Patterns provide a way to declaratively match values. In general, they all implements the `Patterner` interface which requires a `Match(any) bool` method.

Some common patterns included are:

- `Any`
- `Not` & `NotPattern`
- `When`
- `String`
- `Int`
- `Union` & `UnionPattern`
- `Intersection` & `IntersectionPattern`

### `Any` pattern

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

### `Not` pattern

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

### `NotPattern` pattern

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
				pattern.NotPattern(intPattern), // Always matches if not in between 25 & 35
				func() string { return "Its a match" },
			).
			Otherwise(func() string { return "Otherwise" })
}

match(2) // "2"
match(30) // "Otherwise"
match(36) // "Its a match"
match(24) // "Its a match"
```

###

## Examples

You can find more examples and usage scenarios in the `fxStrategy` and `switchUnion` files in the repository. Here are the direct links:

- [fxStrategy](https://github.com/PhakornKiong/go-pattern/blob/master/example/fxstrategy/main.go)
- [switchUnion](https://github.com/PhakornKiong/go-pattern/blob/master/example/switchunion/main.go)

These files a demonstrate its common use case. You may also refer to the test file for more information

## Inspiration

This library is heavily inspired by [ts-pattern](https://github.com/gvergnaud/ts-pattern) which provides powerful pattern matching capabilities for TypeScript. The goal is to provide a similar experience in Go.
