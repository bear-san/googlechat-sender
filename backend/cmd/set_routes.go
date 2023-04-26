package cmd

import (
	"github.com/bear-san/googlechat-sender/backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func setRoutes(engine *gin.Engine) {
	apiGroup := engine.Group("/api")

	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/login", auth.Login)
}
