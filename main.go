package main

import (
	"eliarms.events.com/db"
	"eliarms.events.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.NewSQLite3Repo("events.db")
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}
