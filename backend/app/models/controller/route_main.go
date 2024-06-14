package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func top(c *gin.Context) {
	_, err := session(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome Todo List!",
		})
	} else {
		c.JSON(http.StatusFound, gin.H{"redirect-url": "/todos"})
	}
}