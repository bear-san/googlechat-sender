package auth

import (
	"context"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckStatus(ctx context.Context, s *gin.Context, secretBase string) (*ent.SystemUser, error) {
	token, err := s.Cookie("token")
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid key")
		}

		return []byte(secretBase), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	id := claims["id"].(string)

	return db.Client.SystemUser.Get(ctx, id)
}
