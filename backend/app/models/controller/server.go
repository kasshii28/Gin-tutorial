package controller

import (
    "fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
    "regexp"
    "strconv"
    "backend/app/models"
)


func session(c *gin.Context) (sess models.Session, err error) {
    cookie, err := c.Cookie("_cookie")
    if err == nil {
        sess = models.Session{UUID: cookie} // contextで取得した場合は.Valueは必要ない
        if ok, _ := sess.CheckSession(); !ok {
            err = fmt.Errorf("invalid session")
        }
    }
    return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(*gin.Context, int)) gin.HandlerFunc{
    return func(c *gin.Context){
        path := c.Request.URL.Path
        q := validPath.FindStringSubmatch(path)
        if q == nil {
            c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
            return
        }

        qi, err := strconv.Atoi(q[2])
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
            return
        }
        fn(c, qi)
    }
}

func CorsSettings(r *gin.Engine) {
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
            "Access-Control-Allow-Origin",
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

func StartMainServer(r *gin.Engine) {
    r.GET("/", top)
    r.POST("/signup", signup)
    r.POST("/signin", signin)

	r.Run(":8000")
}