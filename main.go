package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/config"
	"github.com/myselfBZ/BloggingAPI/pkg/handlers"
	"github.com/myselfBZ/BloggingAPI/pkg/models"
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
    h := handlers.NewHandler(&models.Blog{}, &models.Like{}, &models.User{}) 
	routes.BlogRoute(r, h)
	routes.UserRoutes(r, h)
	routes.LikeRoutes(r, h)
	log.Fatal(r.Run(":8080"))
}
