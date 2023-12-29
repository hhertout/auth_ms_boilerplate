package controller

import (
	"auth_ms/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (b BaseController) Login(c *gin.Context) {
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

	user, err := b.repository.FindUserByEmail(body.Email)
	if err != nil || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Check your information",
		})
		return
	}

	passwordIsCorrect, err := service.VerifyPassword(body.Password, user.Password)
	if err != nil || !passwordIsCorrect {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Check your information",
		})
		return
	}

	// logic return goes here
	token, err := service.GenerateJwtToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "An error occurred, please try again later. If the problem persist, contact your platform administrator",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*20, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged",
		"user":    user.Email,
	})
}

func (b BaseController) CheckToken(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, err := service.VerifyJwtToken(cookie)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authorized",
	})
}
