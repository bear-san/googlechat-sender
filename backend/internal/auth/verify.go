package auth

import (
	"context"
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Verify(req *gin.Context) {
	ctx := context.Background()
	_, err := auth.CheckStatus(ctx, req, os.Getenv("SECRET_BASE"))
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

	req.Status(http.StatusNoContent)
}
