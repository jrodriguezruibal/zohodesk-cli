package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/config"
)

type Auth struct {
	config *config.Profile
}

func NewAuth(cfg *config.Profile) *Auth {
	return &Auth{config: cfg}
}

func (a *Auth) GetAccessToken(ctx context.Context) (*TokenResponse, error) {
	authURL, ok := config.AuthURLs[a.config.Region]
	if !ok {
		authURL = config.AuthURLs["com"]
	}

	tokenURL := authURL + "/oauth/v2/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", a.config.ClientID)
	data.Set("client_secret", a.config.ClientSecret)
	data.Set("scope","Desk.tickets.ALL,Desk.contacts.READ,Desk.basic.READ")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth request: %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("auth request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errResp)
		return nil, fmt.Errorf("auth failed (%d): %v", resp.StatusCode, errResp)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	return &tokenResp, nil
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope,omitempty"`
}

func (a *Auth) ValidateCredentials(ctx context.Context) error {
	_, err := a.GetAccessToken(ctx)
	return err
}