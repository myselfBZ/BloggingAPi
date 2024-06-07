package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/handlers"
	"github.com/myselfBZ/BloggingAPI/pkg/middleware"
)

func BlogRoute(r *gin.Engine) {
	r.GET("/blog", middleware.Authenticate(), handlers.GetBlogs)
	r.POST("/blog", middleware.Authenticate(), handlers.CreateBlog)
	r.DELETE("/blog/:id", middleware.Authenticate(), handlers.DeleteBlog)
	r.PUT("/blog/:id", middleware.Authenticate(), handlers.UpdateBlog)
	r.GET("/blog/:id", middleware.Authenticate(), handlers.GetBlog)
}
