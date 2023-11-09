package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouters(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.Use(JwtMiddleware())
	{
	}
}
