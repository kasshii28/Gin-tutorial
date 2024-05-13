package controller

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CorsSettings() {
	r := gin.Default()

    // Corsの設定
    r.Use(cors.New(cors.Config{
        // アクセス許可するオリジン
        AllowOrigins: []string{
            "http://localhost:3000",
        },
        // アクセス許可するHTTPメソッド
        AllowMethods: []string{
            "GET",
            "POST",
			"PUT",
			"DELETE",
            "OPTIONS",
        },
        // 許可するHTTPリクエストヘッダ
        AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
        },
        // cookieなどの情報を必要とするかどうか
        AllowCredentials: true,
        // preflightリクエストの結果をキャッシュする時間
        MaxAge: 24 * time.Hour,
    }))
}

func StartMainServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome todo",
		})
	})

	r.Run(":8000")
}