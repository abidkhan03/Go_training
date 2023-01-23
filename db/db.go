package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
func Init() *gorm.DB {
	dbURL := "postgres://postgres:admin123@localhost:5432/gorm"

	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DB.AutoMigrate(&Object{})

	return DB
}

