package factory

import (
	"github.com/bryanArroyave/eventsplit/back/common/logger/adapter/fmt"
	"github.com/bryanArroyave/eventsplit/back/common/logger/adapter/zerolog"
	"github.com/bryanArroyave/eventsplit/back/common/logger/enums"
	"github.com/bryanArroyave/eventsplit/back/common/logger/ports"
)

func NewLogger(loggerType enums.LoggerType, serviceName string) ports.ILogger {

	switch loggerType {
	case enums.Zerolog:
		return zerolog.NewZerologLoggerAdapter(serviceName)
	default:
		return fmt.NewFmtLoggerAdapter(serviceName)
	}

}
