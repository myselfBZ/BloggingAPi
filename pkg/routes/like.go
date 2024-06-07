package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/handlers"
	"github.com/myselfBZ/BloggingAPI/pkg/middleware"
)

func LikeRoutes(r *gin.Engine) {
	r.POST("/like/:id", middleware.Authenticate(), handlers.Like)
}
