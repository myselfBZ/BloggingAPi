package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/handlers"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/users", handlers.CreateUser)
	r.POST("/log-in", handlers.LogIn)
}
