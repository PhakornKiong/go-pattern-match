package pattern

type Predicate[V any] func(V) bool

type whenPattern[V any] struct {
	predicate Predicate[V]
}

func When[V any](predicate Predicate[V]) whenPattern[V] {
	return whenPattern[V]{predicate: predicate}
}

func (w whenPattern[V]) Match(value any) bool {
	val, ok := value.(V)
	if !ok {
		return false
	}
	return w.predicate(val)
}
