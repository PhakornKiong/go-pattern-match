# Go-Pattern

Pattern Matching library for Go

## Key Components

- `Pattener`: This is an interface that requires the implementation of a `Match` function. Any type that implements this interface can be used as a pattern in the matcher.

- `Handler`: This is a function type that returns a generic type `T`. This function is called when a match is found.

- `Matcher`: This is a struct that holds the value to be matched, a flag indicating if a match has been found, and the response to be returned when a match is found.

## Functions

- `NewMatcher`: This function creates a new Matcher instance.

- `With`: This is a method on the Matcher type. It takes a pattern and a Handler function. If the pattern matches the value in the Matcher, the Handler function is called and the response is set.

- `Otherwise`: This is a method on the Matcher type. It takes a Handler function and calls it if no match has been found.

## Usage

To use the pattern matcher, create a new Matcher with the value to be matched. Then call the `With` method with the pattern and the Handler function. If no match is found, call the `Otherwise` method with a default Handler function.

## Inspiration

This library is heavily inspired by [ts-pattern](https://github.com/gvergnaud/ts-pattern) which provides powerful pattern matching capabilities for TypeScript. The goal is to provide a similar experience in Go.
