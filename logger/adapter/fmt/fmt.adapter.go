package fmt

import (
	"fmt"

	"github.com/bryanArroyave/golang-utils/logger/dtos"
	"github.com/bryanArroyave/golang-utils/logger/ports"
)

type FmtLoggerAdapter struct {
	serviceName string
}

func NewFmtLoggerAdapter(serviceName string) ports.ILogger {

	return &FmtLoggerAdapter{
		serviceName: serviceName,
	}
}

func (z *FmtLoggerAdapter) Info(message string, fields ...*dtos.LoggerFieldsDTO) {
	msg := ""

	for _, field := range fields {
		msg += fmt.Sprintf("%s: %v, ", field.Key, field.Value)
	}

	fmt.Printf("Info - %s: %s - %s\n", z.serviceName, message, msg)
}

func (z *FmtLoggerAdapter) Error(message string, fields ...*dtos.LoggerFieldsDTO) {
	msg := ""

	for _, field := range fields {
		msg += fmt.Sprintf("%s: %v, ", field.Key, field.Value)
	}

	fmt.Printf("Error - %s: %s - %s\n", z.serviceName, message, msg)
}

func (z *FmtLoggerAdapter) Warn(message string, fields ...*dtos.LoggerFieldsDTO) {
	msg := ""

	for _, field := range fields {
		msg += fmt.Sprintf("%s: %v, ", field.Key, field.Value)
	}

	fmt.Printf("Warn - %s: %s - %s\n", z.serviceName, message, msg)
}
