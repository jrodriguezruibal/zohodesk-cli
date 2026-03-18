package config

import (
	"fmt"
	"os"
	"strings"
)

func GetFromEnv() (*Profile, error) {
	clientID := os.Getenv("ZOHO_CLIENT_ID")
	clientSecret := os.Getenv("ZOHO_CLIENT_SECRET")
	orgID := os.Getenv("ZOHO_ORG_ID")
	region := os.Getenv("ZOHO_REGION")

	if clientID == "" || clientSecret == "" || orgID == "" {
		return nil, fmt.Errorf("missing required environment variables: ZOHO_CLIENT_ID, ZOHO_CLIENT_SECRET, ZOHO_ORG_ID")
	}

	if region == "" {
		region = DefaultRegion
	}

	return &Profile{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		OrgID:        orgID,
		Region:       region,
	}, nil
}

func GetActiveProfile(cfg *Config, profileName string) (*Profile, string, error) {
	envProfile, err := GetFromEnv()
	if err == nil {
		return envProfile, "env", nil
	}

	if !Exists() {
		return nil, "", fmt.Errorf("no configuration found. Run 'zohodesk-cli config init' or set environment variables")
	}

	if cfg == nil {
		cfg, err = Load()
		if err != nil {
			return nil, "", fmt.Errorf("failed to load config: %w", err)
		}
	}

	name := profileName
	if name == "" {
		name = cfg.CurrentProfile
	}

	profile, err := GetProfile(cfg, name)
	if err != nil {
		return nil, "", err
	}

	return profile, name, nil
}

func ValidateRegion(region string) error {
	validRegions := []string{"com", "eu", "in", "cn", "au"}
	for _, r := range validRegions {
		if r == region {
			return nil
		}
	}
	return fmt.Errorf("invalid region '%s'. Valid regions: %s", region, strings.Join(validRegions, ", "))
}