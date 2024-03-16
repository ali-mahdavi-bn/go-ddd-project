package service_layer

import "github.com/labstack/echo/v4"

type APIRouter struct {
	Router             *echo.Echo
	Tag                string
	Path               string
	InitializerSwagger bool
}

func (u *APIRouter) Group(prefix string) *APIRouter {
	u.Path += prefix
	return u
}
