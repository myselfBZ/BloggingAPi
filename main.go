package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/config"
	"github.com/myselfBZ/BloggingAPI/pkg/routes"
)

func init() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
		return
	}
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.BlogRoute(r)
	routes.UserRoutes(r)
	routes.LikeRoutes(r)
	log.Fatal(r.Run(":8080"))
}
