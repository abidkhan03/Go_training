package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_"gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

func init() {
	connStr := "host=localhost user=postgres password=admin123 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Karachi"
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}