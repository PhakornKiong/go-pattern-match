package main

// shipping by freighter
type airStrategy struct {
	distance int
	weight   int
	volume   int
}

func (s *airStrategy) CalculateCost() int {
	return s.distance * s.weight * s.volume * 5
}

func NewAirStrategy(distance, weight, volume int) ShippingStrategy {
	return &airStrategy{distance, weight, volume}
}
