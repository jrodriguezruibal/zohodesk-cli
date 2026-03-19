package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
	"github.com/spf13/cobra"
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Manage accounts",
	Long:  `List, create, update, and delete customer accounts.`,
}

var accountsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List accounts",
	Long:  `List all customer accounts.`,
	RunE:  runAccountsList,
}

var accountsGetCmd = &cobra.Command{
	Use:   "get <account-id>",
	Short: "Get account details",
	Long:  `Get detailed information about a specific account.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAccountsGet,
}

var accountsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an account",
	Long:  `Create a new customer account.`,
	RunE:  runAccountsCreate,
}

var accountsUpdateCmd = &cobra.Command{
	Use:   "update <account-id>",
	Short: "Update an account",
	Long:  `Update an existing customer account.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAccountsUpdate,
}

var accountsDeleteCmd = &cobra.Command{
	Use:   "delete <account-id>",
	Short: "Delete an account",
	Long:  `Delete a customer account.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAccountsDelete,
}

var (
	flagAccountName     string
	flagAccountEmail    string
	flagAccountPhone    string
	flagAccountWebsite  string
	flagAccountType     string
	flagAccountIndustry string
	flagAccountOwner    string
)

func init() {
	rootCmd.AddCommand(accountsCmd)
	accountsCmd.AddCommand(accountsListCmd)
	accountsCmd.AddCommand(accountsGetCmd)
	accountsCmd.AddCommand(accountsCreateCmd)
	accountsCmd.AddCommand(accountsUpdateCmd)
	accountsCmd.AddCommand(accountsDeleteCmd)

	accountsCreateCmd.Flags().StringVarP(&flagAccountName, "name", "n", "", "account name (required)")
	accountsCreateCmd.Flags().StringVarP(&flagAccountEmail, "email", "e", "", "account email")
	accountsCreateCmd.Flags().StringVarP(&flagAccountPhone, "phone", "p", "", "account phone")
	accountsCreateCmd.Flags().StringVarP(&flagAccountWebsite, "website", "w", "", "account website")
	accountsCreateCmd.Flags().StringVarP(&flagAccountType, "type", "t", "", "account type")
	accountsCreateCmd.Flags().StringVarP(&flagAccountIndustry, "industry", "i", "", "account industry")
	accountsCreateCmd.Flags().StringVarP(&flagAccountOwner, "owner", "o", "", "owner ID")

	accountsUpdateCmd.Flags().StringVarP(&flagAccountName, "name", "n", "", "account name")
	accountsUpdateCmd.Flags().StringVarP(&flagAccountEmail, "email", "e", "", "account email")
	accountsUpdateCmd.Flags().StringVarP(&flagAccountPhone, "phone", "p", "", "account phone")
	accountsUpdateCmd.Flags().StringVarP(&flagAccountWebsite, "website", "w", "", "account website")
	accountsUpdateCmd.Flags().StringVarP(&flagAccountType, "type", "t", "", "account type")
	accountsUpdateCmd.Flags().StringVarP(&flagAccountIndustry, "industry", "i", "", "account industry")
	accountsUpdateCmd.Flags().StringVarP(&flagAccountOwner, "owner", "o", "", "owner ID")
}

func runAccountsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	accounts, err := api.NewAccountsService(client).List(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list accounts: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAccounts(accounts, out)
}

func runAccountsGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	accountID := args[0]

	account, err := api.NewAccountsService(client).Get(context.Background(), accountID)
	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAccount(account, out)
}

func runAccountsCreate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	if flagAccountName == "" {
		return fmt.Errorf("--name is required")
	}

	req := models.AccountCreateRequest{
		Name:     flagAccountName,
		Email:    flagAccountEmail,
		Phone:    flagAccountPhone,
		Website:  flagAccountWebsite,
		Type:     flagAccountType,
		Industry: flagAccountIndustry,
		OwnerID:  flagAccountOwner,
	}

	account, err := api.NewAccountsService(client).Create(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAccount(account, out)
}

func runAccountsUpdate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	accountID := args[0]

	req := models.AccountUpdateRequest{
		Name:     flagAccountName,
		Email:    flagAccountEmail,
		Phone:    flagAccountPhone,
		Website:  flagAccountWebsite,
		Type:     flagAccountType,
		Industry: flagAccountIndustry,
		OwnerID:  flagAccountOwner,
	}

	account, err := api.NewAccountsService(client).Update(context.Background(), accountID, req)
	if err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAccount(account, out)
}

func runAccountsDelete(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	accountID := args[0]

	if err := api.NewAccountsService(client).Delete(context.Background(), accountID); err != nil {
		return fmt.Errorf("failed to delete account: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "accountId": "%s", "message": "Account deleted successfully"}`, accountID)
	} else {
		fmt.Printf("Account %s deleted successfully\n", accountID)
	}
	return nil
}