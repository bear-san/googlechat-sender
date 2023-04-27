package oauth

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetJwks(metadata GoogleMetadata) (*JwkListPayload, error) {
	res, err := http.Get(metadata.JwksUri)
	if err != nil {
		return nil, err
	}

	keyString, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var p JwkListPayload
	err = json.Unmarshal(keyString, &p)

	return &p, err
}
