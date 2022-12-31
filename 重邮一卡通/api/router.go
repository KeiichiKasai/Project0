package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", login)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.POST("/recharge", recharge)
	}
	r.Run()
}
