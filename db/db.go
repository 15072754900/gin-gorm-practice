package db

import (
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/mysql"

var DB *gorm.DB

func Connect2DB() {
	var err error
	dsn := "root:123456@tcp(localhost:3306)/tiesiyuan?charset=utf8mb4&parseTime=True&loc=Local"
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("error on link database!")
	}
}
