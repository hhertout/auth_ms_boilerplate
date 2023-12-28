package controller

import (
	"auth_ms/src/service"
	"fmt"
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Check your information",
		})
		return
	}
	fmt.Println(passwordIsCorrect)

	// logic return goes here
}

func (b BaseController) CheckToken(c *gin.Context) {

}
