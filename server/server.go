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

type APIRestServer struct {
	echoInstance *echo.Echo
	publicGroup  *echo.Group
	privateGroup *echo.Group
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
		publicGroup:  e.Group(config.GlobalPrefix),
		privateGroup: e.Group(config.GlobalPrefix),
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
	router.Handle(fmt.Sprintf("/%s", prefix), api.publicGroup)
	return api
}

func (api *APIRestServer) RunServer() {

	api.publicGroup.GET("/healthz", func(c echo.Context) error {
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

	api.routes(api.publicGroup)

	api.echoInstance.Logger.Fatal(api.echoInstance.Start(":" + api.port))
}

func (api *APIRestServer) GetEchoInstance() *echo.Echo {
	return api.echoInstance
}

func (api *APIRestServer) GetPublicGroup() *echo.Group {
	return api.publicGroup
}

func (api *APIRestServer) GetPrivateGroup() *echo.Group {
	return api.privateGroup
}
