package main

// shipping by freighter
type freightStrategy struct {
	distance int
	weight   int
	volume   int
}

func (s *freightStrategy) CalculateCost() int {
	return s.distance * s.weight * s.volume
}

func NewFreightStrategy(distance, weight, volume int) ShippingStrategy {
	return &freightStrategy{distance, weight, volume}
}
