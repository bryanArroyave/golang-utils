package postgres

import (
	gormconnect "github.com/bryanArroyave/golang-utils/gorm"
	"github.com/bryanArroyave/golang-utils/gorm/dtos"
	"github.com/bryanArroyave/golang-utils/gorm/ports"
	gormpostgres "gorm.io/driver/postgres"
)

type PostgresDBManager struct {
	dbManager *gormconnect.DBManager
}

func NewPostgresDBManager(options *dtos.ConnectionDTO) ports.IDBManager {
	dialector := gormpostgres.New(gormpostgres.Config{
		DSN:                  options.URI,
		PreferSimpleProtocol: true,
	})
	dbManager := gormconnect.NewDBManager(&dtos.GormConnectionDTO{
		ConnectionDTO: options,
		Dialector:     &(dialector),
	})
	return dbManager
}
