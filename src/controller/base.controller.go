package controller

import (
	"auth_ms/src/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type BaseController struct {
	repository *repository.Repository
}

func NewBaseController() *BaseController {
	r, err := repository.NewRepository(nil)
	if err != nil {
		log.Fatalln("Failed to connect to DB")
	}
	return &BaseController{
		repository: r,
	}
}

func (b BaseController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
