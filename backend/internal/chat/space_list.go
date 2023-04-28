package chat

import (
	"context"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/chat"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SpaceList(req *gin.Context) {
	ctx := context.Background()
	u, err := auth.CheckStatus(ctx, req, os.Getenv("SECRET_BASE"))

	if err != nil {
		req.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":      "error",
				"description": "unauthorized",
			},
		)

		return
	}

	apiKey, err := db.Client.GoogleApiKey.Get(ctx, u.ID)
	if err != nil {
		req.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":      "error",
				"description": "unauthorized",
			},
		)

		return
	}

	spaces, err := chat.GetSpaces(apiKey)
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to get Google Chat Space List",
			},
		)

		return
	}

	lst := make([]chat.Space, 0)
	for _, space := range *spaces {
		if space.SpaceType != "SPACE" {
			continue
		}

		lst = append(lst, space)
	}

	req.JSON(
		http.StatusOK,
		lst,
	)
}
