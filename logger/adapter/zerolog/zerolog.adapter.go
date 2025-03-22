package zerolog

import (
	"github.com/bryanArroyave/eventsplit/back/common/logger/dtos"
	"github.com/bryanArroyave/eventsplit/back/common/logger/ports"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ZerologLoggerAdapter struct {
	logger *zerolog.Logger
}

func NewZerologLoggerAdapter(serviceName string) ports.ILogger {

	_logger := log.With().Str("ms-name", serviceName).Logger()

	return &ZerologLoggerAdapter{
		logger: &_logger,
	}
}

func (z *ZerologLoggerAdapter) Info(message string, fields ...*dtos.LoggerFieldsDTO) {
	logger := z.logger.Info()
	z.print(logger, message, fields...)
}

func (z *ZerologLoggerAdapter) Error(message string, fields ...*dtos.LoggerFieldsDTO) {
	logger := z.logger.Error()
	z.print(logger, message, fields...)
}

func (z *ZerologLoggerAdapter) Warn(message string, fields ...*dtos.LoggerFieldsDTO) {
	logger := z.logger.Warn()
	z.print(logger, message, fields...)
}

func (z *ZerologLoggerAdapter) print(logger *zerolog.Event, message string, fields ...*dtos.LoggerFieldsDTO) {

	for _, value := range fields {
		logger = mapValueEvent(value.Key, value.Value, logger)
	}

	logger.Msg(message)
}

func mapValueEvent(key string, value interface{}, logger *zerolog.Event) *zerolog.Event {
	switch v := value.(type) {
	case string:
		return logger.Str(key, v)
	case int:
		return logger.Int(key, v)
	case int64:
		return logger.Int64(key, v)
	case float64:
		return logger.Float64(key, v)
	case bool:
		return logger.Bool(key, v)
	case error:
		return logger.Err(v)
	default:
		return logger.Interface(key, v)
	}
}
