package pattern

import "reflect"

type not struct {
	pattern any
}

func Not(pattern any) not {
	return not{pattern: pattern}
}

func (n not) Match(value any) bool {
	return !reflect.DeepEqual(value, n.pattern)
}

type notPattern[V AnyPatterner] struct {
	pattern V
}

func NotPattern[V AnyPatterner](pattern V) notPattern[V] {
	return notPattern[V]{pattern: pattern}
}

func (n notPattern[V]) Match(value any) bool {
	return n.pattern.Match(value)
}
