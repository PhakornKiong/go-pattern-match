package pattern

import "reflect"

type unionPattern[V any] struct {
	patterns []V
}

func Union[V any](patterns ...V) unionPattern[V] {
	return unionPattern[V]{patterns: patterns}
}

func (u *unionPattern[V]) Match(value V) bool {
	for _, subPattern := range u.patterns {
		if reflect.DeepEqual(value, subPattern) {
			return true
		}
	}
	return false
}
