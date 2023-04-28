package auth

import (
	"context"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/bear-san/googlechat-sender/backend/pkg/oauth"
	"time"
)

func GetGoogleCredential(ctx context.Context, u *ent.SystemUser, clientInfo oauth.ClientInfo) (*ent.GoogleApiKey, error) {
	cred, err := db.Client.GoogleApiKey.Get(ctx, u.ID)
	if cred.ExpirationDate.Unix() < time.Now().Unix() {
		cred, err = RenewGoogleToken(ctx, cred, clientInfo)
	}

	return cred, err
}
