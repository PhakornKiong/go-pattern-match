package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern-match/pattern"
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
	isUsdStable := pattern.When(func(currency string) bool {
		return pattern.Union(STABLE_CRYPTOCURRENCIES...).Match(currency)
	})

	isFiat := pattern.When(func(currency string) bool {
		return pattern.Union(FIAT_CURRENCIES...).Match(currency)
	})

	isSameCurrency := pattern.When[CurrencyPair](func(currencies CurrencyPair) bool {
		return currencies[0] == currencies[1]
	})

	// Here we are defining patternMatcher based on currency pair
	// In real world, you would return the actual implementations of your abstraction
	patternMatcher := func(input CurrencyPair) string {
		return pattern.NewMatcher[string, CurrencyPair](input).
			// WithValues compare the currency pair by its indexes
			WithValues(
				CurrencyPair{"BTC", "ETH"},
				func() string { return "Concrete BTC to ETH" },
			).
			// WithPatterns will run pattern by index of the input slice
			WithPatterns(
				pattern.Patteners(isUsdStable, isUsdStable),
				func() string { return "Both USD Stables strategy" },
			).
			WithPatterns(
				[]pattern.Patterner{isUsdStable, isFiat},
				func() string { return "USD Stables to fiat" },
			).
			// WithPattern will run pattern on entire input
			WithPattern(
				isSameCurrency,
				func() string { return "Same currency strategy" },
			).
			// Default case
			Otherwise(func() string { return "default strategy" })
	}

	fmt.Println(patternMatcher(CurrencyPair{BTC, ETH}))   // Concrete BTC to ETH
	fmt.Println(patternMatcher(CurrencyPair{USDT, SGD}))  // USD Stables to fiat
	fmt.Println(patternMatcher(CurrencyPair{USDT, USDT})) // Both USD Stables strategy
	fmt.Println(patternMatcher(CurrencyPair{USDT, USDC})) // Both USD Stables strategy
	fmt.Println(patternMatcher(CurrencyPair{BTC, BTC}))   // Same currency strategy
	fmt.Println(patternMatcher(CurrencyPair{USD, SGD}))   // default strategy
}
