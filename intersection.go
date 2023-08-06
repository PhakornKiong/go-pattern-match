package pattern

import "reflect"

type intersectionPattern[V any] struct {
	patterns []V
}

func Intersection[V any](patterns ...V) intersectionPattern[V] {
	return intersectionPattern[V]{patterns: patterns}
}

func (i intersectionPattern[V]) Match(value V) bool {
	for _, subPattern := range i.patterns {
		if !reflect.DeepEqual(value, subPattern) {
			return false
		}
	}
	return true
}
