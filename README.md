# Go-Pattern

[![Test](https://github.com/PhakornKiong/go-pattern/actions/workflows/test.yml/badge.svg)](https://github.com/PhakornKiong/go-pattern/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/PhakornKiong/go-pattern/branch/master/graph/badge.svg?token=IL7G963OAF)](https://codecov.io/gh/PhakornKiong/go-pattern)
[![Go Report Card](https://goreportcard.com/badge/github.com/phakornkiong/go-pattern)](https://goreportcard.com/report/github.com/phakornkiong/go-pattern)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/PhakornKiong/go-pattern/blob/master/LICENSE)

Pattern Matching library for Go

```go
package main

import (
	"fmt"
	"github.com/phakornkiong/go-pattern/pattern"
)

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
	fmt.Println(FoodSorterWithPattern("carrot")) // "vegetable"
	fmt.Println(FoodSorterWithPattern("candy"))  // "unknown"
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

## `NewMatcher[T any, V any](input V) *Matcher[T, V]`

This function creates a new Matcher instance. The generics `T` and `V` represent any types.

`T` is the type that the handler function will return when a match is found.

`V` is the type of the input value that will be matched against the patterns.

## `.WithPattern(pattern Pattener, fn Handler[T]) *Matcher[T, V]`

This is a receiver method on a Matcher instance. It checks if the provided pattern matches the entire input. If a match is found, it calls the provided Handler function and return the response `T`.

## `.WithPatterns(patterns []Pattener, fn Handler[T]) *Matcher[T, V]`

This is a receiver method on a Matcher instance. It checks each of the provided patterns against <b>each of the input</b>. If a match is found, it calls the provided Handler function and return the response `T`.

## `.WithValue(value V, fn Handler[T]) *Matcher[T, V]`

This is a receiver method on a Matcher instance. It checks for deep equality between the provided value and the entire input. If a match is found, it calls the provided Handler function and return the response `T`.

## `.WithValues(values []V, fn Handler[T]) *Matcher[T, V]`

This is a receiver method on a Matcher instance. It checks for deep equality between each of the provided values against <b>each of the input</b>. If a match is found, it calls the provided Handler function and return the response `T`.

## `.Otherwise(fn Handler[T]) *Matcher[T, V]`

This is a receiver method on a Matcher instance. It is called when no match is found for the input. It calls the provided Handler function and return the response `T`.

## Examples

You can find more examples and usage scenarios in the `fxStrategy` and `switchUnion` files in the repository. Here are the direct links:

- [fxStrategy](https://github.com/PhakornKiong/go-pattern/blob/master/example/fxstrategy/main.go)
- [switchUnion](https://github.com/PhakornKiong/go-pattern/blob/master/example/switchunion/main.go)

These files a demonstrate its common use case. You may also refer to the test file for more information

## Inspiration

This library is heavily inspired by [ts-pattern](https://github.com/gvergnaud/ts-pattern) which provides powerful pattern matching capabilities for TypeScript. The goal is to provide a similar experience in Go.
