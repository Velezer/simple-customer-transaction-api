package config

import (
	"sctrans/models"
	"sctrans/utils"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := utils.Getenv("DATABASE_URL", "root:@tcp(127.0.0.1:3306)/db_final_project?charset=utf8mb4&parseTime=True&loc=Local")
	if utils.Getenv("ENV", "") == "production" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&models.Customer{},
		&models.CustomerAddress{},
		&models.Product{},
		&models.PaymentMethod{},
		&models.Order{},
	)

	return db
}
