package pattern

import "reflect"

type notPattern[V any] struct {
	pattern V
}

func Not[V any](pattern V) notPattern[V] {
	return notPattern[V]{pattern: pattern}
}

func (n *notPattern[V]) Match(value V) bool {
	return !reflect.DeepEqual(value, n.pattern)
}
