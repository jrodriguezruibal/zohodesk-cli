package config

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	OrgID        string `yaml:"org_id"`
	Region       string `yaml:"region"`
}

type Config struct {
	Profiles       map[string]Profile `yaml:"profiles"`
	CurrentProfile string             `yaml:"current_profile"`
}

type TokenCache struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	OrgID        string    `json:"org_id"`
}

const (
	DefaultRegion = "com"
)

var RegionURLs = map[string]string{
	"com": "https://desk.zoho.com",
	"eu":  "https://desk.zoho.eu",
	"in":  "https://desk.zoho.in",
	"cn":  "https://desk.zoho.com.cn",
	"au":  "https://desk.zoho.com.au",
}

var AuthURLs = map[string]string{
	"com": "https://accounts.zoho.com",
	"eu":  "https://accounts.zoho.eu",
	"in":  "https://accounts.zoho.in",
	"cn":  "https://accounts.zoho.com.cn",
	"au":  "https://accounts.zoho.com.au",
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, ".config", "zohodesk-cli"), nil
}

func Load() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(configPath, "config.yaml")

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("config file not found: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if len(cfg.Profiles) == 0 {
		return nil, fmt.Errorf("no profiles configured")
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configPath, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configFile := filepath.Join(configPath, "config.yaml")

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

func GetProfile(cfg *Config, name string) (*Profile, error) {
	if name == "" {
		name = cfg.CurrentProfile
	}

	profile, exists := cfg.Profiles[name]
	if !exists {
		return nil, fmt.Errorf("profile '%s' not found", name)
	}

	if profile.Region == ""{
		profile.Region = DefaultRegion
	}

	return &profile, nil
}

func SetProfile(cfg *Config, name string, profile Profile) error {
	if cfg.Profiles == nil {
		cfg.Profiles = make(map[string]Profile)
	}

	cfg.Profiles[name] = profile
	return nil
}

func DeleteProfile(cfg *Config, name string) error {
	if _, exists := cfg.Profiles[name]; !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}

	delete(cfg.Profiles, name)
	return nil
}

func ListProfiles(cfg *Config) []string {
	profiles := make([]string, 0, len(cfg.Profiles))
	for name := range cfg.Profiles {
		profiles = append(profiles, name)
	}
	return profiles
}

func GetTokenCachePath(profile string) (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	cachePath := filepath.Join(configPath, "tokens")
	if err := os.MkdirAll(cachePath, 0700); err != nil {
		return "", fmt.Errorf("failed to create token cache directory: %w", err)
	}

	hash := sha256.Sum256([]byte(profile))
	tokenFile := filepath.Join(cachePath, hex.EncodeToString(hash[:])+".json")

	return tokenFile, nil
}

func LoadTokenCache(profile string) (*TokenCache, error) {
	tokenFile, err := GetTokenCachePath(profile)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("token cache not found: %w", err)
	}

	var cache TokenCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, fmt.Errorf("failed to parse token cache: %w", err)
	}

	return &cache, nil
}

func SaveTokenCache(profile string, cache *TokenCache) error {
	tokenFile, err := GetTokenCachePath(profile)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cache, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal token cache: %w", err)
	}

	if err := os.WriteFile(tokenFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write token cache: %w", err)
	}

	return nil
}

func ClearTokenCache(profile string) error {
	tokenFile, err := GetTokenCachePath(profile)
	if err != nil {
		return err
	}

	if err := os.Remove(tokenFile); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to clear token cache: %w", err)
		}
	}

	return nil
}

func InitConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(configPath, "config.yaml")

	if _, err := os.Stat(configFile); err == nil {
		return Load()
	}

	cfg := &Config{
		Profiles:       make(map[string]Profile),
		CurrentProfile: "default",
	}

	return cfg, nil
}

func Exists() bool {
	configPath, err := getConfigPath()
	if err != nil {
		return false
	}

	configFile := filepath.Join(configPath, "config.yaml")

	_, err = os.Stat(configFile)
	return err == nil
}