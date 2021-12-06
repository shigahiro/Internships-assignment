package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shigahiro/gin-app/db"
	"github.com/shigahiro/gin-app/model"
)

func main() {
	router := gin.Default()
	// HTMLファイルを読み込んだ結果をHTMLを描画するレンダリングエンジンというものに関連付ける
	// router.LoadHTMLGlob("views/*.html")

	db.Init()
	// 登録
	router.POST("/new", Register(c))
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

func getParamId(c *gin.Context) model.Tweet {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return db.GetOne(id)
}

func Register(c *gin.Context) {
	var form model.Tweet
	// 構造体のタグでバリデーションができる
	if err := c.Bind(&form); err != nil {
		tweets := db.GetAll()
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"tweets": tweets, "err": err})
		// 他のハンドラーの呼び出しを防ぐ
		c.Abort()
	} else {
		content := c.PostForm("content")
		db.Insert(content)
		c.Redirect(302, "/")
	}
}

func GetDetail(c *gin.Context) {
	tweet := getParamId(c)
	c.HTML(200, "detail.html", gin.H{"tweet": tweet})
}
