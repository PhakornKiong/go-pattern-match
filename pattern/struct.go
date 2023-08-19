package pattern

import (
	"reflect"
)

type fieldVal struct {
	field string
	val   any
}

type fieldPattern struct {
	field   string
	pattern Patterner
}

type structPattern struct {
	fieldValues   []fieldVal
	fieldPatterns []fieldPattern
}

func Struct() structPattern {
	return structPattern{}
}

func (s structPattern) clone() structPattern {
	return structPattern{
		fieldValues:   s.fieldValues,
		fieldPatterns: s.fieldPatterns,
	}
}

func (m structPattern) FieldValue(fieldName string, v any) structPattern {
	newPattern := m.clone()
	newPattern.fieldValues = append(newPattern.fieldValues, fieldVal{fieldName, v})
	return newPattern
}

func (m structPattern) FieldPattern(fieldName string, p Patterner) structPattern {
	newPattern := m.clone()
	newPattern.fieldPatterns = append(newPattern.fieldPatterns, fieldPattern{fieldName, p})
	return newPattern
}

func (m structPattern) Match(value any) bool {
	v := reflect.ValueOf(value)

	// Check if it is struct
	if v.Kind() != reflect.Struct {
		return false
	}

	for _, fv := range m.fieldValues {
		value, ok := getFieldValue(v, fv.field)

		if !ok {
			return false
		}

		if !reflect.DeepEqual(value, fv.val) {
			return false
		}
	}

	for _, fp := range m.fieldPatterns {
		value, ok := getFieldValue(v, fp.field)

		if !ok {
			return false
		}

		if !fp.pattern.Match(value) {
			return false
		}
	}
	return true
}

func getFieldValue(v reflect.Value, fieldName string) (any, bool) {
	field := v.FieldByName(fieldName)

	// Check if field exists
	if !field.IsValid() {
		return nil, false
	}

	// Check if field is exported
	if !field.CanInterface() {
		return nil, false
	}

	return field.Interface(), true

}
