package database

import (
	"api-apps/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func InitDatabase() {
	dsn := "root:ahay@tcp(localhost:3306)/api_go"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Filed to connect to database")
	}
	err := DB.AutoMigrate(&model.Products{})
	if err != nil {
		return
	}
}
