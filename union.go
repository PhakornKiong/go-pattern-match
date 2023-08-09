package pattern

import (
	"reflect"
)

type union[V any] struct {
	patterns []V
}

func Union[V any](patterns ...V) union[V] {
	return union[V]{patterns: patterns}
}

func (u union[V]) Match(value any) bool {
	for _, subPattern := range u.patterns {
		if reflect.DeepEqual(value, subPattern) {
			return true
		}
	}
	return false
}

type unionPattern[V AnyPatterner] struct {
	patterns []V
}

func UnionPattern[V AnyPatterner](patterns ...V) unionPattern[V] {
	return unionPattern[V]{patterns: patterns}
}

func (u unionPattern[V]) Match(value any) bool {
	for _, subPattern := range u.patterns {
		if subPattern.Match(value) {
			return true
		}
	}
	return false
}
