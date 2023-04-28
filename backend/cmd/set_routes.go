package cmd

import (
	"github.com/bear-san/googlechat-sender/backend/internal/auth"
	"github.com/bear-san/googlechat-sender/backend/internal/chat"
	"github.com/gin-gonic/gin"
)

func setRoutes(engine *gin.Engine) {
	apiGroup := engine.Group("/api")

	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/login", auth.Login)
	authGroup.GET("/callback", auth.Callback)

	spaceGroup := apiGroup.Group("/spaces")
	spaceGroup.GET("/", chat.SpaceList)
	spaceGroup.POST("/:sid", chat.SpacePost)
}
