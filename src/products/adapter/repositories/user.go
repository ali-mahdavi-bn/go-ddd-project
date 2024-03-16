package repositories

import (
	"go-e-s/src/backbone/adapter"
	"go-e-s/src/products/domain/entities"
	"gorm.io/gorm"
)

type UserRepo interface {
	adapter.AbstractRepository[entities.User]
}

type GormUserRepo struct {
	adapter.AbstractRepository[entities.User]
}

func NewGormUserRepo(db *gorm.DB) UserRepo {
	return &GormUserRepo{AbstractRepository: adapter.NewGormRepository[entities.User](db)}
}
