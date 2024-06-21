package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func(h *Handler) Like(c *gin.Context) {
	userId := c.MustGet("id").(uint)
	blogId := c.Param("id")
	validatedId, err := strconv.Atoi(blogId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.LikeStorage.Like(userId, uint(validatedId))
	if err != nil {
		if err == gorm.ErrForeignKeyViolated {
			c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"liked": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"liked": true})
}
