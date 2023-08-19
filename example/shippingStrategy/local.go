package main

// shopping by local partner
type localStrategy struct {
	distance int
	weight   int
}

func (s *localStrategy) CalculateCost() int {
	return s.distance * s.weight * 2
}

func NewLocalStrategy(distance, weight int) ShippingStrategy {
	return &localStrategy{distance, weight}
}
