package adapter

import (
	"context"
	"gorm.io/gorm"
)

type AbstractRepository[T IBaseEntity] interface {
	ByID(ctx context.Context, id uint) (T, error)
	ByFiled(ctx context.Context, field string, id uint) (T, error)
	Save(ctx context.Context, model *T) error
}
type GormRepository[T IBaseEntity] struct {
	// implement: AbstractRepository
	db *gorm.DB
}

func NewGormRepository[T IBaseEntity](db *gorm.DB) AbstractRepository[T] {
	return &GormRepository[T]{
		db: db,
	}
}

func (g GormRepository[T]) ByID(ctx context.Context, id uint) (T, error) {
	return g.ByFiled(ctx, "id", id)
}

func (g GormRepository[T]) ByFiled(ctx context.Context, field string, id uint) (T, error) {
	var t T
	err := g.db.WithContext(ctx).Model(t).Where(field+"=?", id).First(&t).Error
	return t, err
}

func (g GormRepository[T]) Save(ctx context.Context, model *T) error {
	return g.db.WithContext(ctx).Save(model).Error
}

func (g GormRepository[T]) Delete(ctx context.Context, model T, softDelete bool) error {
	return g.db.WithContext(ctx).Delete(model).Error

}
