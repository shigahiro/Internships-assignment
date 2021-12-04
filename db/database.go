package db

import (
	"github.com/jinzhu/gorm"
	model "github.com/shigahiro/gin-app/models"
)

func Init() {
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
