package auth

import (
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Login(req *gin.Context) {
	metadata, err := oauth.GetMetadata(os.Getenv("GOOGLE_OAUTH_METADATA_PATH"))
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to get OAuth metadata",
			},
		)

		return
	}

	redirectUrl, err := auth.CreateLoginUrl(
		*metadata,
		GetOAuthClientInfo(),
	)

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

	req.Redirect(
		http.StatusFound,
		*redirectUrl,
	)
}
