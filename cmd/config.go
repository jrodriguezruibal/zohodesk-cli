package cmd

import (
		"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/config"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/tui"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
		Long:  `Manage zohodesk-cli configuration and profiles.`,
}

var configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration",
		Long:  `Initialize configuration with interactive setup wizard.`,
		RunE:  runConfigInit,
}

var configListCmd = &cobra.Command{
		Use:   "list",
		Short: "List profiles",
		Long:  `List all configured profiles.`,
		RunE:  runConfigList,
}

var configSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set configuration values",
		Long:  `Set configuration values for a profile.`,
		RunE:  runConfigSet,
}

var configDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a profile",
		Long:  `Delete a configuration profile.`,
		RunE:  runConfigDelete,
}

var configUseCmd = &cobra.Command{
		Use:   "use",
		Short: "Set default profile",
		Long:  `Set the default profile to use.`,
		RunE:  runConfigUse,
}

var configClearTokensCmd = &cobra.Command{
		Use:   "clear-tokens",
		Short: "Clear cached tokens",
		Long:  `Clear cached authentication tokens.`,
		RunE:  runConfigClearTokens,
}

var (
	configFlagProfile     string
	configFlagClientID    string
	configFlagClientSecret string
	configFlagOrgID       string
	configFlagRegion      string
)

func init() {
		rootCmd.AddCommand(configCmd)
		configCmd.AddCommand(configInitCmd)
		configCmd.AddCommand(configListCmd)
		configCmd.AddCommand(configSetCmd)
		configCmd.AddCommand(configDeleteCmd)
		configCmd.AddCommand(configUseCmd)
		configCmd.AddCommand(configClearTokensCmd)

		configSetCmd.Flags().StringVarP(&configFlagProfile, "profile", "p", "default", "profile name")
		configSetCmd.Flags().StringVarP(&configFlagClientID, "client-id", "c", "", "Zoho Client ID")
		configSetCmd.Flags().StringVarP(&configFlagClientSecret, "client-secret", "s", "", "Zoho Client Secret")
		configSetCmd.Flags().StringVarP(&configFlagOrgID, "org-id", "o", "", "Zoho Organization ID")
		configSetCmd.Flags().StringVarP(&configFlagRegion, "region", "r", "com", "Zoho region (com, eu, in, cn, au)")

		configDeleteCmd.Flags().StringVarP(&configFlagProfile, "profile", "p", "", "profile name to delete")
		configDeleteCmd.MarkFlagRequired("profile")

		configUseCmd.Flags().StringVarP(&configFlagProfile, "profile", "p", "", "profile name to set as default")
		configUseCmd.MarkFlagRequired("profile")

		configClearTokensCmd.Flags().StringVarP(&configFlagProfile, "profile", "p", "", "profile name (default: all)")
}

func runConfigInit(cmd *cobra.Command, args []string) error {
		return tui.RunSetup()
}

func runConfigList(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
	if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
	}

		fmt.Printf("Current profile: %s\n\n", cfg.CurrentProfile)
		fmt.Println("Profiles:")
		for name, profile := range cfg.Profiles {
				marker := " "
				if name == cfg.CurrentProfile {
						marker = "*"
				}
				fmt.Printf("  %s %s (region: %s)\n", marker, name, profile.Region)
		}

		return nil
}

func runConfigSet(cmd *cobra.Command, args []string) error {
		cfg, err := config.InitConfig()
	if err != nil {
			return fmt.Errorf("failed to init config: %w", err)
	}

		profileName := configFlagProfile
		if profileName == "" {
				profileName = "default"
		}

		existingProfile, exists := cfg.Profiles[profileName]

		newProfile := config.Profile{
				ClientID:     configFlagClientID,
				ClientSecret: configFlagClientSecret,
				OrgID:        configFlagOrgID,
				Region:       configFlagRegion,
		}

		if exists {
				if newProfile.ClientID == "" {
						newProfile.ClientID = existingProfile.ClientID
				}
				if newProfile.ClientSecret == "" {
						newProfile.ClientSecret = existingProfile.ClientSecret
				}
				if newProfile.OrgID == "" {
						newProfile.OrgID = existingProfile.OrgID
				}
				if newProfile.Region == "" {
						newProfile.Region = existingProfile.Region
				}
		}

		if newProfile.Region == "" {
				newProfile.Region = "com"
		}

		if err := config.ValidateRegion(newProfile.Region); err != nil {
				return err
		}

		cfg.Profiles[profileName] = newProfile

		if cfg.CurrentProfile == "" {
				cfg.CurrentProfile = profileName
		}

		if err := config.Save(cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
	}

		fmt.Printf("Profile '%s' configured successfully\n", profileName)
		if newProfile.ClientID != "" {
				fmt.Printf("Client ID: %s...%s\n", newProfile.ClientID[:10], newProfile.ClientID[len(newProfile.ClientID)-4:])
		}
		fmt.Printf("Region: %s\n", newProfile.Region)

		return nil
}

func runConfigDelete(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
	if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
	}

		profileName := configFlagProfile

		if err := config.DeleteProfile(cfg, profileName); err != nil {
				return err
		}

		if cfg.CurrentProfile == profileName {
				for name := range cfg.Profiles {
						cfg.CurrentProfile = name
						break
				}
		}

		if err := config.Save(cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
	}

		fmt.Printf("Profile '%s' deleted successfully\n", profileName)
		return nil
}

func runConfigUse(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
	if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
	}

		profileName := configFlagProfile

		if _, exists := cfg.Profiles[profileName]; !exists {
				return fmt.Errorf("profile '%s' not found", profileName)
		}

		cfg.CurrentProfile = profileName

		if err := config.Save(cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
	}

		fmt.Printf("Default profile set to '%s'\n", profileName)
		return nil
}

func runConfigClearTokens(cmd *cobra.Command, args []string) error {
		profileName := configFlagProfile

		if profileName != "" {
				if err := config.ClearTokenCache(profileName); err != nil {
						return fmt.Errorf("failed to clear token cache: %w", err)
				}
				fmt.Printf("Token cache cleared for profile '%s'\n", profileName)
		} else {
				cfg, err := config.Load()
				if err != nil {
						return fmt.Errorf("failed to load config: %w", err)
				}
				for name := range cfg.Profiles {
						config.ClearTokenCache(name)
				}
				fmt.Println("All token caches cleared")
		}

		return nil
}