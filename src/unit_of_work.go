package src

import (
	"fmt"
	"go-e-s/src/backbone/helpers/errors_handler"
	"go-e-s/src/backbone/infrastructr/databases"
	"gorm.io/gorm"
)

type GormUnitOfWork struct {
	DB *gorm.DB
}

func NewGormUnitOfWork() GormUnitOfWork {
	db, err := databases.SetupDatabase()
	fmt.Println(err)
	return GormUnitOfWork{DB: db}
}

func (uow *GormUnitOfWork) Begin() (*gorm.DB, error) {
	tx := uow.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (uow *GormUnitOfWork) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (uow *GormUnitOfWork) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (uow *GormUnitOfWork) Transaction(f func(tx *gorm.DB) (any, error)) any {
	tx, err := uow.Begin()
	errors_handler.ErrorExist(err)
	value, errF := f(tx)
	if errF != nil {
		err = uow.Rollback(tx)
		errors_handler.ErrorExist(err)
		return errF
	}
	return value

}
