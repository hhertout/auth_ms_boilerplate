package controller

import (
	"auth_ms/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (b BaseController) CreateUser(c *gin.Context) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body Body

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		return
	}

	if err := service.ValidateUserCreationData(body.Email, body.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		return
	}

	user, err := b.repository.FindUserByEmail(body.Email)
	if err != nil || user.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exist",
		})
		return
	}

	encryptPassword, err := service.HashPassword(body.Password)
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = b.repository.CreateUser(body.Email, encryptPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": body.Email,
	})
}
