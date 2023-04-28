package auth

import (
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
	"os"
)

func GetOAuthClientInfo() oauth.ClientInfo {
	return oauth.ClientInfo{
		ClientId:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		Scope:        os.Getenv("OAUTH_SCOPE"),
		RedirectUri:  fmt.Sprintf("%s/api/auth/callback", os.Getenv("SERVER_HOST")),
		Domains:      os.Getenv("GOOGLE_DOMAIN_RESTRICTION"),
	}
}
