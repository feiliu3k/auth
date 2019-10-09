package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Gender string `json:"gender"`
}

func AddUser(user *User) {
	DB.Create(&user)
	return
}

func UserDetailByName(name string) (user User) {
	DB.Where("name = ?", name).First(&user)
	return
}

func UserDetailByEmail(email string) (user User) {
	DB.Where("email = ?", email).First(&user)
	return
}

func UserDetail(id uint) (user User) {
	DB.Where("id = ?", id).First(&user)
	return
}