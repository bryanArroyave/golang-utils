package customerrors

import "fmt"

type PatternError struct {
}

func NewPatternError() *PatternError {
	return &PatternError{}
}

func (e *PatternError) Error() string {
	return fmt.Sprintf("invalid value")
}
