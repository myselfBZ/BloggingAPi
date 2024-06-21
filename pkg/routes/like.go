package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/handlers"
	"github.com/myselfBZ/BloggingAPI/pkg/middleware"
)

func LikeRoutes(r *gin.Engine, h *handlers.Handler) {
	r.POST("/like/:id", middleware.Authenticate(), h.Like)
}
