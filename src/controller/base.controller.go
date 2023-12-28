package controller

import (
	"auth_ms/src/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ApiController struct {
	repository *repository.Repository
}

func NewBaseController() *ApiController {
	r, err := repository.NewRepository(nil)
	if err != nil {
		log.Fatalln("Failed to connect to DB")
	}
	return &ApiController{
		repository: r,
	}
}

func (a ApiController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
