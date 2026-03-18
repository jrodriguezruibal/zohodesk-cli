package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Manage knowledge base categories",
	Long:  `List and get knowledge base categories.`,
}

var categoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List categories",
	Long:  `List all knowledge base categories.`,
	RunE:  runCategoriesList,
}

var categoriesGetCmd = &cobra.Command{
	Use:   "get <category-id>",
	Short: "Get category details",
	Long:  `Get detailed information about a specific category.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCategoriesGet,
}

func init() {
	rootCmd.AddCommand(categoriesCmd)
	categoriesCmd.AddCommand(categoriesListCmd)
	categoriesCmd.AddCommand(categoriesGetCmd)
}

func runCategoriesList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	categories, err := api.NewCategoriesService(client).List(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list categories: %w", err)
	}

	out := getOutputFormat()
	return output.PrintCategories(categories, out)
}

func runCategoriesGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	categoryID := args[0]

	category, err := api.NewCategoriesService(client).Get(context.Background(), categoryID)
	if err != nil {
		return fmt.Errorf("failed to get category: %w", err)
	}

	out := getOutputFormat()
	return output.PrintCategory(category, out)
}