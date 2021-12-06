package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shigahiro/gin-app/db"
	"github.com/shigahiro/gin-app/model"
)

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
