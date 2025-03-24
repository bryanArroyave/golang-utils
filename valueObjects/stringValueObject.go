package valueobjects

import (
	"regexp"

	customerrors "github.com/bryanArroyave/golang-utils/valueObjects/customErrors"
)

type StringValueObject struct {
	*BaseValueObject[string]
	maxLength        int
	hasMaxValidation bool
	minLength        int
	hasMinValidation bool
	pattern          string
	hasPattern       bool
	includeValues    []string
}

func NewStringValueObject(value string) *StringValueObject {
	v := &StringValueObject{}

	base := &BaseValueObject[string]{
		value:    &value,
		validate: v.validate,
	}

	v.BaseValueObject = base

	return v
}

func (s *StringValueObject) MaxLength(length int) *StringValueObject {
	s.maxLength = length
	s.hasMaxValidation = true
	return s
}

func (s *StringValueObject) MinLength(length int) *StringValueObject {
	s.minLength = length
	s.hasMinValidation = true
	return s
}

func (s *StringValueObject) Pattern(pattern string) *StringValueObject {
	s.pattern = pattern
	s.hasPattern = true
	return s
}

func (s *StringValueObject) Optional() *StringValueObject {
	s.optional = true
	return s
}

func (s *StringValueObject) Include(values []string) *StringValueObject {
	s.includeValues = values
	return s
}

func (s *StringValueObject) validate() {
	s.validateMaxLength()
	s.validateMinLength()
	s.validatePattern()
	s.validateInclude()
}

func (s *StringValueObject) validateMaxLength() {

	if *s.value == "" && s.optional {
		return
	}

	if s.hasMaxValidation {
		if len(*s.value) > s.maxLength {
			s.errors = append(s.errors, customerrors.NewMaxLengthError(s.maxLength))
		}
	}
}

func (s *StringValueObject) validateMinLength() {

	if *s.value == "" && s.optional {
		return
	}

	if s.hasMinValidation {
		if len(*s.value) < s.minLength {
			s.errors = append(s.errors, customerrors.NewMinLengthError(s.minLength))
		}
	}
}

func (s *StringValueObject) validatePattern() {

	if *s.value == "" && s.optional {
		return
	}

	if s.hasPattern {
		match, err := regexp.MatchString(s.pattern, *s.value)

		if err != nil {
			s.errors = append(s.errors, err)
		}

		if !match {
			s.errors = append(s.errors, customerrors.NewPatternError())
		}
	}
}

func (s *StringValueObject) validateInclude() {

	if *s.value == "" && s.optional {
		return
	}

	if len(s.includeValues) > 0 {
		include := false
		for _, value := range s.includeValues {
			if *s.value == value {
				include = true
				break
			}
		}

		if !include {
			s.errors = append(s.errors, customerrors.NewIncludeError(s.includeValues))
		}
	}
}
