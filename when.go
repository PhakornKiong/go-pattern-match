package pattern

import "fmt"

type Predicate[V any] func(V) bool

type whenPattern[V any] struct {
	predicate Predicate[V]
}

func When[V any](predicate Predicate[V]) whenPattern[V] {
	return whenPattern[V]{predicate: predicate}
}

func (w *whenPattern[V]) Match(value V) bool {
	fmt.Println(value)
	fmt.Println(w.predicate(value))
	return w.predicate(value)
}
