package factory

import (
	"github.com/bryanArroyave/golang-utils/logger/adapter/fmt"
	"github.com/bryanArroyave/golang-utils/logger/adapter/zerolog"
	"github.com/bryanArroyave/golang-utils/logger/enums"
	"github.com/bryanArroyave/golang-utils/logger/ports"
)

func NewLogger(loggerType enums.LoggerType, serviceName string) ports.ILogger {

	switch loggerType {
	case enums.Zerolog:
		return zerolog.NewZerologLoggerAdapter(serviceName)
	default:
		return fmt.NewFmtLoggerAdapter(serviceName)
	}

}
