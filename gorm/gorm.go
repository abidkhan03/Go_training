package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB // database

func init() {
	var err error
	db, err = gorm.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Object{})
}

// func init() {
// 	var err error
// 	db, err = gorm.Open("sqlite3", "./gorm.db")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	db.AutoMigrate(&Object{})
// }

// Object model
type Object struct {
	gorm.Model
	Name string
}

// GetObjects get all objects
func GetObjects(c *gin.Context) {
	var objects []Object
	if err := db.Find(&objects).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, objects)
	}
}

// GetObject get object by id
func GetObject(c *gin.Context) {
	id := c.Params.ByName("id")
	var object Object
	if err := db.Where("id = ?", id).First(&object).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, object)
	}
}

// CreateObject create object
func CreateObject(c *gin.Context) {
	var object Object
	c.BindJSON(&object)
	db.Create(&object)
	c.JSON(http.StatusOK, object)
}

// UpdateObject update object
func UpdateObject(c *gin.Context) {
	var object Object
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&object).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}
	c.BindJSON(&object)
	db.Save(&object)
	c.JSON(http.StatusOK, object)
}

// DeleteObject delete object
func DeleteObject(c *gin.Context) {
	var object Object
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&object)
	fmt.Println(d)
	c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}

func main() {
	r := gin.Default()
	r.GET("/object", GetObjects)
	r.GET("/object/:id", GetObject)
	r.POST("/object", CreateObject)
	r.PUT("/object/:id", UpdateObject)
	r.DELETE("/object/:id", DeleteObject)
	r.Run(":8080")
}
