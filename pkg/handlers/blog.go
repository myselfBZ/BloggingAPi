package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/models"
	"github.com/myselfBZ/BloggingAPI/pkg/utils"
	"gorm.io/gorm"
)

func (h *Handler) CreateBlog(c *gin.Context) {
	userID := c.MustGet("id").(uint)
	var blog models.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	blog.Likes = make([]models.Like, 0) 
	blog.UserID = userID
	err := h.BlogStorage.Create(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, blog)
}

func (h *Handler)DeleteBlog(c *gin.Context) {
    id := c.Param("id")
	userId := c.MustGet("id").(uint)
	validID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.BlogStorage.Delete(uint(validID), userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server sucks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})

}

func (h *Handler) GetBlogs(c *gin.Context) {

	var blogs, err = h.BlogStorage.GetBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server is not okay"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

func (h *Handler)UpdateBlog(c *gin.Context) {
	var blog *models.Blog
	userId := c.MustGet("id").(uint)
    id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
    UintId := uint(id)
	if err := c.BindJSON(blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	blog, err = h.BlogStorage.Update(UintId, blog, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server is ill"})
		return
	}
	updatedBlog := map[string]any{
		"Content": blog.Content,
		"Title":   blog.Title,
	}
	c.JSON(http.StatusOK, updatedBlog)
}

func (h *Handler)GetBlog(c *gin.Context) {
	id := c.Param("id")
	validatedId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	fetchedBlog, err := h.BlogStorage.GetBlog(uint(validatedId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, utils.InternaleServer)
		return
	}
	c.JSON(http.StatusOK, fetchedBlog)

}
