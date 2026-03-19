package cmd

import (
	"context"
	"encoding/json"
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
	Long:  `Create a new ticket in Zoho Desk. Use --batch to create multiple tickets from JSON input.`,
	RunE:  runTicketsCreate,
}

var ticketsUpdateCmd = &cobra.Command{
	Use:   "update <ticket-id>",
	Short: "Update a ticket",
	Long:  `Update ticket status, priority, or other fields. Use --batch to update multiple tickets from JSON input.`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runTicketsUpdate,
}

var ticketsCloseCmd = &cobra.Command{
	Use:   "close <ticket-id>",
	Short: "Close a ticket",
	Long:  `Close a ticket with an optional resolution message. Use --batch to close multiple tickets from JSON input.`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runTicketsClose,
}

var ticketsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search tickets",
	Long:  `Search tickets by email, status, or other criteria.`,
	RunE:  runTicketsSearch,
}

var ticketsAssignCmd = &cobra.Command{
	Use:   "assign <ticket-id>",
	Short: "Assign a ticket",
	Long:  `Assign a ticket to an agent and/or department.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTicketsAssign,
}

var ticketsMergeCmd = &cobra.Command{
	Use:   "merge <source-id> <target-id>",
	Short: "Merge tickets",
	Long:  `Merge a source ticket into a target ticket.`,
	Args:  cobra.ExactArgs(2),
	RunE:  runTicketsMerge,
}

var ticketsFollowCmd = &cobra.Command{
	Use:   "follow <ticket-id>",
	Short: "Follow a ticket",
	Long:  `Follow a ticket to receive notifications.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTicketsFollow,
}

