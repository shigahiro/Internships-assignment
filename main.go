package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shigahiro/gin-app/model"
)

func main() {
	router := gin.Default()
	dbInit()
	router.GET("/")
	router.Run()
}

func dbInit() {
	db := gormConnect()

	// コネクション解放解放
	defer db.Close()
	db.AutoMigrate(&model.Tweet{}) //構造体に基づいてテーブルを作成
}

func gormConnect() *gorm.DB {
	// MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	DBMS := "mysql"
	USER := "test"
	PASS := "12345678"
	DBNAME := "test"
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
