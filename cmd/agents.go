package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var agentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "Manage Zoho Desk agents",
	Long:  `Commands for listing and viewing agents (team members) in Zoho Desk.`,
}

var agentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List agents",
	Long:  `List all agents (team members) in Zoho Desk.`,
	RunE:  runAgentsList,
}

var agentsGetCmd = &cobra.Command{
	Use:   "get <agent-id>",
	Short: "Get agent details",
	Long:  `Get detailed information about a specific agent.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAgentsGet,
}

var agentsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search agents",
	Long:  `Search agents by email or name.`,
	RunE:  runAgentsSearch,
}

var (
	agentFlagLimit  int
	agentFlagEmail  string
	agentFlagName   string
)

func init() {
	rootCmd.AddCommand(agentsCmd)
	agentsCmd.AddCommand(agentsListCmd)
	agentsCmd.AddCommand(agentsGetCmd)
	agentsCmd.AddCommand(agentsSearchCmd)

	agentsListCmd.Flags().IntVarP(&agentFlagLimit, "limit", "l", 50, "maximum number of agents to return")

	agentsSearchCmd.Flags().StringVarP(&agentFlagEmail, "email", "e", "", "search by email")
	agentsSearchCmd.Flags().StringVarP(&agentFlagName, "name", "n", "", "search by name")
}

func runAgentsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := map[string]string{
		"limit": fmt.Sprintf("%d", agentFlagLimit),
	}

	agents, err := api.NewAgentsService(client).List(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to list agents: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAgents(agents, out)
}

func runAgentsGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	agentID := args[0]

	agent, err := api.NewAgentsService(client).Get(context.Background(), agentID)
	if err != nil {
		return fmt.Errorf("failed to get agent: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAgent(agent, out)
}

func runAgentsSearch(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	agents, err := api.NewAgentsService(client).Search(context.Background(), agentFlagEmail, agentFlagName)
	if err != nil {
		return fmt.Errorf("failed to search agents: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAgents(agents, out)
}