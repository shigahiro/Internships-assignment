package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shigahiro/gin-app/db"
	"github.com/shigahiro/gin-app/handler"
)

func main() {
	router := gin.Default()
	// HTMLファイルを読み込んだ結果をHTMLを描画するレンダリングエンジンというものに関連付ける
	router.LoadHTMLGlob("views/*.html")

	db.Init()
	// 登録
	router.POST("/new", handler.Register)
	// 登録内容詳細
	router.GET("/detail/:id")
	// 一覧
	router.GET("/")
	// 更新
	router.POST("/update/:id")
	// 削除
	router.POST("/delete/:id")
	// 削除内容
	router.GET("/delete_check/:id")

	router.Run()
}
