package singleton

import (
	"github.com/bryanArroyave/golang-utils/logger/enums"
	"github.com/bryanArroyave/golang-utils/logger/factory"
	"github.com/bryanArroyave/golang-utils/logger/ports"
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
