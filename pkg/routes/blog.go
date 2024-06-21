package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/handlers"
	"github.com/myselfBZ/BloggingAPI/pkg/middleware"
)

func BlogRoute(r *gin.Engine, h *handlers.Handler) {
	r.GET("/blog", middleware.Authenticate(), h.GetBlogs)
	r.POST("/blog", middleware.Authenticate(), h.CreateBlog)
	r.DELETE("/blog/:id", middleware.Authenticate(), h.DeleteBlog)
	r.PUT("/blog/:id", middleware.Authenticate(), h.UpdateBlog)
	r.GET("/blog/:id", middleware.Authenticate(), h.GetBlog)
}
