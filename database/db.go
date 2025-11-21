package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	connectionString := "host=localhost user=root dbname=go_rest_api password=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic("failed to connect database")
	} else {
		println("Database connected")
	}
}
