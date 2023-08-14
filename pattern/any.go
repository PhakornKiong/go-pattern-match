package pattern

type anyPattern struct {
}

func Any() anyPattern {
	return anyPattern{}
}

func (a anyPattern) Match(value any) bool {
	return true
}
