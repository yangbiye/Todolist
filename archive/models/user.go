package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password string `gorm:"type:varchar(256);not null"`
	Email    string `gorm:"type:varchar(50);not null"`
}

//用户登录之后给出用户信息
func LoginUser(userID string) (User, *gorm.DB) {
	user := User{}
	result := db.First(&user, userID)
	return user, result
}

//用户注册
func RegisterUser(password string, email string) (User, *gorm.DB) {
	user := User{Password: password, Email: email}
	result := db.Create(&user)
	return user, result
}
