package middlewares

import (
	"auth_ms/src/repository"
	"auth_ms/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthenticationMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	tokenSplit := strings.Split(tokenString, "bearer ")
	strings.TrimSpace(tokenSplit[1])

	valid, claims, err := service.VerifyJwtToken(tokenSplit[1])
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := findUser(claims.Issuer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}

func AuthenticationCookieMiddleware(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	valid, claims, err := service.VerifyJwtToken(cookie)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	user, err := findUser(claims.Issuer)
	if err != nil || user.Email == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}

func findUser(email string) (repository.User, error) {
	r, err := repository.NewRepository(nil)
	if err != nil {
		return repository.User{}, err
	}

	return r.FindUserByEmail(email)
}
