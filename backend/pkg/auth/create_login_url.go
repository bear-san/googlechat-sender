package auth

import (
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
)

func CreateLoginUrl(metadata oauth.GoogleMetadata, info oauth.ClientInfo) (*string, error) {
	authUrl := fmt.Sprintf(
		"%s?client_id=%s&scope=%s&hd=%s&redirect_uri=%s&response_type=code&access_type=offline&prompt=consent",
		metadata.AuthorizationEndpoint,
		info.ClientId,
		info.Scope,
		info.Domains,
		info.RedirectUri,
	)

	return &authUrl, nil
}
