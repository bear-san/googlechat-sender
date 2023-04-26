package auth

import (
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
	"os"
)

func CreateLoginUrl() (*string, error) {
	metadata, err := oauth.GetMetadata()
	if err != nil {
		return nil, err
	}

	authUrl := fmt.Sprintf(
		"%s?client_id=%s&scope=%s&hd=%s&redirect_uri=%s",
		metadata.AuthorizationEndpoint,
		os.Getenv("OAUTH_CLIENT_ID"),
		"openid profile email",
		os.Getenv("GOOGLE_DOMAIN_RESTRICTION"),
		fmt.Sprintf("%s/api/auth/callback", os.Getenv("SERVER_HOST")),
	)

	return &authUrl, nil
}
