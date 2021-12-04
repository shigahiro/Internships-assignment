package model

import "github.com/jinzhu/gorm"

type Tweet struct {
	gorm.Model
	Content string `form:"content" binding:"required"`
}

type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}
