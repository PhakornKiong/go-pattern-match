package main

// default strategy, just take flat rate of 25
type defaultStrategy struct{}

func (s *defaultStrategy) CalculateCost() int {
	return 25
}

func NewDefaultStrategy() ShippingStrategy {
	return &defaultStrategy{}
}
