package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
	"github.com/spf13/cobra"
)

var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "Manage knowledge base articles",
	Long:  `List, search, create, update, and delete knowledge base articles.`,
}

var articlesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List articles",
	Long:  `List all knowledge base articles.`,
	RunE:  runArticlesList,
}

var articlesGetCmd = &cobra.Command{
	Use:   "get <article-id>",
	Short: "Get article details",
	Long:  `Get detailed information about a specific article.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runArticlesGet,
}

var articlesSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search articles",
	Long:  `Search knowledge base articles by keyword.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runArticlesSearch,
}

var articlesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new article",
	Long:  `Create a new knowledge base article.`,
	RunE:  runArticlesCreate,
}

var articlesUpdateCmd = &cobra.Command{
	Use:   "update <article-id>",
	Short: "Update an article",
	Long:  `Update an existing knowledge base article.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runArticlesUpdate,
}

var articlesDeleteCmd = &cobra.Command{
	Use:   "delete <article-id>",
	Short: "Delete an article",
	Long:  `Delete a knowledge base article.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runArticlesDelete,
}

var (
	flagArticleTitle      string
	flagArticleContent    string
	flagArticleSummary    string
	flagArticleCategory   string
	flagArticleStatus     string
	flagArticleLimit      int
	flagArticleSearchQuery string
)

func init() {
	rootCmd.AddCommand(articlesCmd)
	articlesCmd.AddCommand(articlesListCmd)
	articlesCmd.AddCommand(articlesGetCmd)
	articlesCmd.AddCommand(articlesSearchCmd)
	articlesCmd.AddCommand(articlesCreateCmd)
	articlesCmd.AddCommand(articlesUpdateCmd)
	articlesCmd.AddCommand(articlesDeleteCmd)

	articlesListCmd.Flags().IntVarP(&flagArticleLimit, "limit", "l", 50, "maximum number of articles to return")

	articlesCreateCmd.Flags().StringVarP(&flagArticleTitle, "title", "t", "", "article title (required)")
	articlesCreateCmd.Flags().StringVarP(&flagArticleContent, "content", "c", "", "article content (required)")
	articlesCreateCmd.Flags().StringVarP(&flagArticleSummary, "summary", "s", "", "article summary")
	articlesCreateCmd.Flags().StringVarP(&flagArticleCategory, "category", "C", "", "category ID")

	articlesUpdateCmd.Flags().StringVarP(&flagArticleTitle, "title", "t", "", "article title")
	articlesUpdateCmd.Flags().StringVarP(&flagArticleContent, "content", "c", "", "article content")
	articlesUpdateCmd.Flags().StringVarP(&flagArticleSummary, "summary", "s", "", "article summary")
	articlesUpdateCmd.Flags().StringVarP(&flagArticleCategory, "category", "C", "", "category ID")
	articlesUpdateCmd.Flags().StringVarP(&flagArticleStatus, "status", "S", "", "article status")
}

func runArticlesList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	params := map[string]string{
		"limit": fmt.Sprintf("%d", flagArticleLimit),
	}

	articles, err := api.NewArticlesService(client).List(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to list articles: %w", err)
	}

	out := getOutputFormat()
	return output.PrintArticles(articles, out)
}

func runArticlesGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	articleID := args[0]

	article, err := api.NewArticlesService(client).Get(context.Background(), articleID)
	if err != nil {
		return fmt.Errorf("failed to get article: %w", err)
	}

	out := getOutputFormat()
	return output.PrintArticle(article, out)
}

func runArticlesSearch(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	query := args[0]

	articles, err := api.NewArticlesService(client).Search(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to search articles: %w", err)
	}

	out := getOutputFormat()
	return output.PrintArticles(articles, out)
}

func runArticlesCreate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	if flagArticleTitle == "" {
		return fmt.Errorf("--title is required")
	}
	if flagArticleContent == "" {
		return fmt.Errorf("--content is required")
	}

	req := models.ArticleCreateRequest{
		Title:      flagArticleTitle,
		Content:    flagArticleContent,
		Summary:    flagArticleSummary,
		CategoryID: flagArticleCategory,
	}

	article, err := api.NewArticlesService(client).Create(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}

	out := getOutputFormat()
	return output.PrintArticle(article, out)
}

func runArticlesUpdate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	articleID := args[0]

	req := models.ArticleUpdateRequest{
		Title:      flagArticleTitle,
		Content:    flagArticleContent,
		Summary:    flagArticleSummary,
		CategoryID: flagArticleCategory,
		Status:     flagArticleStatus,
	}

	article, err := api.NewArticlesService(client).Update(context.Background(), articleID, req)
	if err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}

	out := getOutputFormat()
	return output.PrintArticle(article, out)
}

func runArticlesDelete(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	articleID := args[0]

	if err := api.NewArticlesService(client).Delete(context.Background(), articleID); err != nil {
		return fmt.Errorf("failed to delete article: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "articleId": "%s", "message": "Article deleted successfully"}`, articleID)
	} else {
		fmt.Printf("Article %s deleted successfully\n", articleID)
	}
	return nil
}