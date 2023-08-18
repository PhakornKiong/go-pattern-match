package pattern

import (
	"reflect"
)

type mapPattern[K comparable, V any] struct {
	keyVals        []keyVal[K, V]
	keys           []K
	vals           []V
	keyValPatterns []keyVal[K, Patterner]
}

type keyVal[K comparable, V any] struct {
	key K
	val V
}

func Map[K comparable, V any]() mapPattern[K, V] {
	return mapPattern[K, V]{}
}

func (s mapPattern[K, V]) clone() mapPattern[K, V] {
	return mapPattern[K, V]{
		keyVals:        s.keyVals,
		keys:           s.keys,
		vals:           s.vals,
		keyValPatterns: s.keyValPatterns,
	}
}

func (m mapPattern[K, V]) KeyVal(key K, val V) mapPattern[K, V] {
	newPattern := m.clone()
	newPattern.keyVals = append(newPattern.keyVals, keyVal[K, V]{key, val})
	return newPattern
}

func (m mapPattern[K, V]) Key(key K) mapPattern[K, V] {
	newPattern := m.clone()
	newPattern.keys = append(newPattern.keys, key)
	return newPattern
}

func (m mapPattern[K, V]) Val(val V) mapPattern[K, V] {
	newPattern := m.clone()
	newPattern.vals = append(newPattern.vals, val)
	return newPattern
}

func (m mapPattern[K, V]) KeyValPatterns(key K, p Patterner) mapPattern[K, V] {
	newPattern := m.clone()
	newPattern.keyValPatterns = append(newPattern.keyValPatterns, keyVal[K, Patterner]{key, p})
	return newPattern
}

func (m mapPattern[K, V]) Match(value any) bool {
	input, ok := value.(map[K]V)

	if !ok {
		return false
	}

	// Check if key and value pair exists in the input map
	for _, kv := range m.keyVals {
		val, ok := input[kv.key]

		if !ok {
			return false
		}

		if !reflect.DeepEqual(val, kv.val) {
			return false
		}
	}

	// Check if key exists in the input map
	for _, k := range m.keys {
		_, ok := input[k]
		if !ok {
			return false
		}
	}

	// Check if value exists in the input map
	for _, v := range m.vals {
		found := false
		for _, kv := range input {
			if reflect.DeepEqual(kv, v) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// Check if value in the key matches the provided pattern
	for _, kv := range m.keyValPatterns {
		val, ok := input[kv.key]

		if !ok {
			return false
		}

		if !kv.val.Match(val) {
			return false
		}
	}

	return true
}
