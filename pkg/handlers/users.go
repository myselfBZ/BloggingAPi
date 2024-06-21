package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myselfBZ/BloggingAPI/pkg/middleware"
	"github.com/myselfBZ/BloggingAPI/pkg/models"
	"github.com/myselfBZ/BloggingAPI/pkg/utils"
	"gorm.io/gorm"
)

func (h *Handler)CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidJSON)
		return
	}
	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.InternaleServer)
		return
	}
    err = h.UserStorage.Create(&user)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "username is already taken"})
		return

	}
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.InternaleServer)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (h *Handler) LogIn(c *gin.Context) {
	var creadentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&creadentials); err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidJSON)
		return
	}
	var fetchedUser *models.User
	var err error
	fetchedUser, err = h.UserStorage.Get(creadentials.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, utils.InvalidCreadentials)
			return
		}
		c.JSON(http.StatusInternalServerError, utils.InternaleServer)
		return
	}
	isValid := utils.CompareHash(creadentials.Password, fetchedUser.Password)
	if !isValid {
		c.JSON(http.StatusUnauthorized, utils.InvalidCreadentials)
		return
	}
	token, err := middleware.GenerateToken(fetchedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.InternaleServer)
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
