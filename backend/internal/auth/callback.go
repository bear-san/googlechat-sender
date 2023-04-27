package auth

import (
	"encoding/json"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Callback(req *gin.Context) {
	metadata, err := oauth.GetMetadata(os.Getenv("GOOGLE_OAUTH_METADATA_PATH"))
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

	token, err := oauth.GetToken(
		req.Query("code"),
		oauth.ClientInfo{
			ClientId:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			Scope:        os.Getenv("OAUTH_SCOPE"),
			RedirectUri:  fmt.Sprintf("%s/api/auth/callback", os.Getenv("SERVER_HOST")),
			Domains:      os.Getenv("GOOGLE_DOMAIN_RESTRICTION"),
		},
		*metadata,
	)

	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to get credential from google",
			},
		)

		return
	}

	var cred OAuthCredential

	err = json.Unmarshal([]byte(*token), &cred)
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "invalid credential payload",
			},
		)

		return
	}

	jwks, err := oauth.GetJwks(*metadata)
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to get public key from Google IdP",
			},
		)

		return
	}

	claims, err := oauth.ParseIdToken(cred.IdToken, *jwks)

	req.JSON(http.StatusOK, claims)
}

type OAuthCredential struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token"`
}
