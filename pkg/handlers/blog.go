package handlers

import (
	"log"
	"net/http"
	"strconv"
    "time"
	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/models"
	"github.com/myselfBZ/BloggingAPI/pkg/utils"
	"gorm.io/gorm"
)

func CreateBlog(c *gin.Context) {
	userID := c.MustGet("id").(uint)
	var blog models.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var likes []models.Like
	blog.Likes = likes
	blog.UserID = userID
	err := blog.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, blog)
}

func DeleteBlog(c *gin.Context) {
    start := time.Now()	
    id := c.Param("id")
	userId := c.MustGet("id").(uint)
	validID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
    var tempBlog *models.Blog
	if err := tempBlog.Delete(uint(validID), userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server sucks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
    log.Println("this operation took: ", time.Since(start)) 

}

func GetBlogs(c *gin.Context) {
	var blog *models.Blog

	var blogs, err = blog.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server is not okay"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

func UpdateBlog(c *gin.Context) {
	var blog *models.Blog
	userId := c.MustGet("id").(uint)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	blog, err = blog.Update(uint(id), blog, userId)
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

func GetBlog(c *gin.Context) {
	id := c.Param("id")
	validatedId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var blog models.Blog
	fetchedBlog, err := blog.GetBlog(uint(validatedId))
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
