package cmd

import (
	"fmt"
	"os"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgVersion   string
	cfgBuildTime string
	cfgFile      string
	profile      string
	outputFormat string
)

var rootCmd = &cobra.Command{
	Use:   "zohodesk-cli",
	Short: "CLI for Zoho Desk API",
	Long: `A command-line interface for Zoho Desk API.

Designed for both human interaction and AI agent integration.
Supports multiple profiles, output formats (JSON, YAML, Table), and batch operations.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func SetVersion(v, bt string) {
	cfgVersion = v
	cfgBuildTime = bt
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/zohodesk-cli/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "", "profile to use (default is 'default')")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format (json, yaml, table)")

	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configPath := home + "/.config/zohodesk-cli"
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("zoho")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}

func getProfile() string {
	if profile != "" {
		return profile
	}
	if viper.IsSet("current_profile") {
		return viper.GetString("current_profile")
	}
	if viper.IsSet("profile") {
		return viper.GetString("profile")
	}
	return "default"
}

func getOutputFormat() string {
	if outputFormat != "" {
		return outputFormat
	}
	return viper.GetString("output")
}

func getConfig() (*config.Config, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w. Run 'zohodesk-cli config init' first", err)
	}

	p := getProfile()
	_, exists := cfg.Profiles[p]
	if !exists {
		return nil, fmt.Errorf("profile '%s' not found. Run 'zohodesk-cli config init' or 'zohodesk-cli config set --profile %s'", p, p)
	}

	cfg.CurrentProfile = p
	return cfg, nil
}