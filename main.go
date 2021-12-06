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

	// 投稿一覧
	router.GET("/", handler.GetPosts)

	//登録
	router.POST("/new", handler.RegisterPost)

	//投稿詳細
	router.GET("/detail/:id", handler.GetDetailPost)

	//削除確認
	router.GET("/delete_check/:id", handler.CheckDeletion)

	//更新
	router.POST("/update/:id", handler.UpdatePost)

	//削除
	router.POST("/delete/:id", handler.RemovePost)

	router.Run()
}
