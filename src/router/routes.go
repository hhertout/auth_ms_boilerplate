package router

import (
	"auth_ms/src/controller"
	middlewares "auth_ms/src/middleware"
	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	r := gin.Default()
	c := controller.NewBaseController()

	r.Use(middlewares.CORSMiddleware())

	r.GET("/ping", c.Ping)

	r.POST("/api/user/new", c.CreateUser)
	r.DELETE("/api/user/delete", c.SoftDeleteUser)
	r.POST("/api/login", c.Login)

	r.GET("/api/auth/check-token", c.CheckToken)

	return r
}
