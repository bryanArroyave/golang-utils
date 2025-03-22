package dtos

import (
	"github.com/bryanArroyave/eventsplit/back/common/logger/enums"
	mongodtos "github.com/bryanArroyave/eventsplit/back/common/mongo/dtos"
)

type AppConfigDTO struct {
	MongoConnection *mongodtos.MongoConnectionDTO
	LoggerConfig    *LoggerConfigDTO
}

type LoggerConfigDTO struct {
	LoggerType  enums.LoggerType
	ServiceName string
}
