package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var contactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Manage Zoho Desk contacts",
	Long:  `Commands for listing and searching contacts in Zoho Desk.`,
}

var contactsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts",
	Long:  `List all contacts with optional filters.`,
	RunE:  runContactsList,
}

var contactsGetCmd = &cobra.Command{
	Use:   "get <contact-id>",
	Short: "Get contact details",
	Long:  `Get detailed information about a specific contact.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runContactsGet,
}

var contactsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search contacts",
	Long:  `Search contacts by email or name.`,
	RunE:  runContactsSearch,
}

var (
	contactFlagLimit int
	contactFlagEmail string
	contactFlagName  string
)

func init() {
	rootCmd.AddCommand(contactsCmd)
	contactsCmd.AddCommand(contactsListCmd)
	contactsCmd.AddCommand(contactsGetCmd)
	contactsCmd.AddCommand(contactsSearchCmd)

	contactsListCmd.Flags().IntVarP(&contactFlagLimit, "limit", "l", 50, "maximum number of contacts to return")

	contactsSearchCmd.Flags().StringVarP(&contactFlagEmail, "email", "e", "", "search by email")
	contactsSearchCmd.Flags().StringVarP(&contactFlagName, "name", "n", "", "search by name")
}

func runContactsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := map[string]string{
		"limit": fmt.Sprintf("%d", contactFlagLimit),
	}

	contacts, err := api.NewContactsService(client).List(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to list contacts: %w", err)
	}

	out := getOutputFormat()
	return output.PrintContacts(contacts, out)
}

func runContactsGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	contactID := args[0]

	contact, err := api.NewContactsService(client).Get(context.Background(), contactID)
	if err != nil {
		return fmt.Errorf("failed to get contact: %w", err)
	}

	out := getOutputFormat()
	return output.PrintContact(contact, out)
}

func runContactsSearch(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	contacts, err := api.NewContactsService(client).Search(context.Background(), contactFlagEmail, contactFlagName)
	if err != nil {
		return fmt.Errorf("failed to search contacts: %w", err)
	}

	out := getOutputFormat()
	return output.PrintContacts(contacts, out)
}