package handlers

import (
	"context"
	"go-e-s/src"
	"go-e-s/src/backbone/helpers/errors_handler"
	"go-e-s/src/backbone/service_layer"
	"go-e-s/src/products/adapter/repositories"
	"go-e-s/src/products/domain/entities"
	"gorm.io/gorm"
	"net/http"
)

type HelloCommand struct {
	// implement: Command
	service_layer.CommandHandles
}

func NewHelloCommand() *HelloCommand {
	dependencies := []string{
		"uow",
		"request",
	}
	helloCmd := &HelloCommand{}
	helloCmd.SetDependencies(dependencies)
	return helloCmd
}

func (c *HelloCommand) CommandHandler(dependencies map[string]any) error {
	uow := dependencies["uow"].(src.GormUnitOfWork)
	user4 := uow.Transaction(func(tx *gorm.DB) (any, error) {
		b := repositories.NewGormUserRepo(tx)
		id, err := b.ByID(context.Background(), 4)
		data, err := errors_handler.ErrorExistOrSuccess(id, err)
		return data, err
	})
	user3 := uow.Transaction(func(tx *gorm.DB) (any, error) {
		b := repositories.NewGormUserRepo(tx)
		id, err := b.ByID(context.Background(), 3)
		data, err := errors_handler.ErrorExistOrSuccess(id, err)
		return data, err
	})
	a := []entities.User{
		user3.(entities.User),
		user4.(entities.User),
	}
	return c.Request.JSON(http.StatusOK, a)
}
