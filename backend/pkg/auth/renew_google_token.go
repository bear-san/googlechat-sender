package auth

import (
	"context"
	"encoding/json"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func RenewGoogleToken(ctx context.Context, t *ent.GoogleApiKey, clientInfo oauth.ClientInfo) (*ent.GoogleApiKey, error) {
	oauthMetadata, err := oauth.GetMetadata(os.Getenv("GOOGLE_OAUTH_METADATA_PATH"))
	if err != nil {
		return nil, err
	}

	val := url.Values{}
	val.Set("client_id", clientInfo.ClientId)
	val.Set("client_secret", clientInfo.ClientSecret)
	val.Set("grant_type", "refresh_token")
	val.Set("refresh_token", t.RefreshToken)

	resp, err := http.PostForm(oauthMetadata.TokenEndpoint, val)
	if err != nil {
		return nil, err
	}

	respTxt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p RefreshTokenPayload
	err = json.Unmarshal(respTxt, &p)

	if err != nil {
		return nil, err
	}

	return db.Client.GoogleApiKey.UpdateOneID(t.ID).
		SetAccessToken(p.AccessToken).
		SetExpirationDate(time.Now().Add(time.Duration(p.ExpiresIn) * time.Second)).
		Save(ctx)
}

type RefreshTokenPayload struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
