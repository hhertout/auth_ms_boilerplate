package controller

import (
	"auth_ms/src/repository"
	"auth_ms/src/service"
	"fmt"
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

func (b BaseController) UpdateUserPassword(c *gin.Context) {
	type Body struct {
		Password string `json:"password"`
	}
	var body Body

	err := c.BindJSON(&body)
	if err != nil || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	u, ok := user.(repository.User)
	fmt.Println(u)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	encryptPassword, err := service.HashPassword(body.Password)
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = b.repository.UpdatePassword(u.Id, encryptPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password update successfully",
	})
}

func (b BaseController) SoftDeleteUser(c *gin.Context) {
	type Body struct {
		Email string `json:"email"`
	}
	var body Body

	if err := c.BindJSON(&body); err != nil || body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	user, err := b.repository.FindUserByEmail(body.Email)
	if err != nil || user.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User doesn't exist",
		})
		return
	}

	_, err = b.repository.SoftDelete(body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully deleted",
	})
}
