package app

import (
	"context"

	appdtos "github.com/bryanArroyave/golang-utils/app/dtos"
	"github.com/bryanArroyave/golang-utils/events/enums"
	eventsfactory "github.com/bryanArroyave/golang-utils/events/factory"
	"github.com/bryanArroyave/golang-utils/gorm/dtos"
	"github.com/bryanArroyave/golang-utils/gorm/ports"
	"github.com/bryanArroyave/golang-utils/gorm/postgres"
	"github.com/bryanArroyave/golang-utils/logger/adapter/singleton"
	loggerports "github.com/bryanArroyave/golang-utils/logger/ports"
	mongoadapter "github.com/bryanArroyave/golang-utils/mongo"
	mongodtos "github.com/bryanArroyave/golang-utils/mongo/dtos"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	logger            loggerports.ILogger
	mongoInstances    map[string]*mongo.Database
	postgresInstances map[string]ports.IDBManager
	messageBrokers    map[string]*eventsfactory.MessageBroker
}

func NewApp(config *appdtos.LoggerConfigDTO) *App {

	app := &App{}
	app.initLogger(config)
	return app
}

func (a *App) initLogger(config *appdtos.LoggerConfigDTO) {
	singleton.InitLogger(config.LoggerType, config.ServiceName)
	a.logger = singleton.GetLogger()
}

func (a *App) AddMessageBroker(name string, adapterType enums.BrokerType, config *eventsfactory.FactoryConfig) *App {
	if a.messageBrokers == nil {
		a.messageBrokers = make(map[string]*eventsfactory.MessageBroker)
	}

	broker, err := eventsfactory.NewMessageBroker(adapterType, config)
	if err != nil {
		panic(err)
	}

	a.messageBrokers[name] = broker
	return a
}

func (a *App) AddMongoConnection(name string, config *mongodtos.MongoConnectionDTO) *App {
	if a.mongoInstances == nil {
		a.mongoInstances = make(map[string]*mongo.Database)
	}

	db, err := mongoadapter.InitMongoDB(context.TODO(), config)
	if err != nil {
		panic(err)
	}

	a.mongoInstances[name] = db
	return a

}

func (a *App) AddPostgresConnection(name string, config *dtos.ConnectionDTO) *App {
	if a.postgresInstances == nil {
		a.postgresInstances = make(map[string]ports.IDBManager)
	}

	db := postgres.NewPostgresDBManager(config)

	a.postgresInstances[name] = db

	return a
}

func (a *App) Close() {
	for _, d := range a.postgresInstances {
		d.Close()
	}

}

// GETTERS
func (a *App) GetLogger() loggerports.ILogger {
	return a.logger
}

func (a *App) GetMongoConnection(name string) *mongo.Database {
	return a.mongoInstances[name]
}

func (a *App) GetPostgresConnection(name string) ports.IDBManager {
	return a.postgresInstances[name]
}

func (a *App) GetMessageBroker(name string) *eventsfactory.MessageBroker {
	return a.messageBrokers[name]
}
