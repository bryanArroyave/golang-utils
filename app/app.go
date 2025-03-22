package app

import (
	"context"
	"fmt"

	appdtos "github.com/bryanArroyave/golang-utils/app/dtos"
	"github.com/bryanArroyave/golang-utils/logger/adapter/singleton"
	"github.com/bryanArroyave/golang-utils/logger/ports"
	mongoadapter "github.com/bryanArroyave/golang-utils/mongo"
	"github.com/bryanArroyave/golang-utils/mongo/dtos"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	logger         ports.ILogger
	mongoInstances map[string]*mongo.Database
}

func NewApp(config *appdtos.LoggerConfigDTO) *App {

	err := godotenv.Load()

	fmt.Println(err)

	app := &App{}
	app.initLogger(config)
	return app
}

func (a *App) initLogger(config *appdtos.LoggerConfigDTO) {
	singleton.InitLogger(config.LoggerType, config.ServiceName)
	a.logger = singleton.GetLogger()
}

func (a *App) AddMongoConnection(name string, config *dtos.MongoConnectionDTO) *App {
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

// GETTERS
func (a *App) GetLogger() ports.ILogger {
	return a.logger
}

func (a *App) GetMongoConnection(name string) *mongo.Database {
	return a.mongoInstances[name]
}
