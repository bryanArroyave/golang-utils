package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	appadapter "github.com/bryanArroyave/golang-utils/app"
	loggerdtos "github.com/bryanArroyave/golang-utils/logger/dtos"
	serverdtos "github.com/bryanArroyave/golang-utils/server/dtos"
	"github.com/bryanArroyave/golang-utils/server/ports"

	"github.com/labstack/echo/v4"
)

// import (
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/labstack/echo/middleware"
// 	"github.com/labstack/echo/v4"
// )

// type Server struct {
// }

// // NewServer inicializa un servidor Echo con configuraciones comunes.
// func NewServer() *echo.Echo {
// 	e := echo.New()

// 	// Middleware común

// 	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			start := time.Now()
// 			_ = next(c)
// 			log.Info().
// 				Int("status", c.Response().Status).
// 				Dur("latency", time.Since(start)).
// 				Str("client_ip", c.RealIP()).
// 				Str("method", c.Request().Method).
// 				Str("path", c.Path()).
// 				Msg("")
// 			return nil
// 		}
// 	})

// 	e.Use(middleware.Recover())
// 	e.Use(middleware.CORS())

// 	return e
// }

// // StartServer inicia el servidor en un puerto dado.
// func StartServer(e *echo.Echo) {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	log.Fatal(e.Start(":" + port))
// }

type APIRestServer struct {
	echoInstance *echo.Echo
	globalGroup  *echo.Group
	port         string
	routes       func(*echo.Group)
	app          *appadapter.App
}

func NewAPIRestServer(config *serverdtos.APIRestServerConfigDTO) *APIRestServer {
	e := echo.New()

	if config.App == nil {
		panic("App instance is required")
	}

	return &APIRestServer{
		echoInstance: e,
		globalGroup:  e.Group(config.GlobalPrefix),
		port:         config.Port,
		app:          config.App,
		routes: func(group *echo.Group) {
			group.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					return c.JSON(http.StatusNotFound, map[string]interface{}{
						"message": fmt.Sprintf("%s not found", c.Request().RequestURI),
					})
				}
			})
		},
	}
}

func (api *APIRestServer) AddRoute(prefix string, router ports.IRouter) *APIRestServer {
	router.Handle(fmt.Sprintf("/%s", prefix), api.globalGroup)
	return api
}

func (api *APIRestServer) RunServer() {

	api.globalGroup.GET("/healthz", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	api.echoInstance.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			_ = next(c)
			if strings.Contains(c.Path(), "/healthz") || c.Path() == "/" || c.Path() == "" {
				return nil
			}

			api.app.GetLogger().Info("",
				&loggerdtos.LoggerFieldsDTO{Key: "status", Value: c.Response().Status},
				&loggerdtos.LoggerFieldsDTO{Key: "latency", Value: time.Since(start)},
				&loggerdtos.LoggerFieldsDTO{Key: "client_ip", Value: c.RealIP()},
				&loggerdtos.LoggerFieldsDTO{Key: "method", Value: c.Request().Method},
				&loggerdtos.LoggerFieldsDTO{Key: "path", Value: fmt.Sprintf("%s%s", c.Request().Host, c.Path())},
			)

			return nil
		}
	})

	api.routes(api.globalGroup)

	api.echoInstance.Logger.Fatal(api.echoInstance.Start(":" + api.port))
}

func (api *APIRestServer) GetEchoInstance() *echo.Echo {
	return api.echoInstance
}
