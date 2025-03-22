package main

import (
	"os"

	appadapter "github.com/bryanArroyave/golang-utils/app"
	appdtos "github.com/bryanArroyave/golang-utils/app/dtos"
	mongodtos "github.com/bryanArroyave/golang-utils/mongo/dtos"
	"github.com/bryanArroyave/golang-utils/server"
	serverdtos "github.com/bryanArroyave/golang-utils/server/dtos"
	"github.com/bryanArroyave/golang-utils/server/ports"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: main para probar la conexión a MongoDB
/*func main() {

	os.Setenv("MONGO_HOST", "localhost")
	os.Setenv("MONGO_PORT", "27017")
	os.Setenv("MONGO_INITDB_DATABASE", "eventSplit")
	os.Setenv("MONGO_INITDB_ROOT_USERNAME", "root")
	os.Setenv("MONGO_INITDB_ROOT_PASSWORD", "example")

	ctx := context.Background()
	config := &dtos.MongoConnectionDTO{
		User:     os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		DBName:   os.Getenv("MONGO_INITDB_DATABASE"),
		Port:     os.Getenv("MONGO_PORT"),
	}

	err := mongo.InitMongoDB(ctx, config)
	if err != nil {
		log.Fatal("Error al conectar a MongoDB:", err)
	}

	// Obtener la base de datos en cualquier parte del código
	db := mongo.GetMongoDB()
	fmt.Println("Conexión exitosa a MongoDB:", db.Name())
}
*/

// maqin para iniciar servidor
func main() {

	os.Setenv("MONGO_HOST", "localhost")
	os.Setenv("MONGO_PORT", "27017")
	os.Setenv("MONGO_INITDB_DATABASE", "journey")
	os.Setenv("MONGO_INITDB_ROOT_USERNAME", "root")
	os.Setenv("MONGO_INITDB_ROOT_PASSWORD", "example")

	// Inicializar configuración del servidor
	config := &appdtos.LoggerConfigDTO{
		LoggerType:  "logrus",
		ServiceName: "s1",
	}

	// Inicializar app
	app := appadapter.
		NewApp(config).
		AddMongoConnection("clients", &mongodtos.MongoConnectionDTO{
			Host:     os.Getenv("MONGO_HOST"),
			Port:     os.Getenv("MONGO_PORT"),
			DBName:   os.Getenv("MONGO_INITDB_DATABASE"),
			User:     os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
			Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		})

	server1 := server.NewAPIRestServer(&serverdtos.APIRestServerConfigDTO{
		GlobalPrefix: "/s1/v1",
		Port:         "8080",
		App:          app,
	})

	server1.AddRoute("prueba", NewPrueba1Routes(app.GetMongoConnection("clients")))

	go server1.RunServer()

	for {
	}

}

type Prueba1Routes struct {
	mongoDB *mongo.Database
}

func NewPrueba1Routes(mongoDB *mongo.Database) ports.IRouter {
	return &Prueba1Routes{
		mongoDB: mongoDB,
	}
}

func (route *Prueba1Routes) Handle(groupPath string, group *echo.Group) {
	roleGroup := group.Group(groupPath)

	roleGroup.GET("", func(c echo.Context) error {
		route.mongoDB.Collection("prueba").InsertOne(c.Request().Context(), map[string]interface{}{
			"message": "Hello World from s1",
		})
		return c.JSON(200, "Hello World from s1")
	})

}
