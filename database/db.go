package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectToDatabase() {
  dsn := "root:root@tcp(127.0.0.1:3306)/expenses_manager?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error Establishing a Database Connection")
	}
}
