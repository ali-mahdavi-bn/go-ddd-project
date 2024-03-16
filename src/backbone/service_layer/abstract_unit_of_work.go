package service_layer

import "gorm.io/gorm"

type UnitOfWork interface {
	BeginTransaction() (*gorm.DB, error)
	CommitTransaction(tx *gorm.DB) error
	RollbackTransaction(tx *gorm.DB) error
}
