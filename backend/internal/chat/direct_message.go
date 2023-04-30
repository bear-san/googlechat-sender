package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/ent"
	auth2 "github.com/bear-san/googlechat-sender/backend/internal/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/chat"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func DirectMessagePost(req *gin.Context) {
	ctx := context.Background()
	u, err := auth.CheckStatus(ctx, req, os.Getenv("SECRET_BASE"))
	var apiKey *ent.GoogleApiKey
	if err == nil {
		apiKey, err = auth.GetGoogleCredential(ctx, u, auth2.GetOAuthClientInfo())
	}

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

	space, err := chat.FindDirectMessage(apiKey, req.Param("uid"))
	if err != nil || space.Name == nil {
		space, err = chat.CreateDirectMessage(apiKey, req.Param("uid"))
	}

	if err != nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "bad request",
			},
		)

		return
	}

	reqTxt, err := io.ReadAll(req.Request.Body)
	if err != nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "bad request",
			},
		)

		return
	}

	var msgReq chat.Message
	err = json.Unmarshal(reqTxt, &msgReq)
	if err != nil || msgReq.Text == nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "bad request",
			},
		)
	}

	msg, err := space.Post(apiKey, msgReq)

	if err != nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": fmt.Sprintf("error: %v", err),
			},
		)

		return
	}

	req.JSON(
		http.StatusOK,
		msg,
	)
}
