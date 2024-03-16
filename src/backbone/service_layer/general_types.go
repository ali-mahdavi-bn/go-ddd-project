package service_layer

import (
	"github.com/labstack/echo/v4"
)

type IEventDispatcher interface {
	Dispatch(events []IBaseEvent)
}
type HandlerFunc echo.HandlerFunc
type IBaseEvent interface {
}

type Command interface {
	GetDependencies() []string
	SetRequest(request echo.Context)
	CommandHandler(dependencies map[string]any) error
}

type CommandHandles struct {
	Request      echo.Context
	Dependencies []string
}

func (c *CommandHandles) GetDependencies() []string {
	return c.Dependencies
}
func (c *CommandHandles) SetDependencies(dependencies []string) {
	c.Dependencies = dependencies
}
func (c *CommandHandles) SetRequest(request echo.Context) {
	c.Request = request
}
