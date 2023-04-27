package oauth

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/big"
)

func ParseIdToken(tokenString string, jwks JwkListPayload) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("invalid Sign Method")
		}

		var keyDict JwkPayload
		for _, key := range jwks.Keys {
			if key.Kid != token.Header["kid"] {
				continue
			}

			keyDict = key
			break
		}

		n, err := base64.RawURLEncoding.DecodeString(keyDict.N)
		if err != nil {
			return nil, fmt.Errorf("invalid key")
		}

		rsaKey := &rsa.PublicKey{
			N: new(big.Int).SetBytes(n),
			E: 65537,
		}

		return rsaKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	i, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claim")
	}

	return &i, nil
}

type JwkListPayload struct {
	Keys []JwkPayload `json:"keys"`
}

type JwkPayload struct {
	Use string `json:"use"`
	N   string `json:"n"`
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kty string `json:"kty"`
	Kid string `json:"kid"`
}
