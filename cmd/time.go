package cmd

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Manage time entries",
	Long:  `Manage time entries on tickets for tracking agent work time.`,
}

var timeListCmd = &cobra.Command{
	Use:   "list <ticket-id>",
	Short: "List time entries for a ticket",
	Long:  `List all time entries recorded for a specific ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTimeList,
}

var timeAddCmd = &cobra.Command{
	Use:   "add <ticket-id>",
	Short: "Add a time entry to a ticket",
	Long:  `Record time spent working on a ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTimeAdd,
}

var timeDeleteCmd = &cobra.Command{
	Use:   "delete <time-entry-id>",
	Short: "Delete a time entry",
	Long:  `Delete a specific time entry by ID.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTimeDelete,
}

var timeReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate time report",
	Long:  `Generate a time report showing hours logged across tickets and agents.`,
	RunE:  runTimeReport,
}

var (
	flagTimeDuration    string
	flagTimeExecutedAt  string
	flagTimeAgentID     string
	flagTimeFromDate    string
	flagTimeToDate      string
)

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.AddCommand(timeListCmd)
	timeCmd.AddCommand(timeAddCmd)
	timeCmd.AddCommand(timeDeleteCmd)
	timeCmd.AddCommand(timeReportCmd)

	timeAddCmd.Flags().StringVarP(&flagTimeDuration, "duration", "d", "", "time duration (e.g., 1h30m, 2h, 45m)")
	timeAddCmd.Flags().StringVarP(&flagDescription, "description", "D", "", "description of work performed")
	timeAddCmd.Flags().StringVarP(&flagTimeAgentID, "agent", "a", "", "agent ID (default: current user)")
	timeAddCmd.Flags().StringVarP(&flagTimeExecutedAt, "executed", "e", "", "when the work was performed (YYYY-MM-DD)")

	timeReportCmd.Flags().StringVarP(&flagTimeAgentID, "agent", "a", "", "filter by agent ID")
	timeReportCmd.Flags().StringVarP(&flagTimeFromDate, "from", "f", "", "from date (YYYY-MM-DD)")
	timeReportCmd.Flags().StringVarP(&flagTimeToDate, "to", "t", "", "to date (YYYY-MM-DD)")
}

func parseDuration(duration string) (hours, minutes int, err error) {
	duration = strings.ToLower(strings.TrimSpace(duration))
	
	hours = 0
	minutes = 0

	if strings.Contains(duration, "h") {
		parts := strings.SplitN(duration, "h", 2)
		h, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, fmt.Errorf("invalid duration format: %s", duration)
		}
		hours = h
		duration = parts[1]
	}

	if strings.Contains(duration, "m") {
		parts := strings.SplitN(duration, "m", 2)
		m, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, fmt.Errorf("invalid duration format: %s", duration)
		}
		minutes = m
	}

	if hours == 0 && minutes == 0 && duration != "" {
		return 0, 0, fmt.Errorf("invalid duration format: %s (use format like 1h30m, 2h, or 45m)", duration)
	}

	return hours, minutes, nil
}

func runTimeList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	entries, err := api.NewTimeService(client).List(context.Background(), ticketID)
	if err != nil {
		return fmt.Errorf("failed to list time entries: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTimeEntries(entries, out)
}

func runTimeAdd(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	if flagTimeDuration == "" {
		return fmt.Errorf("--duration is required (e.g., 1h30m, 2h, 45m)")
	}

	hours, minutes, err := parseDuration(flagTimeDuration)
	if err != nil {
		return err
	}

	req := models.TimeEntryCreateRequest{
		Hours:       hours,
		Minutes:     minutes,
		Description: flagDescription,
		AgentID:     flagTimeAgentID,
		ExecutedAt:  flagTimeExecutedAt,
	}

	entry, err := api.NewTimeService(client).Add(context.Background(), ticketID, req)
	if err != nil {
		return fmt.Errorf("failed to add time entry: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTimeEntry(entry, out)
}

func runTimeDelete(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	timeEntryID := args[0]

	if err := api.NewTimeService(client).Delete(context.Background(), timeEntryID); err != nil {
		return fmt.Errorf("failed to delete time entry: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "timeEntryId": "%s", "message": "Time entry deleted successfully"}`, timeEntryID)
	} else {
		fmt.Printf("Time entry %s deleted successfully\n", timeEntryID)
	}
	return nil
}

func runTimeReport(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := make(map[string]string)
	if flagTimeAgentID != "" {
		params["agentId"] = flagTimeAgentID
	}
	if flagTimeFromDate != "" {
		params["from"] = flagTimeFromDate
	}
	if flagTimeToDate != "" {
		params["to"] = flagTimeToDate
	}

	report, err := api.NewTimeService(client).GetReport(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to get time report: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTimeReport(report, out)
}