package cmd

import (
	"github.com/bear-san/googlechat-sender/backend/internal/auth"
	"github.com/bear-san/googlechat-sender/backend/internal/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setRoutes(engine *gin.Engine) {
	apiGroup := engine.Group("/api")

	apiGroup.GET("/health", func(req *gin.Context) {
		req.String(http.StatusOK, "pong")
	})

	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/login", auth.Login)
	authGroup.GET("/callback", auth.Callback)
	authGroup.GET("/verify", auth.Verify)

	spaceGroup := apiGroup.Group("/spaces")
	spaceGroup.GET("/", chat.SpaceList)
	spaceGroup.POST("/:sid/messages", chat.SpacePost)

	memberGroup := apiGroup.Group("/members")
	memberGroup.GET("/", chat.GWSMemberList)
	memberGroup.GET("/:uid/space", chat.DirectMessageFind)
	memberGroup.POST("/:uid/messages", chat.DirectMessagePost)

	scheduleGroup := apiGroup.Group("/schedules")
	scheduleGroup.GET("/", chat.ScheduleList)
	scheduleGroup.POST("/", chat.Schedule)
	scheduleGroup.PATCH("/:sid", chat.Reschedule)
	scheduleGroup.DELETE("/:sid", chat.UnSchedule)
}
