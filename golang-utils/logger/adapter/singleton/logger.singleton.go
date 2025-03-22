package singleton

import (
	"github.com/bryanArroyave/eventsplit/back/common/logger/enums"
	"github.com/bryanArroyave/eventsplit/back/common/logger/factory"
	"github.com/bryanArroyave/eventsplit/back/common/logger/ports"
)

var (
	logger ports.ILogger
)

func InitLogger(loggerType enums.LoggerType, serviceName string) {
	logger = factory.NewLogger(loggerType, serviceName)
}

func GetLogger() ports.ILogger {
	return logger
}
