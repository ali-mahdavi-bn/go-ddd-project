package middleware

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"go-e-s/src/backbone/service_layer"
	"time"
)

// LoggerMiddleware logs the details of each request.
func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		elapsed := time.Since(start)

		req := c.Request()
		res := c.Response()
		methodColor := color.New(color.FgGreen).SprintFunc()
		pathColor := color.New(color.FgCyan).SprintFunc()
		statusColor := color.New(color.FgYellow).SprintFunc()
		elapsedColor := color.New(color.FgMagenta).SprintFunc()
		errorColor := color.New(color.FgRed).SprintFunc()

		fmt.Printf("Method: %s Path: %s Status: %s Time: %s error: %s\n",
			methodColor(req.Method),
			pathColor(req.URL.Path),
			statusColor(res.Status),
			elapsedColor(elapsed),
			errorColor(err))

		return err
	}
}

func NewAPIRouter(InitializerSwagger bool) *service_layer.APIRouter {
	return &service_layer.APIRouter{InitializerSwagger: InitializerSwagger}
}
