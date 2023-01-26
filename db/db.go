package db

import (
	"log"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dbURL := "postgres://postgres:postgres@localhost:5432/postgres?dbname"

	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully connected to database...")
}
