package main

import (
	"fmt"

	"github.com/phakornkiong/go-pattern-match/pattern"
)

type ShippingStrategy interface {
	CalculateCost() int
}

type Order struct {
	Country  string
	Distance int
	Weight   int
	Volume   int
}

const (
	MY = "Malaysia"
	AU = "Australia"
	US = "US"
	CN = "China"
)

func shippingStrategyFactory(o Order) ShippingStrategy {
	switch o.Country {
	case MY:
		return NewLocalStrategy(o.Distance, o.Weight)
	case US, AU, CN:
		if o.Volume > 100 || o.Weight > 250 {
			return NewFreightStrategy(o.Distance, o.Weight, o.Volume)
		}
		return NewAirStrategy(o.Distance, o.Weight, o.Volume)
	default:
		return NewDefaultStrategy()
	}
}

// This is more declarative, and can be unit tested separately
// Albeit it is much verbose

var LargeVolumePattern = pattern.Int().Gt(100)
var LargeWeightPattern = pattern.Int().Gt(250)
var IsOverseasPattern = pattern.Union[string](US, AU, CN)

var airPattern = pattern.Struct().
	FieldPattern("Country", IsOverseasPattern)

var freightVolPattern = pattern.Struct().
	FieldPattern("Country", IsOverseasPattern).
	FieldPattern("Volume", LargeVolumePattern)

var freightWeightPattern = pattern.Struct().
	FieldPattern("Country", IsOverseasPattern).
	FieldPattern("Weight", LargeWeightPattern)

var freightPattern = pattern.UnionPattern(freightVolPattern, freightWeightPattern)

func shippingStrategyFactoryPattern(o Order) ShippingStrategy {
	return pattern.NewMatcher[ShippingStrategy](o).
		WithPattern(
			pattern.Struct().FieldValue("Country", MY),
			func() ShippingStrategy {
				return NewLocalStrategy(o.Distance, o.Weight)
			},
		).
		WithPattern(
			freightPattern,
			func() ShippingStrategy {
				return NewFreightStrategy(o.Distance, o.Weight, o.Volume)
			},
		).
		WithPattern(
			airPattern,
			func() ShippingStrategy {
				return NewAirStrategy(o.Distance, o.Weight, o.Volume)
			},
		).
		Otherwise(func() ShippingStrategy { return NewDefaultStrategy() })
}

// o.Volume > 100 || o.Weight > 250
func main() {
	freightWeightOrder := Order{AU, 100, 251, 99}
	freightVolOrder := Order{AU, 100, 249, 101}
	airOrder := Order{US, 1, 1, 1}
	localOrder := Order{Country: MY, Distance: 2, Weight: 5}
	defaultOrder := Order{Country: "Singapore"}

	// *main.FreightStrategy
	fmt.Printf("%T\n", shippingStrategyFactory(freightWeightOrder))
	// *main.FreightStrategy
	fmt.Printf("%T\n", shippingStrategyFactory(freightVolOrder))
	// *main.AirStrategy
	fmt.Printf("%T\n", shippingStrategyFactory(airOrder))
	// *main.LocalStrategy
	fmt.Printf("%T\n", shippingStrategyFactory(localOrder))
	// *main.DefaultStrategy
	fmt.Printf("%T\n", shippingStrategyFactory(defaultOrder))

	// *main.FreightStrategy
	fmt.Printf("%T\n", shippingStrategyFactoryPattern(freightWeightOrder))
	// *main.FreightStrategy
	fmt.Printf("%T\n", shippingStrategyFactoryPattern(freightVolOrder))
	// *main.AirStrategy
	fmt.Printf("%T\n", shippingStrategyFactoryPattern(airOrder))
	// *main.LocalStrategy
	fmt.Printf("%T\n", shippingStrategyFactoryPattern(localOrder))
	// *main.DefaultStrategy
	fmt.Printf("%T\n", shippingStrategyFactoryPattern(defaultOrder))
}
