package valueobjects

import (
	customerrors "github.com/bryanArroyave/golang-utils/valueObjects/customErrors"
)

type FloatValueObject struct {
	*BaseValueObject[float64]
	minValue         float64
	hasMinValidation bool
	maxValue         float64
	hasMaxValidation bool
}

func NewFloatValueObject(value float64) *FloatValueObject {
	v := &FloatValueObject{}

	base := &BaseValueObject[float64]{
		value:    &value,
		validate: v.validate,
	}

	v.BaseValueObject = base

	return v
}

func (s *FloatValueObject) Max(value float64) *FloatValueObject {
	s.maxValue = value
	s.hasMaxValidation = true
	return s
}

func (s *FloatValueObject) Min(value float64) *FloatValueObject {

	s.minValue = value
	s.hasMinValidation = true
	return s
}

func (s *FloatValueObject) Optional() *FloatValueObject {
	s.optional = true
	return s
}

func (s *FloatValueObject) validate() {
	s.validateMax()
	s.validateMin()
}

func (s *FloatValueObject) validateMin() {

	if *s.value == 0 && s.optional {
		return
	}
	if s.hasMinValidation {
		if *s.value < s.minValue {
			s.errors = append(s.errors, customerrors.NewMinError(s.minValue))
		}
	}
}

func (s *FloatValueObject) validateMax() {

	if *s.value == 0 && s.optional {
		return
	}
	if s.hasMaxValidation {
		if *s.value > s.maxValue {
			s.errors = append(s.errors, customerrors.NewMaxError(s.maxValue))
		}
	}
}
