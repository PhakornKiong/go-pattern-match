package pattern

import (
	"reflect"
	"regexp"
	"strings"
)

type stringPattern struct {
	startsWith string
	endsWith   string
	minLength  int
	maxLength  int
	includes   string
	regex      *regexp.Regexp
}

func String() *stringPattern {
	return &stringPattern{}
}

func (s *stringPattern) clone() *stringPattern {
	return &stringPattern{
		startsWith: s.startsWith,
		endsWith:   s.endsWith,
		minLength:  s.minLength,
		maxLength:  s.maxLength,
		includes:   s.includes,
		regex:      s.regex,
	}
}

func (s *stringPattern) StartsWith(value string) *stringPattern {
	newPattern := s.clone()
	newPattern.startsWith = value
	return newPattern
}

func (s *stringPattern) EndsWith(value string) *stringPattern {
	newPattern := s.clone()
	newPattern.endsWith = value
	return newPattern
}

func (s *stringPattern) MinLength(value int) *stringPattern {
	newPattern := s.clone()
	newPattern.minLength = value
	return newPattern
}

func (s *stringPattern) MaxLength(value int) *stringPattern {
	newPattern := s.clone()
	newPattern.maxLength = value
	return newPattern
}

func (s *stringPattern) Includes(value string) *stringPattern {
	newPattern := s.clone()
	newPattern.includes = value
	return newPattern
}

func (s *stringPattern) Regex(value *regexp.Regexp) *stringPattern {
	newPattern := s.clone()
	newPattern.regex = value
	return newPattern
}

func (s stringPattern) Match(value any) bool {
	if reflect.TypeOf(value).Kind() != reflect.String {
		return false
	}

	str := value.(string)

	if s.startsWith != "" && !strings.HasPrefix(str, s.startsWith) {
		return false
	}

	if s.endsWith != "" && !strings.HasSuffix(str, s.endsWith) {
		return false
	}

	if s.minLength != 0 && len(str) < s.minLength {
		return false
	}

	if s.maxLength != 0 && len(str) > s.maxLength {
		return false
	}

	if s.includes != "" && !strings.Contains(str, s.includes) {
		return false
	}

	if s.regex != nil && !s.regex.MatchString(str) {
		return false
	}

	return true
}
