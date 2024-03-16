package entrypoints

import (
	"github.com/labstack/echo/v4"
	"go-e-s/src/backbone/service_layer"
	"go-e-s/src/products"
	"go-e-s/src/products/service_layer/handlers"
)

func InitProductsRoutes(e *service_layer.APIRouter) {
	e.Tag = "products"
	bootstrap := products.NewBootstrap()

	service_layer.Handlers(e, "GET", "/aa",
		func(c echo.Context) error {
			command := handlers.NewHelloCommand()
			return bootstrap.Handle(c, command)
		}, map[string]any{
			"summary": "create user",
		})

	service_layer.Handlers(e, "GET", "/bbaa",
		func(c echo.Context) error {
			command := handlers.NewHelloCommand()
			return bootstrap.Handle(c, command)
		}, map[string]any{
			"summary": "create user",
		})
}
