package api

import (
	"go-e-s/src/backbone/service_layer"
	products "go-e-s/src/products/entrypoints"
)

func InitRoute(e *service_layer.APIRouter) {
	// Product
	products.InitProductsRoutes(e)
}
