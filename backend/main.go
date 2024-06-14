package main

import (
	"backend/app/models/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.CorsSettings(r)
	controller.StartMainServer(r)
}