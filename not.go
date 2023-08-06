package pattern

import "reflect"

type notPattern struct {
	pattern any
}

func Not(pattern any) notPattern {
	return notPattern{pattern: pattern}
}

func (n *notPattern) Match(value any) bool {
	return !reflect.DeepEqual(value, n.pattern)
}
