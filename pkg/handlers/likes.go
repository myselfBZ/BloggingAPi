package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/models"
	"gorm.io/gorm"
)

func Like(c *gin.Context) {
	userId := c.MustGet("id").(uint)
	blogId := c.Param("id")
	validatedId, err := strconv.Atoi(blogId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	like := &models.Like{
		UserID: userId,
		BlogID: uint(validatedId),
	}
	err = like.Like()
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
