package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var reportsCmd = &cobra.Command{
	Use:   "reports",
	Short: "Generate reports and statistics",
	Long:  `Generate various reports and statistics for your helpdesk.`,
}

var reportsTicketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "Ticket statistics",
	Long:  `Get ticket statistics by status, priority, and department.`,
	RunE:  runReportsTickets,
}

var reportsAgentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "Agent statistics",
	Long:  `Get agent performance statistics.`,
	RunE:  runReportsAgents,
}

var reportsSLACmd = &cobra.Command{
	Use:   "sla",
	Short: "SLA statistics",
	Long:  `Get SLA compliance statistics.`,
	RunE:  runReportsSLA,
}

var (
	flagReportFrom string
	flagReportTo    string
	flagReportDept  string
	flagReportAgent string
)

func init() {
	rootCmd.AddCommand(reportsCmd)
	reportsCmd.AddCommand(reportsTicketsCmd)
	reportsCmd.AddCommand(reportsAgentsCmd)
	reportsCmd.AddCommand(reportsSLACmd)

	reportsTicketsCmd.Flags().StringVarP(&flagReportFrom, "from", "f", "", "from date (YYYY-MM-DD)")
	reportsTicketsCmd.Flags().StringVarP(&flagReportTo, "to", "t", "", "to date (YYYY-MM-DD)")
	reportsTicketsCmd.Flags().StringVarP(&flagReportDept, "department", "d", "", "filter by department ID")

	reportsAgentsCmd.Flags().StringVarP(&flagReportFrom, "from", "f", "", "from date (YYYY-MM-DD)")
	reportsAgentsCmd.Flags().StringVarP(&flagReportTo, "to", "t", "", "to date (YYYY-MM-DD)")
	reportsAgentsCmd.Flags().StringVarP(&flagReportAgent, "agent", "a", "", "filter by agent ID")

	reportsSLACmd.Flags().StringVarP(&flagReportFrom, "from", "f", "", "from date (YYYY-MM-DD)")
	reportsSLACmd.Flags().StringVarP(&flagReportTo, "to", "t", "", "to date (YYYY-MM-DD)")
	reportsSLACmd.Flags().StringVarP(&flagReportDept, "department", "d", "", "filter by department ID")
}

func runReportsTickets(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := make(map[string]string)
	if flagReportFrom != "" {
		params["from"] = flagReportFrom
	}
	if flagReportTo != "" {
		params["to"] = flagReportTo
	}
	if flagReportDept != "" {
		params["departmentId"] = flagReportDept
	}

	stats, err := api.NewReportsService(client).GetTicketStats(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to get ticket stats: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTicketStats(stats, out)
}

func runReportsAgents(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := make(map[string]string)
	if flagReportFrom != "" {
		params["from"] = flagReportFrom
	}
	if flagReportTo != "" {
		params["to"] = flagReportTo
	}
	if flagReportAgent != "" {
		params["agentId"] = flagReportAgent
	}

	stats, err := api.NewReportsService(client).GetAgentStats(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to get agent stats: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAgentStats(stats.Data, out)
}

func runReportsSLA(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := make(map[string]string)
	if flagReportFrom != "" {
		params["from"] = flagReportFrom
	}
	if flagReportTo != "" {
		params["to"] = flagReportTo
	}
	if flagReportDept != "" {
		params["departmentId"] = flagReportDept
	}

	stats, err := api.NewReportsService(client).GetSLAStats(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to get SLA stats: %w", err)
	}

	out := getOutputFormat()
	return output.PrintSLAStats(stats, out)
}