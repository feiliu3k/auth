package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Setup() {
	var err error
	// var user userModel.User
	DB, err = gorm.Open("mysql", "homestead:secret@tcp(127.0.0.1:33060)/golive?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
	AutoMigrateAll()
}

func AutoMigrateAll() {
	DB.AutoMigrate(&User{})
}