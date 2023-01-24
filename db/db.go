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
	//connStr := "host=localhost user=postgres password=admin123 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Karachi"
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}