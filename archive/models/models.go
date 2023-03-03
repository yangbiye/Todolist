package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Setup 启动数据库
func Setup() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("models/modles.go Setup() error=", err.Error())
	}

	db.AutoMigrate(&User{}, &Item{})
}
