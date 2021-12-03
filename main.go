package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Tweet struct {
	gorm.Model
	Content string `form:"content" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("detail/:id")
	router.Run()
}
