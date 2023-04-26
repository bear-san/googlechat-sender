package oauth

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func GetMetadata() (*GoogleMetadata, error) {
	resp, err := http.Get(os.Getenv("GOOGLE_OAUTH_METADATA_PATH"))
	if err != nil {
		return nil, err
	}

	bodyString, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m GoogleMetadata
	if err = json.Unmarshal(bodyString, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

type GoogleMetadata struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	DeviceAuthorizationEndpoint       string   `json:"device_authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	RevocationEndpoint                string   `json:"revocation_endpoint"`
	JwksUri                           string   `json:"jwks_uri"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
}

type GoogleJwk struct {
	Keys []struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kty string `json:"kty"`
		Use string `json:"use"`
		Kid string `json:"kid"`
		N   string `json:"n"`
	} `json:"keys"`
}
