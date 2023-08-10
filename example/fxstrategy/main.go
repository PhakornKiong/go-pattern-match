package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern"
)

const (
	USD  = "USD"
	SGD  = "SGD"
	EUR  = "EUR"
	BTC  = "BTC"
	ETH  = "ETH"
	USDT = "USDT"
	USDC = "USDC"
)

var FIAT_CURRENCIES = []string{USD, SGD, EUR}
var CRYPTOCURRENCIES = []string{BTC, ETH, USDT, USDC}
var STABLE_CRYPTOCURRENCIES = []string{USDT, USDC}

// Slice of currency pair
type CurrencyPair []string

func main() {
	// isCrypto := pattern.When[[]string](func(currency string) bool {
	// 	return pattern.Union(CRYPTOCURRENCIES...).Match(currency)
	// })

	isUsdStable := pattern.When(func(currency string) bool {
		return pattern.Union(STABLE_CRYPTOCURRENCIES...).Match(currency)
	})

	isFiat := pattern.When(func(currency string) bool {
		return pattern.Union(FIAT_CURRENCIES...).Match(currency)
	})

	isSameCurrency := pattern.When[CurrencyPair](func(currencies CurrencyPair) bool {
		return currencies[0] == currencies[1]
	})

	patternMatcher := func(input CurrencyPair) string {
		return pattern.NewMatcher[string, CurrencyPair](input).
			With(
				CurrencyPair{"BTC", "ETH"},
				func() string { return "Concrete BTC to ETH" },
			).
			With(
				[]pattern.AnyPatterner{isUsdStable, isUsdStable},
				func() string { return "both USD Stables strategy" },
			).
			With(
				[]pattern.AnyPatterner{isUsdStable, isFiat},
				func() string { return "USD Stables to fiat" },
			).
			With(
				isSameCurrency,
				func() string { return "same currency strategy" },
			).
			Otherwise(func() string { return "default strategy" })
	}

	fmt.Println(patternMatcher(CurrencyPair{BTC, ETH}))   // Concrete BTC to ETH
	fmt.Println(patternMatcher(CurrencyPair{USDT, SGD}))  // USD Stables to fiat
	fmt.Println(patternMatcher(CurrencyPair{USDT, USDT})) // both USD Stables strategy
	fmt.Println(patternMatcher(CurrencyPair{USDT, USDC})) // both USD Stables strategy
	fmt.Println(patternMatcher(CurrencyPair{BTC, BTC}))   // same currency strategy
	fmt.Println(patternMatcher(CurrencyPair{USD, SGD}))   // default strategy
}
