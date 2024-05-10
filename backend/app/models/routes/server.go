package routes
import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CorsSettings(){
	r := gin.Default()

	r.Use(cors.New(cors.Config{
	//許可するアクセス元
	AllowOrigins: []string{
		"http://localhost:3000",
	},
	//許可するHTTPメソッド
	AllowMethods: []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
	},
	//許可するHTTPリクエストヘッダ
	AllowHeaders: []string {
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	},
	//cookieなどの情報を必要とするか
	AllowCredentials: true,
	//preflightリクエストの結果をキャッシュする時間
	MaxAge: 24*time.Hour,
	}))
}

func StartMainServer() error {
	r := gin.Default()

	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome Todo",
		})
	})

	r.Run(":8000")
}