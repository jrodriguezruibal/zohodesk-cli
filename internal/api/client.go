package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/config"
	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

const (
 DefaultTimeout =30 * time.Second
 MaxRetries     =3
 RetryDelay     = 1 * time.Second
)

type Client struct {
	httpClient *http.Client
	config     *config.Profile
	profileName string
	tokenCache *config.TokenCache
}

func NewClient(cfg *config.Profile, profileName string) (*Client, error) {
	return &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		config:      cfg,
		profileName: profileName,
	}, nil
}

func (c *Client) ensureToken(ctx context.Context) error {
	if c.tokenCache != nil && time.Now().Before(c.tokenCache.ExpiresAt) {
		return nil
	}

	cache, err := config.LoadTokenCache(c.profileName)
	if err == nil && time.Now().Before(cache.ExpiresAt) {
		c.tokenCache = cache
		return nil
	}

	auth := NewAuth(c.config)
	token, err := auth.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	cache = &config.TokenCache{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
		OrgID:        c.config.OrgID,
	}

	if err := config.SaveTokenCache(c.profileName, cache); err != nil {
		return fmt.Errorf("failed to save token cache: %w", err)
	}

	c.tokenCache = cache
	return nil
}

func (c *Client) doRequest(ctx context.Context, method, endpoint string, params url.Values, body interface{}) (*http.Response, error) {
	if err := c.ensureToken(ctx); err != nil {
		return nil, err
	}

	baseURL, ok := config.RegionURLs[c.config.Region]
	if !ok {
		baseURL = config.RegionURLs["com"]
	}

	fullURL := baseURL + "/api/v1" + endpoint
	if len(params) > 0 {
		fullURL = fullURL + "?" + params.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Zoho-oauthtoken "+c.tokenCache.AccessToken)
	req.Header.Set("orgId", c.config.OrgID)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	var lastErr error

	for i := 0; i < MaxRetries; i++ {
		resp, lastErr = c.httpClient.Do(req)
		if lastErr == nil && resp.StatusCode <500 {
			break
		}

		if resp != nil && resp.StatusCode ==401 {
			c.tokenCache = nil
			if err := c.ensureToken(ctx); err != nil {
				return nil, err
			}
			req.Header.Set("Authorization", "Zoho-oauthtoken "+c.tokenCache.AccessToken)
		}

		time.Sleep(RetryDelay)
	}

	return resp, lastErr
}

func (c *Client) parseResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >=400 {
		var apiErr models.APIError
		if err := json.Unmarshal(body, &apiErr); err == nil {
			return fmt.Errorf("API error (%d): %s - %s", resp.StatusCode, apiErr.Code, apiErr.Message)
		}
		return fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}

func (c *Client) Get(ctx context.Context, endpoint string, params url.Values) ([]byte, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *Client) Post(ctx context.Context, endpoint string, body interface{}) ([]byte, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *Client) Patch(ctx context.Context, endpoint string, body interface{}) ([]byte, error) {
	resp, err := c.doRequest(ctx, http.MethodPatch, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *Client) Delete(ctx context.Context, endpoint string) error {
	resp, err := c.doRequest(ctx, http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >=400 {
		return fmt.Errorf("delete failed with status %d", resp.StatusCode)
	}

	return nil
}