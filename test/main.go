package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"im/models"
)

func main() {
	dsn := "root:pwassword@tcp(127.0.0.1:3306/hi_chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		panic(err)
	}
}
