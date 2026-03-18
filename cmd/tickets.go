package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/config"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
	"github.com/spf13/cobra"
)

var ticketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "Manage Zoho Desk tickets",
	Long:  `Commands for listing, creating, updating, and searching tickets in Zoho Desk.`,
}

var ticketsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tickets",
	Long:  `List all tickets with optional filters.`,
	RunE:  runTicketsList,
}

var ticketsGetCmd = &cobra.Command{
	Use:   "get <ticket-id>",
	Short: "Get ticket details",
	Long:  `Get detailed information about a specific ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTicketsGet,
}

var ticketsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ticket",
	Long:  `Create a new ticket in Zoho Desk.`,
	RunE:  runTicketsCreate,
}

var ticketsUpdateCmd = &cobra.Command{
	Use:   "update <ticket-id>",
	Short: "Update a ticket",
	Long:  `Update ticket status, priority, or other fields.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTicketsUpdate,
}

var ticketsCloseCmd = &cobra.Command{
	Use:   "close <ticket-id>",
	Short: "Close a ticket",
	Long:  `Close a ticket with an optional resolution message.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTicketsClose,
}

var ticketsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search tickets",
	Long:  `Search tickets by email, status, or other criteria.`,
	RunE:  runTicketsSearch,
}

var (
	flagStatus     string
	flagPriority   string
	flagLimit      int
	flagSubject    string
	flagDescription string
	flagEmail      string
	flagDepartment string
	flagResolution string
	flagFull       bool
	flagQuery      string
	flagFromDate   string
	flagToDate     string
)

func init() {
	rootCmd.AddCommand(ticketsCmd)
	ticketsCmd.AddCommand(ticketsListCmd)
	ticketsCmd.AddCommand(ticketsGetCmd)
	ticketsCmd.AddCommand(ticketsCreateCmd)
	ticketsCmd.AddCommand(ticketsUpdateCmd)
	ticketsCmd.AddCommand(ticketsCloseCmd)
	ticketsCmd.AddCommand(ticketsSearchCmd)

	ticketsListCmd.Flags().StringVarP(&flagStatus, "status", "s", "", "filter by status (Open, Closed, On Hold)")
	ticketsListCmd.Flags().StringVarP(&flagPriority, "priority", "P", "", "filter by priority (Low, Medium, High)")
	ticketsListCmd.Flags().IntVarP(&flagLimit, "limit", "l", 50, "maximum number of tickets to return")

	ticketsGetCmd.Flags().BoolVarP(&flagFull, "full", "f", false, "show full ticket with threads and contact")

	ticketsCreateCmd.Flags().StringVarP(&flagSubject, "subject", "s", "", "ticket subject (required)")
	ticketsCreateCmd.Flags().StringVarP(&flagDescription, "description", "d", "", "ticket description")
	ticketsCreateCmd.Flags().StringVarP(&flagEmail, "email", "e", "", "contact email (required)")
	ticketsCreateCmd.Flags().StringVarP(&flagPriority, "priority", "P", "Medium", "ticket priority (Low, Medium, High)")
	ticketsCreateCmd.Flags().StringVarP(&flagDepartment, "department", "D", "", "department ID")
	ticketsCreateCmd.Flags().StringVarP(&flagResolution, "resolution", "r", "", "resolution message")
	ticketsCreateCmd.MarkFlagRequired("subject")
	ticketsCreateCmd.MarkFlagRequired("email")

	ticketsUpdateCmd.Flags().StringVarP(&flagStatus, "status", "s", "", "new status")
	ticketsUpdateCmd.Flags().StringVarP(&flagPriority, "priority", "P", "", "new priority")
	ticketsUpdateCmd.Flags().StringVarP(&flagResolution, "resolution", "r", "", "resolution message")

	ticketsCloseCmd.Flags().StringVarP(&flagResolution, "resolution", "r", "", "resolution message")

	ticketsSearchCmd.Flags().StringVarP(&flagEmail, "email", "e", "", "search by contact email")
	ticketsSearchCmd.Flags().StringVarP(&flagQuery, "query", "q", "", "search by keyword")
	ticketsSearchCmd.Flags().StringVarP(&flagStatus, "status", "s", "", "filter by status")
	ticketsSearchCmd.Flags().StringVarP(&flagFromDate, "from", "f", "", "from date (YYYY-MM-DD)")
	ticketsSearchCmd.Flags().StringVarP(&flagToDate, "to", "t", "", "to date (YYYY-MM-DD)")
}

func newClient() (*api.Client, string, error) {
	profile, profileName, err := getActiveProfile()
	if err != nil {
		return nil, "", err
	}

	client, err := api.NewClient(profile, profileName)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create API client: %w", err)
	}

	return client, profileName, nil
}

func getActiveProfile() (*config.Profile, string, error) {
	cfg, err := getConfig()
	if err == nil {
		p := getProfile()
		profile, err := config.GetProfile(cfg, p)
		if err == nil {
			return profile, p, nil
		}
	}

	profile, name, err := config.GetActiveProfile(nil, "")
	if err != nil {
		return nil, "", fmt.Errorf("no profile configured. Run 'zohodesk-cli config init' or set ZOHO_CLIENT_ID, ZOHO_CLIENT_SECRET, ZOHO_ORG_ID environment variables")
	}
	return profile, name, nil
}

func runTicketsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := map[string]string{
		"limit": fmt.Sprintf("%d", flagLimit),
	}
	if flagStatus != "" {
		params["status"] = flagStatus
	}
	if flagPriority != "" {
		params["priority"] = flagPriority
	}

	tickets, err := api.NewTicketsService(client).List(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to list tickets: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTickets(tickets, out)
}

func runTicketsGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	out := getOutputFormat()

	if flagFull {
		fullTicket, err := api.NewTicketsService(client).GetFull(context.Background(), ticketID)
		if err != nil {
			return fmt.Errorf("failed to get ticket: %w", err)
		}
		return output.PrintFullTicket(fullTicket, out)
	}

	ticket, err := api.NewTicketsService(client).Get(context.Background(), ticketID)
	if err != nil {
		return fmt.Errorf("failed to get ticket: %w", err)
	}

	return output.PrintTicket(ticket, out)
}

func runTicketsCreate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	req := models.TicketCreateRequest{
		Subject:     flagSubject,
		Description: flagDescription,
		Priority:    flagPriority,
	}

	if flagDepartment != "" {
		req.DepartmentID = flagDepartment
	}

	ticket, err := api.NewTicketsService(client).Create(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTicket(ticket, out)
}

func runTicketsUpdate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	req := models.TicketUpdateRequest{}
	if flagStatus != "" {
		req.Status = flagStatus
	}
	if flagPriority != "" {
		req.Priority = flagPriority
	}
	if flagResolution != "" {
		req.Resolution = flagResolution
	}

	ticket, err := api.NewTicketsService(client).Update(context.Background(), ticketID, req)
	if err != nil {
		return fmt.Errorf("failed to update ticket: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTicket(ticket, out)
}

func runTicketsClose(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	if err := api.NewTicketsService(client).Close(context.Background(), ticketID, flagResolution); err != nil {
		return fmt.Errorf("failed to close ticket: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "ticket_id": "%s", "message": "Ticket closed successfully"}`, ticketID)
	} else {
		fmt.Printf("Ticket %s closed successfully\n", ticketID)
	}
	return nil
}

func runTicketsSearch(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := models.SearchParams{
		Email:  flagEmail,
		Status: flagStatus,
	}

	tickets, err := api.NewTicketsService(client).Search(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to search tickets: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTickets(tickets, out)
}