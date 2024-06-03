package controller

import (
	"log"
	"net/http"
	"backend/app/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context){
	if c.Request.Method == http.MethodGet {
		_, err := session(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"redirect": "/signup"})
		} else {
			c.JSON(http.StatusFound, gin.H{"redirect": "/todos"})
		}
	} else if c.Request.Method == http.MethodPost {
		err := c.Request.ParseForm()
		if err != nil {
			log.Fatalln("failed to parse",err)
		}
		user := models.User{
			Name: c.PostForm("name"),
			Email: c.PostForm("email"),
			PassWord: c.PostForm("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Fatalln("failed to create user",err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return 
		}

		c.JSON(http.StatusFound, gin.H{"redirect": "/"})
	}
}