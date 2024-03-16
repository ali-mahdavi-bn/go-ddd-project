package products

import (
	"github.com/labstack/echo/v4"
	"go-e-s/src"
	"go-e-s/src/backbone/service_layer"
)

var Dependencies = map[string]any{
	"uow": src.NewGormUnitOfWork(),
}
var name = src.NewGormUnitOfWork()

type Bootstrap struct {
	commands map[string]service_layer.Command
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		commands: make(map[string]service_layer.Command),
	}
}

func (h *Bootstrap) Handle(c echo.Context, command service_layer.Command) error {
	command.SetRequest(c)
	commandDependencies := command.GetDependencies()
	injectDependency := make(map[string]any)
	for _, i := range commandDependencies {
		if v, ok := Dependencies[i]; ok {
			injectDependency[i] = v
		}
	}
	resultCommandHandler := command.CommandHandler(injectDependency)
	return resultCommandHandler
}
