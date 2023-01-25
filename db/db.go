package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_"gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

func init() {
	dbURL := "postgres://postgres:postgres@localhost:5432/postgres?dbname"
	
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}