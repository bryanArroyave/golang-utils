package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/bryanArroyave/golang-utils/mongo/dtos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(ctx context.Context, config *dtos.MongoConnectionDTO) (*mongo.Database, error) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.User, config.Password, config.Host, config.Port)

	connectionURI := fmt.Sprintf("%s/?retryWrites=true&w=majority&authSource=admin&authMechanism=SCRAM-SHA-256", mongoURI)

	clientOptions := options.Client().ApplyURI(connectionURI)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(10)
	clientOptions.SetMaxConnIdleTime(10 * time.Minute)
	clientOptions.SetMaxConnecting(10)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	mongoDB := client.Database(config.DBName)

	return mongoDB, err
}
