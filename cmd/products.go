package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Manage products",
	Long:  `List and get information about products.`,
}

var productsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List products",
	Long:  `List all products.`,
	RunE:  runProductsList,
}

var productsGetCmd = &cobra.Command{
	Use:   "get <product-id>",
	Short: "Get product details",
	Long:  `Get detailed information about a specific product.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runProductsGet,
}

func init() {
	rootCmd.AddCommand(productsCmd)
	productsCmd.AddCommand(productsListCmd)
	productsCmd.AddCommand(productsGetCmd)
}

func runProductsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	products, err := api.NewProductsService(client).List(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list products: %w", err)
	}

	out := getOutputFormat()
	return output.PrintProducts(products, out)
}

func runProductsGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	productID := args[0]

	product, err := api.NewProductsService(client).Get(context.Background(), productID)
	if err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	out := getOutputFormat()
	return output.PrintProduct(product, out)
}