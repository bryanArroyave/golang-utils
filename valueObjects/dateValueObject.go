package valueobjects

import (
	"time"

	customerrors "github.com/bryanArroyave/golang-utils/valueObjects/customErrors"
)

type DateValueObject struct {
	*BaseValueObject[time.Time]
	minValue         time.Time
	hasMinValidation bool
	maxValue         time.Time
	hasMaxValidation bool
}

func NewDateValueObject(value time.Time) *DateValueObject {
	v := &DateValueObject{}

	base := &BaseValueObject[time.Time]{
		value:    &value,
		validate: v.validate,
	}

	v.BaseValueObject = base

	return v
}

func (s *DateValueObject) Max(value time.Time) *DateValueObject {
	s.maxValue = value
	s.hasMaxValidation = true
	return s
}

func (s *DateValueObject) Min(value time.Time) *DateValueObject {

	s.minValue = value
	s.hasMinValidation = true
	return s
}

func (s *DateValueObject) Optional() *DateValueObject {
	s.optional = true
	return s
}

func (s *DateValueObject) validate() {
	s.validateMax()
	s.validateMin()
}

func (s *DateValueObject) validateMin() {

	zeroTime := time.Time{}
	if *s.value == zeroTime && s.optional {
		return
	}
	if s.hasMinValidation {
		if s.value.Before(s.minValue) {
			s.errors = append(s.errors, customerrors.NewMinError(s.minValue))
		}
	}
}

func (s *DateValueObject) validateMax() {
	zeroTime := time.Time{}

	if *s.value == zeroTime && s.optional {
		return
	}
	if s.hasMaxValidation {
		if s.value.After(s.maxValue) {
			s.errors = append(s.errors, customerrors.NewMaxError(s.maxValue))
		}
	}
}
