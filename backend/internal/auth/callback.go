package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
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

	ctx := context.Background()

	var user *ent.SystemUser

	user, err = db.Client.SystemUser.Get(ctx, (*claims)["sub"].(string))
	if err != nil {
		user, err = db.Client.SystemUser.Create().
			SetID((*claims)["sub"].(string)).
			SetName((*claims)["name"].(string)).
			SetEmail((*claims)["email"].(string)).
			Save(ctx)
	}

	uid := (*claims)["sub"].(string)

	_, err = db.Client.GoogleApiKey.Get(ctx, uid)
	if err != nil {
		_, err = db.Client.GoogleApiKey.Create().
			SetID(uid).
			SetAccessToken(cred.AccessToken).
			SetRefreshToken(cred.RefreshToken).
			Save(ctx)
	} else {
		_, err = db.Client.GoogleApiKey.UpdateOneID(uid).
			SetAccessToken(cred.AccessToken).
			SetRefreshToken(cred.RefreshToken).
			Save(ctx)
	}

	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": fmt.Sprintf("failed to create new user, %v", err),
			},
		)

		return
	}

	expTimeStmp := time.Now().Add(time.Hour * 3).Unix()

	localTokenClaims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   expTimeStmp,
		"aud":   os.Getenv("SERVER_HOST"),
		"iss":   os.Getenv("SERVER_HOST"),
	}

	localToken := jwt.NewWithClaims(jwt.SigningMethodHS256, localTokenClaims)
	localTokenString, err := localToken.SignedString([]byte(os.Getenv("SECRET_BASE")))

	req.SetCookie(
		"token",
		localTokenString,
		3600*3,
		"/",
		"",
		true,
		true,
	)

	req.Redirect(
		http.StatusFound,
		"/",
	)
}

type OAuthCredential struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token"`
}
