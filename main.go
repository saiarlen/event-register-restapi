package main

import (
	"eventapi/database"
	"eventapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()       //init db
	server := gin.Default() //init server

	routes.RegisterRoutes(server) //registering routes

	server.Run(":9000") // localhost:9000

}
