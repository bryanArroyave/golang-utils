package dtos

import (
	appadapter "github.com/bryanArroyave/golang-utils/app"
	"github.com/bryanArroyave/golang-utils/server/ports"
)

type APIRestServerConfigDTO struct {
	GlobalPrefix string
	Port         string
	App          *appadapter.App
}

type RouterHandler struct {
	Prefix string
	Router ports.IRouter
}
