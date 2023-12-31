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
	r.POST("/api/login", c.Login)
	r.DELETE("/api/user/delete", c.SoftDeleteUser)
	r.PATCH("/user/password/update", middlewares.AuthenticationCookieMiddleware, c.UpdateUserPassword)
	r.POST("/user/password/reinitialise", c.ReinitializePassword)

	// Cookie methods
	r.GET("/api/auth/check-cookie", c.CheckAuthCookie)

	// Token methods
	r.GET("/api/auth/check-token", c.CheckAuthHeader)

	return r
}
