package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// サーバーの初期設定
	r := gin.Default()

	// ルート（/）にアクセスが来たらJSONを返す設定
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Game Server is running!",
		})
	})

	// 8080ポートで起動
	r.Run(":8080")
}
