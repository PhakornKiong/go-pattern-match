package pattern

import "reflect"

type slicePattern[V any] struct {
	containsElement []V
	containsPattern []Patterner
	headElement     *V
	headPattern     *Patterner
	tailElement     *V
	tailPattern     *Patterner
}

func Slice[V any]() slicePattern[V] {
	return slicePattern[V]{}
}

func (s slicePattern[V]) clone() slicePattern[V] {
	return slicePattern[V]{
		containsElement: s.containsElement,
		containsPattern: s.containsPattern,
		headElement:     s.headElement,
		headPattern:     s.headPattern,
		tailElement:     s.tailElement,
		tailPattern:     s.tailPattern,
	}
}

func (s slicePattern[V]) Head(v V) slicePattern[V] {
	newPattern := s.clone()
	newPattern.headElement = &v
	return newPattern
}

func (s slicePattern[V]) HeadPattern(p Patterner) slicePattern[V] {
	newPattern := s.clone()
	newPattern.headPattern = &p
	return newPattern
}

func (s slicePattern[V]) Tail(v V) slicePattern[V] {
	newPattern := s.clone()
	newPattern.tailElement = &v
	return newPattern
}

func (s slicePattern[V]) TailPattern(p Patterner) slicePattern[V] {
	newPattern := s.clone()
	newPattern.tailPattern = &p
	return newPattern
}

func (s slicePattern[V]) Contains(v V) slicePattern[V] {
	newPattern := s.clone()
	newPattern.containsElement = append(newPattern.containsElement, v)
	return newPattern
}

func (s slicePattern[V]) ContainsPattern(p Patterner) slicePattern[V] {
	newPattern := s.clone()
	newPattern.containsPattern = append(newPattern.containsPattern, p)
	return newPattern
}

func (s slicePattern[V]) Match(value any) bool {
	// Implement match logic
	// Check if the value is of type slice[V]
	valueSlice, ok := value.([]V)
	if !ok {
		return false
	}

	if s.headElement != nil && !reflect.DeepEqual(valueSlice[0], *s.headElement) {
		return false
	}

	if s.headPattern != nil && !(*s.headPattern).Match(valueSlice[0]) {
		return false
	}

	if s.tailElement != nil && !reflect.DeepEqual(valueSlice[len(valueSlice)-1], *s.tailElement) {
		return false
	}

	if s.tailPattern != nil && !(*s.tailPattern).Match(valueSlice[len(valueSlice)-1]) {
		return false
	}

	// TODO: optimize this to avoid duplicate iteration

	for _, v := range s.containsElement {
		// Check if v is in valueSlice
		found := false
		for _, val := range valueSlice {
			if reflect.DeepEqual(val, v) {
				found = true
				break
			}
		}
		// If v is not found in valueSlice, return false
		if !found {
			return false
		}
	}

	for _, p := range s.containsPattern {
		// Check if p matches any element in valueSlice
		matched := false
		for _, val := range valueSlice {
			if p.Match(val) {
				matched = true
				break
			}
		}
		// If p does not match any element, return false
		if !matched {
			return false
		}
	}

	return true
}