var ticketsUnfollowCmd = &cobra.Command{
	Use:   "unfollow <ticket-id>",
	Short: "Unfollow a ticket",
	Long:  `Stop following a ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTicketsUnfollow,
}

var (
	flagStatus      string
	flagPriority    string
	flagLimit       int
	flagSubject     string
	flagDescription string
	flagEmail       string
	flagContactID string
	flagDepartment  string
	flagResolution  string
	flagFull        bool
	flagQuery       string
	flagFromDate    string
	flagToDate      string
	flagBatch       bool
	flagBatchFile   string
	flagContext     bool
	flagAssignee    string
	flagCustomFields string
)

func init() {
	rootCmd.AddCommand(ticketsCmd)
	ticketsCmd.AddCommand(ticketsListCmd)
	ticketsCmd.AddCommand(ticketsGetCmd)
	ticketsCmd.AddCommand(ticketsCreateCmd)
	ticketsCmd.AddCommand(ticketsUpdateCmd)
	ticketsCmd.AddCommand(ticketsCloseCmd)
	ticketsCmd.AddCommand(ticketsSearchCmd)
	ticketsCmd.AddCommand(ticketsAssignCmd)
	ticketsCmd.AddCommand(ticketsMergeCmd)
	ticketsCmd.AddCommand(ticketsFollowCmd)
	ticketsCmd.AddCommand(ticketsUnfollowCmd)

	ticketsListCmd.Flags().StringVarP(&flagStatus, "status", "s", "", "filter by status (Open, Closed, On Hold)")
	ticketsListCmd.Flags().StringVarP(&flagPriority, "priority", "P", "", "filter by priority (Low, Medium, High)")
	ticketsListCmd.Flags().IntVarP(&flagLimit, "limit", "l", 50, "maximum number of tickets to return")

	ticketsGetCmd.Flags().BoolVarP(&flagFull, "full", "f", false, "show full ticket with threads and contact")
	ticketsGetCmd.Flags().BoolVarP(&flagContext, "context", "C", false, "show ticket with full context (department, assignee, SLA)")

	ticketsCreateCmd.Flags().StringVarP(&flagSubject, "subject", "s", "", "ticket subject (required)")
	ticketsCreateCmd.Flags().StringVarP(&flagDescription, "description", "d", "", "ticket description")
	ticketsCreateCmd.Flags().StringVarP(&flagEmail, "email", "e", "", "contact email")
	ticketsCreateCmd.Flags().StringVar(&flagContactID, "contact-id", "", "contact ID")
	ticketsCreateCmd.Flags().StringVarP(&flagPriority, "priority", "P", "Medium", "ticket priority (Low, Medium, High)")
	ticketsCreateCmd.Flags().StringVarP(&flagDepartment, "department", "D", "", "department ID (required)")
	ticketsCreateCmd.Flags().StringVarP(&flagResolution, "resolution", "r", "", "resolution message")
	ticketsCreateCmd.Flags().BoolVarP(&flagBatch, "batch", "b", false, "read batch input from stdin (JSON array)")
	ticketsCreateCmd.Flags().StringVarP(&flagBatchFile, "file", "f", "", "read batch input from file (JSON array)")

	ticketsUpdateCmd.Flags().StringVarP(&flagStatus, "status", "s", "", "new status")
	ticketsUpdateCmd.Flags().StringVarP(&flagPriority, "priority", "P", "", "new priority")
	ticketsUpdateCmd.Flags().StringVarP(&flagResolution, "resolution", "r", "", "resolution message")
	ticketsUpdateCmd.Flags().StringVarP(&flagCustomFields, "custom-fields", "", "", "custom fields as JSON")
	ticketsUpdateCmd.Flags().BoolVarP(&flagBatch, "batch", "b", false, "read batch input from stdin (JSON array)")
	ticketsUpdateCmd.Flags().StringVarP(&flagBatchFile, "file", "f", "", "read batch input from file (JSON array)")

	ticketsCloseCmd.Flags().StringVarP(&flagResolution, "resolution", "r", "", "resolution message")
	ticketsCloseCmd.Flags().BoolVarP(&flagBatch, "batch", "b", false, "read batch input from stdin (JSON array)")
	ticketsCloseCmd.Flags().StringVarP(&flagBatchFile, "file", "f", "", "read batch input from file (JSON array)")

	ticketsSearchCmd.Flags().StringVarP(&flagEmail, "email", "e", "", "search by contact email")
	ticketsSearchCmd.Flags().StringVarP(&flagQuery, "query", "q", "", "search by keyword")
	ticketsSearchCmd.Flags().StringVarP(&flagStatus, "status", "s", "", "filter by status")
	ticketsSearchCmd.Flags().StringVarP(&flagFromDate, "from", "f", "", "from date (YYYY-MM-DD)")
	ticketsSearchCmd.Flags().StringVarP(&flagToDate, "to", "t", "", "to date (YYYY-MM-DD)")

	ticketsAssignCmd.Flags().StringVarP(&flagAssignee, "agent", "a", "", "agent ID to assign")
	ticketsAssignCmd.Flags().StringVarP(&flagDepartment, "department", "D", "", "department ID")
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

	if flagContext {
		ctxTicket, err := api.NewTicketsService(client).GetWithContext(context.Background(), ticketID)
		if err != nil {
			return fmt.Errorf("failed to get ticket context: %w", err)
		}
		return output.PrintTicketContext(ctxTicket, out)
	}

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
	if flagBatch || flagBatchFile != "" {
		return runBatchCreate()
	}

	client, _, err := newClient()
	if err != nil {
		return err
	}

	req := models.TicketCreateRequest{
		Subject:     flagSubject,
		Description: flagDescription,
		Priority:    flagPriority,
	}

	if flagContactID != "" {
		req.ContactID = flagContactID
	} else if flagEmail != "" {
		contact, err := api.NewContactsService(client).SearchByEmail(context.Background(), flagEmail)
		if err != nil {
			return fmt.Errorf("failed to find contact by email: %w", err)
		}
		if contact == nil {
			return fmt.Errorf("no contact found with email '%s'. Please provide --contact-id or create the contact first", flagEmail)
		}
		req.ContactID = contact.ID
	}

	if flagDepartment != "" {
		req.DepartmentID = flagDepartment
	}

	if req.DepartmentID == "" {
		return fmt.Errorf("department ID is required. Use --department flag")
	}

	if req.ContactID == "" {
		return fmt.Errorf("contact ID is required. Use --contact-id flag or --email to search for a contact")
	}

	ticket, err := api.NewTicketsService(client).Create(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTicket(ticket, out)
}

func runTicketsUpdate(cmd *cobra.Command, args []string) error {
	if flagBatch || flagBatchFile != "" {
		return runBatchUpdate()
	}

	if len(args) == 0 {
		return fmt.Errorf("ticket-id is required when not using --batch")
	}

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
	if flagCustomFields != "" {
		var customFields map[string]interface{}
		if err := json.Unmarshal([]byte(flagCustomFields), &customFields); err != nil {
			return fmt.Errorf("invalid custom fields JSON: %w", err)
		}
		req.CustomFields = customFields
	}

	ticket, err := api.NewTicketsService(client).Update(context.Background(), ticketID, req)
	if err != nil {
		return fmt.Errorf("failed to update ticket: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTicket(ticket, out)
}

func runTicketsClose(cmd *cobra.Command, args []string) error {
	if flagBatch || flagBatchFile != "" {
		return runBatchClose()
	}

	if len(args) == 0 {
		return fmt.Errorf("ticket-id is required when not using --batch")
	}

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

func runTicketsAssign(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	if flagAssignee == "" && flagDepartment == "" {
		return fmt.Errorf("at least one of --agent or --department is required")
	}

	ticket, err := api.NewTicketsService(client).Assign(context.Background(), ticketID, flagAssignee, flagDepartment)
	if err != nil {
		return fmt.Errorf("failed to assign ticket: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTicket(ticket, out)
}

func runTicketsMerge(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	sourceID := args[0]
	targetID := args[1]

	if err := api.NewTicketsService(client).Merge(context.Background(), sourceID, targetID); err != nil {
		return fmt.Errorf("failed to merge tickets: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "sourceId": "%s", "targetId": "%s", "message": "Tickets merged successfully"}`, sourceID, targetID)
	} else {
		fmt.Printf("Ticket %s merged into %s successfully\n", sourceID, targetID)
	}
	return nil
}

func runTicketsFollow(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	if err := api.NewTicketsService(client).Follow(context.Background(), ticketID); err != nil {
		return fmt.Errorf("failed to follow ticket: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "ticketId": "%s", "message": "Now following ticket"}`, ticketID)
	} else {
		fmt.Printf("Now following ticket %s\n", ticketID)
	}
	return nil
}

func runTicketsUnfollow(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	if err := api.NewTicketsService(client).Unfollow(context.Background(), ticketID); err != nil {
		return fmt.Errorf("failed to unfollow ticket: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "ticketId": "%s", "message": "Stopped following ticket"}`, ticketID)
	} else {
		fmt.Printf("Stopped following ticket %s\n", ticketID)
	}
	return nil
}