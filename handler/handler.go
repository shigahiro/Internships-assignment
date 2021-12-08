package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shigahiro/gin-app/db"
	"github.com/shigahiro/gin-app/model"
)

func getParamId(c *gin.Context) int {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return id
}

func GetPosts(c *gin.Context) {
	tweets := db.GetAll()
	// index.htmlを描画する。"tweets"をキーとしてバリューがtweetsのマップを作る
	c.HTML(200, "index.html", gin.H{"tweets": tweets})
}

func RegisterPost(c *gin.Context) {
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

func GetDetailPost(c *gin.Context) {
	id := getParamId(c)
	tweet := db.GetOne(id)
	c.HTML(200, "detail.html", gin.H{"tweet": tweet})
}

func CheckDeletion(c *gin.Context) {
	id := getParamId(c)
	tweet := db.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"tweet": tweet})
}

func UpdatePost(c *gin.Context) {
	id := getParamId(c)
	tweet := c.PostForm("tweet")
	db.Update(id, tweet)
	c.Redirect(302, "/")
}

func RemovePost(c *gin.Context) {
	id := getParamId(c)
	db.Delete(id)
	c.Redirect(302, "/")
}
