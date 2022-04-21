package config

import (
	"sctrans/models"
	"sctrans/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := utils.Getenv("DATABASE_URL", "postgres://postgres:password@localhost:5432/sctrans")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(4)

	db.AutoMigrate(
		&models.Customer{},
		&models.CustomerAddress{},
		&models.Product{},
		&models.PaymentMethod{},
		&models.Order{},
	)

	return db
}
