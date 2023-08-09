package pattern

import "reflect"

type intPattern struct {
	between []int
	lt      int
	gt      int
	lte     int
	gte     int
	isPos   bool
	isNeg   bool
}

func Int() intPattern {
	return intPattern{}
}

func (n intPattern) Between(min, max int) intPattern {
	newPattern := n
	newPattern.between = []int{min, max}
	return newPattern
}

func (n intPattern) Lt(max int) intPattern {
	newPattern := n
	newPattern.lt = max
	return newPattern
}

func (n intPattern) Gt(min int) intPattern {
	newPattern := n
	newPattern.gt = min
	return newPattern
}

func (n intPattern) Lte(max int) intPattern {
	newPattern := n
	newPattern.lte = max
	return newPattern
}

func (n intPattern) Gte(min int) intPattern {
	newPattern := n
	newPattern.gte = min
	return newPattern
}

func (n intPattern) Positive() intPattern {
	newPattern := n
	newPattern.isPos = true
	return newPattern
}

func (n intPattern) Negative() intPattern {
	newPattern := n
	newPattern.isNeg = true
	return newPattern
}

func (n intPattern) Match(value any) bool {
	if reflect.TypeOf(value).Kind() != reflect.Int {
		return false
	}

	input := value.(int)

	if n.between != nil && (input < n.between[0] || input > n.between[1]) {
		return false
	}
	if n.lt != 0 && input >= n.lt {
		return false
	}
	if n.gt != 0 && input <= n.gt {
		return false
	}
	if n.lte != 0 && input > n.lte {
		return false
	}
	if n.gte != 0 && input < n.gte {
		return false
	}

	if n.isPos && input <= 0 {
		return false
	}
	if n.isNeg && input >= 0 {
		return false
	}
	return true
}
