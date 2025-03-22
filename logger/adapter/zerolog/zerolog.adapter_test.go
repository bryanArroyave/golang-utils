package zerolog

import (
	"errors"
	"testing"

	"github.com/bryanArroyave/golang-utils/logger/dtos"
)

func TestPrueb(t *testing.T) {

	logger := NewZerologLoggerAdapter("ms-prueba")

	tests := []struct {
		message string
		fields  []*dtos.LoggerFieldsDTO
	}{
		{
			message: "hola mundo",
			fields: []*dtos.LoggerFieldsDTO{
				{Key: "key_string", Value: "value_string"},
				{Key: "key_int", Value: 1},
				{Key: "key_int64", Value: int64(1)},
				{Key: "key_float64", Value: float64(1.1)},
				{Key: "key_bool", Value: true},
				{Key: "key_error", Value: errors.New("error")},
				{Key: "key_interface", Value: struct {
					Foo string
					Bar string
				}{
					Foo: "foo",
					Bar: "bar",
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			logger.Info(tt.message, tt.fields...)
			logger.Error(tt.message, tt.fields...)
			logger.Warn(tt.message, tt.fields...)
		})
	}

}
