package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern/pattern"
)

func Fib(input int) (output int) {
	output = pattern.NewMatcher[int](input).
		WithValue(
			1,
			func() int { return 1 },
		).
		WithValue(
			2,
			func() int { return 1 },
		).
		Otherwise(func() int { return Fib(input-1) + Fib(input-2) })

	return output
}

func main() {
	fmt.Println(Fib(1))  // 1
	fmt.Println(Fib(2))  // 1
	fmt.Println(Fib(10)) // 55
	fmt.Println(Fib(12)) // 144
}
