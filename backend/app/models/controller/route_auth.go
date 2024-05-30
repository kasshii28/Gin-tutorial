package controller

// import (
// 	"log"
// 	"net/http"
// 	"backend/app/models"
// 	"backend/utils"
// 	"github.com/gin-gonic/gin"
// )

// func signup(c *gin.Context){
// 	if c.Request.Method == "GET" {
// 		_, err := utils.GenerateTokenJWT()
// 		if err != nil {
// 			c.Redirect(http.StatusFound, "/login")
// 		} else {
// 			c.Redirect(http.StatusFound, "/todos")
// 		}
// 	} else if c.Request.Method == "POST" {
// 		err := c.ParseForm()
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		user := models.User{
// 			Name: c.Request.Form["name"],
// 			Email: c.Request.Form["email"],
// 			PassWord: c.Request.Form["password"],
// 		}
// 		if er := user.CreateUser(); err != nil {
// 			log.Fatalln("failed to create user",err)
// 		}

// 		c.Redirect(http.StatusFound, "/")
// 	}
// }