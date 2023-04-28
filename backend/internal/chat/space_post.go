package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/chat"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func SpacePost(req *gin.Context) {
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

	space, err := chat.GetOneSpace(apiKey, fmt.Sprintf("spaces/%s", req.Param("sid")))
	if err != nil {
		req.JSON(
			http.StatusNotFound,
			gin.H{
				"status":      "error",
				"description": "not found",
			},
		)

		return
	}

	bodyTxt, err := io.ReadAll(req.Request.Body)

	var p chat.Message
	err = json.Unmarshal(bodyTxt, &p)
	if err != nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "invalid request",
			},
		)

		return
	}

	res, err := space.Post(apiKey, p)
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to post message",
			},
		)

		return
	}

	req.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": res,
		},
	)
}
