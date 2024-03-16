package databases

import (
	"fmt"
	products "go-e-s/src/products/adapter/data_models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := "user=postgres password=ali3z110 host=localhost dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}
func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&products.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate schema: %v", err)
	}
	return nil
}
