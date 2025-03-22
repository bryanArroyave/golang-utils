package ports

import "github.com/labstack/echo/v4"

type IRouter interface {
	Handle(groupPath string, group *echo.Group)
}
