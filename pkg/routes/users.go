package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/handlers"
)

func UserRoutes(r *gin.Engine, h *handlers.Handler) {
	r.POST("/users", h.CreateUser)
	r.POST("/log-in", h.LogIn)
}
