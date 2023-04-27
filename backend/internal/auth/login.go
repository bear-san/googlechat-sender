package auth

import (
	"fmt"
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
		oauth.ClientInfo{
			ClientId:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			Scope:        os.Getenv("OAUTH_SCOPE"),
			RedirectUri:  fmt.Sprintf("%s/api/auth/callback", os.Getenv("SERVER_HOST")),
			Domains:      os.Getenv("GOOGLE_DOMAIN_RESTRICTION"),
		},
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
