package main

import (
	"log"

	"github.com/myselfBZ/BloggingAPI/pkg/config"
	"github.com/myselfBZ/BloggingAPI/pkg/models"
)

func init() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(
			"You are actually bad at this",
		)
	}
	config.ConnectDB()

}

func main() {
	err := config.DB.AutoMigrate(&models.User{}, &models.Blog{}, &models.Like{})
	if err != nil {
		log.Fatal("You are bad at migrating the datbase", err)
	}
}
