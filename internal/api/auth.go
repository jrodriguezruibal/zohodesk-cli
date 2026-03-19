package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (a *Auth) ExchangeCode(ctx context.Context, code string) (*TokenResponse, error) {
	authURL, ok := config.AuthURLs[a.config.Region]
	if !ok {
		authURL = config.AuthURLs["com"]
	}

	tokenURL := authURL + "/oauth/v2/token"

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", a.config.ClientID)
	data.Set("client_secret", a.config.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", "self")

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		json.Unmarshal(body, &errResp)
		return nil, fmt.Errorf("auth failed (%d): %v - body: %s", resp.StatusCode, errResp, string(body))
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w - body: %s", err, string(body))
	}

	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("empty access token in response: %s", string(body))
	}

	tokenResp.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	return &tokenResp, nil
}

func (a *Auth) RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	authURL, ok := config.AuthURLs[a.config.Region]
	if !ok {
		authURL = config.AuthURLs["com"]
	}

	tokenURL := authURL + "/oauth/v2/token"

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", a.config.ClientID)
	data.Set("client_secret", a.config.ClientSecret)
	data.Set("refresh_token", refreshToken)
	data.Set("redirect_uri", "self")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh request: %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("refresh request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		json.Unmarshal(body, &errResp)
		return nil, fmt.Errorf("refresh failed (%d): %v - body: %s", resp.StatusCode, errResp, string(body))
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w - body: %s", err, string(body))
	}

	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("empty access token in response: %s", string(body))
	}

	tokenResp.RefreshToken = refreshToken
	tokenResp.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	return &tokenResp, nil
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
	data.Set("org_id", a.config.OrgID)
	data.Set("scope", "Desk.tickets.ALL,Desk.contacts.READ,Desk.basic.READ")

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		json.Unmarshal(body, &errResp)
		return nil, fmt.Errorf("auth failed (%d): %v - body: %s", resp.StatusCode, errResp, string(body))
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w - body: %s", err, string(body))
	}

	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("empty access token in response: %s", string(body))
	}

	tokenResp.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	return &tokenResp, nil
}

type TokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiresIn    int       `json:"expires_in"`
	TokenType    string    `json:"token_type"`
	Scope        string    `json:"scope,omitempty"`
	ExpiresAt    time.Time `json:"-"`
}

func (a *Auth) ValidateCredentials(ctx context.Context) error {
	_, err := a.GetAccessToken(ctx)
	return err
}