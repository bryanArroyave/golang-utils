package dtos

import (
	"github.com/bryanArroyave/golang-utils/logger/enums"
	mongodtos "github.com/bryanArroyave/golang-utils/mongo/dtos"
)

type AppConfigDTO struct {
	MongoConnection *mongodtos.MongoConnectionDTO
	LoggerConfig    *LoggerConfigDTO
}

type LoggerConfigDTO struct {
	LoggerType  enums.LoggerType
	ServiceName string
}
