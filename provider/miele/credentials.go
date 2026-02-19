package miele

import (
	"encoding/json"
	"fmt"
	"os"
)

type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

func LoadCredentials(path string) (*Credentials, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file: %w", err)
	}

	var creds Credentials
	if err := json.Unmarshal(b, &creds); err != nil {
		return nil, fmt.Errorf("failed to parse credentials file: %w", err)
	}

	if creds.ClientID == "" || creds.ClientSecret == "" {
		return nil, fmt.Errorf("credentials file must contain client_id and client_secret")
	}

	return &creds, nil
}
