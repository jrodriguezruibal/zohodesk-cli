package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var departmentsCmd = &cobra.Command{
	Use:   "departments",
	Short: "Manage Zoho Desk departments",
	Long:  `Commands for listing and viewing departments in Zoho Desk.`,
}

var departmentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List departments",
	Long:  `List all departments in Zoho Desk.`,
	RunE:  runDepartmentsList,
}

var departmentsGetCmd = &cobra.Command{
	Use:   "get <department-id>",
	Short: "Get department details",
	Long:  `Get detailed information about a specific department.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runDepartmentsGet,
}

var (
	departmentFlagLimit int
)

func init() {
	rootCmd.AddCommand(departmentsCmd)
	departmentsCmd.AddCommand(departmentsListCmd)
	departmentsCmd.AddCommand(departmentsGetCmd)

	departmentsListCmd.Flags().IntVarP(&departmentFlagLimit, "limit", "l", 50, "maximum number of departments to return")
}

func runDepartmentsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := map[string]string{
		"limit": fmt.Sprintf("%d", departmentFlagLimit),
	}

	departments, err := api.NewDepartmentsService(client).List(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to list departments: %w", err)
	}

	out := getOutputFormat()
	return output.PrintDepartments(departments, out)
}

func runDepartmentsGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	departmentID := args[0]

	department, err := api.NewDepartmentsService(client).Get(context.Background(), departmentID)
	if err != nil {
		return fmt.Errorf("failed to get department: %w", err)
	}

	out := getOutputFormat()
	return output.PrintDepartment(department, out)
}