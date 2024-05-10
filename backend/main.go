package main
import (
	"github.com/Gin-tutorial/backend/app/models/routes"
)

func main() {
	routes.CorsSettings()
	routes.StartMainServer()
}