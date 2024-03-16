package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-e-s/src/backbone/api"
	"go-e-s/src/backbone/api/middleware"
	_ "go-e-s/src/backbone/docs" // Import the docs
	"go-e-s/src/backbone/service_layer"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.LoggerMiddleware)

	// rotes
	apiRouter := service_layer.APIRouter{Router: e}
	api.InitRoute(apiRouter.Group("/api"))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// started app
	e.Logger.Fatal(e.Start(":8080"))

}
