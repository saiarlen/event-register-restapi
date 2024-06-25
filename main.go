package main

import (
	"eventapi/database"
	"eventapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":9000") // localhsot:9000

}
