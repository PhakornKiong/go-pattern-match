package pattern

import (
	"reflect"
)

type intersection[V any] struct {
	patterns []V
}

func Intersection[V any](patterns ...V) intersection[V] {
	return intersection[V]{patterns: patterns}
}

func (i intersection[V]) Match(value any) bool {
	for _, subPattern := range i.patterns {
		if !reflect.DeepEqual(value, subPattern) {
			return false
		}
	}
	return true
}

type intersectionPattern[V Patterner] struct {
	patterns []V
}

func IntersectionPattern[V Patterner](patterns ...V) intersectionPattern[V] {
	return intersectionPattern[V]{patterns: patterns}
}

func (u intersectionPattern[V]) Match(value any) bool {
	for _, subPattern := range u.patterns {
		if !subPattern.Match(value) {
			return false
		}
	}
	return true
}
