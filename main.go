package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/db"
	routes "main.go/route"
)

func main() {

	//Initializong Database with Init Method defined in db.go
	db.InitDB()
	//Creates a Server
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	routes.RegisterRoute(server)
	//Actives the Server at port :8080
	server.Run(":8080") //localhost:8080

}
