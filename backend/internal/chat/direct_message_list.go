package chat

import (
	"context"
	auth2 "github.com/bear-san/googlechat-sender/backend/internal/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/chat"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func DirectMessageList(req *gin.Context) {
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

	act, err := auth.GetGoogleCredential(ctx, u, auth2.GetOAuthClientInfo())
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "internal server error",
			},
		)

		return
	}

	spaces, err := chat.GetDirectMessages(ctx, act, os.Getenv("GOOGLE_USER_LIST_ID"))

	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "internal server error",
			},
		)

		return
	}

	req.JSON(
		http.StatusOK,
		spaces,
	)
}
