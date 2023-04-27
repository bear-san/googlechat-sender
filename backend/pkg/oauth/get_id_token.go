package oauth

import (
	"io"
	"net/http"
	"net/url"
)

func GetToken(token string, clientInfo ClientInfo, metadata GoogleMetadata) (*string, error) {
	val := url.Values{}
	val.Set("client_id", clientInfo.ClientId)
	val.Set("client_secret", clientInfo.ClientSecret)
	val.Set("scope", clientInfo.Scope)
	val.Set("grant_type", "authorization_code")
	val.Set("redirect_uri", clientInfo.RedirectUri)
	val.Set("code", token)

	resp, err := http.PostForm(metadata.TokenEndpoint, val)

	defer resp.Body.Close()

	res_s, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	str := string(res_s)

	return &str, nil
}
