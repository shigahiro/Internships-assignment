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
	router.LoadHTMLGlob("views/*.html")

	db.Init()
	router.GET("/", func(c *gin.Context) {
		tweets := db.GetAll()
		c.HTML(200, "index.html", gin.H{"tweets": tweets})
	})

	//登録
	router.POST("/new", func(c *gin.Context) {
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
	})

	//投稿詳細
	router.GET("/detail/:id", func(c *gin.Context) {
		id := getParamId(c)
		tweet := db.GetOne(id)
		c.HTML(200, "detail.html", gin.H{"tweet": tweet})
	})

	//削除確認
	router.GET("/delete_check/:id", func(c *gin.Context) {
		id := getParamId(c)
		tweet := db.GetOne(id)
		c.HTML(200, "delete.html", gin.H{"tweet": tweet})
	})

	//更新
	router.POST("/update/:id", func(c *gin.Context) {
		id := getParamId(c)
		tweet := c.PostForm("tweet")
		db.Update(id, tweet)
		c.Redirect(302, "/")
	})

	//削除
	router.POST("/delete/:id", func(c *gin.Context) {
		id := getParamId(c)
		db.Delete(id)
		c.Redirect(302, "/")

	})

	router.Run()
}

func getParamId(c *gin.Context) int {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return id
}
