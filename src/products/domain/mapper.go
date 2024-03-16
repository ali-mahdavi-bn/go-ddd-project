package domain

import (
	"go-e-s/src/products/adapter/data_models"
	"go-e-s/src/products/domain/entities"
)

func MapNewUserModel(user *entities.User) *data_models.User {
	return &data_models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Password:  user.Password,
	}
}
