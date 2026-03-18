package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/batch"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "Manage ticket comments",
	Long:  `Manage comments and responses on Zoho Desk tickets.`,
}

var commentsListCmd = &cobra.Command{
	Use:   "list <ticket-id>",
	Short: "List comments on a ticket",
	Long:  `List all comments and responses on a specific ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCommentsList,
}

var commentsReplyCmd = &cobra.Command{
	Use:   "reply <ticket-id>",
	Short: "Add a public reply to a ticket",
	Long:  `Add a public reply to a ticket that the customer can see.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCommentsReply,
}

var commentsNoteCmd = &cobra.Command{
	Use:   "note <ticket-id>",
	Short: "Add a private note to a ticket",
	Long:  `Add a private note to a ticket that only agents can see.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCommentsNote,
}

var (
	commentMessage string
	commentPublic  bool
	batchFlag      bool
	batchFile      string
)

func init() {
	rootCmd.AddCommand(commentsCmd)
	commentsCmd.AddCommand(commentsListCmd)
	commentsCmd.AddCommand(commentsReplyCmd)
	commentsCmd.AddCommand(commentsNoteCmd)

	commentsListCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "output format (json, yaml, table)")

	commentsReplyCmd.Flags().StringVarP(&commentMessage, "message", "m", "", "reply message (required)")
	commentsReplyCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "output format")
	commentsReplyCmd.MarkFlagRequired("message")

	commentsNoteCmd.Flags().StringVarP(&commentMessage, "message", "m", "", "note message (required)")
	commentsNoteCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "output format")
	commentsNoteCmd.MarkFlagRequired("message")

	// Add batch flags to tickets commands
	ticketsCreateCmd.Flags().BoolVarP(&batchFlag, "batch", "b", false, "read batch input from stdin (JSON array)")
	ticketsCreateCmd.Flags().StringVarP(&batchFile, "file", "f", "", "read batch input from file (JSON array)")

	ticketsUpdateCmd.Flags().BoolVarP(&batchFlag, "batch", "b", false, "read batch input from stdin (JSON array)")
	ticketsUpdateCmd.Flags().StringVarP(&batchFile, "file", "f", "", "read batch input from file (JSON array)")

	ticketsCloseCmd.Flags().BoolVarP(&batchFlag, "batch", "b", false, "read batch input from stdin (JSON array of ticket IDs)")
	ticketsCloseCmd.Flags().StringVarP(&batchFile, "file", "f", "", "read batch input from file (JSON array of ticket IDs)")
}

func runCommentsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	comments, err := api.NewCommentsService(client).List(context.Background(), ticketID)
	if err != nil {
		return fmt.Errorf("failed to list comments: %w", err)
	}

	out := getOutputFormat()
	return output.PrintComments(comments, out)
}

func runCommentsReply(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	comment, err := api.NewCommentsService(client).Reply(context.Background(), ticketID, commentMessage, true)
	if err != nil {
		return fmt.Errorf("failed to add reply: %w", err)
	}

	out := getOutputFormat()
	return output.PrintComment(comment, out)
}

func runCommentsNote(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	comment, err := api.NewCommentsService(client).AddNote(context.Background(), ticketID, commentMessage)
	if err != nil {
		return fmt.Errorf("failed to add note: %w", err)
	}

	out := getOutputFormat()
	return output.PrintComment(comment, out)
}

func getBatchReader() (io.Reader, error) {
	if batchFile != "" {
		file, err := os.Open(batchFile)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		return file, nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, fmt.Errorf("no data on stdin. Pipe JSON data or use --file option")
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("failed to read stdin: %w", err)
	}

	return bytes.NewReader(data), nil
}

func runBatchCreate() error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	reader, err := getBatchReader()
	if err != nil {
		return err
	}
	if closer, ok := reader.(io.Closer); ok {
		defer closer.Close()
	}

	processor := batch.NewProcessor()
	processor.SetRateLimitDelay(100 * time.Millisecond)

	ticketsService := api.NewTicketsService(client)
	result := processor.ProcessTicketsCreate(context.Background(), reader, ticketsService.Create)

	out := getOutputFormat()
	if out == "json" {
		return output.PrintJSON(result)
	}

	fmt.Printf("Batch operation completed:\n")
	fmt.Printf("  Total: %d\n", result.Total)
	fmt.Printf("  Succeeded: %d\n", result.Succeeded)
	fmt.Printf("  Failed: %d\n", result.Failed)

	if result.Failed > 0 {
		fmt.Println("\nFailed items:")
		for _, r := range result.Results {
			if r.Status == "failed" {
				fmt.Printf("  Index %d: %s\n", r.Index, r.Error)
			}
		}
	}

	return nil
}

func runBatchUpdate() error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	reader, err := getBatchReader()
	if err != nil {
		return err
	}
	if closer, ok := reader.(io.Closer); ok {
		defer closer.Close()
	}

	processor := batch.NewProcessor()
	processor.SetRateLimitDelay(100 * time.Millisecond)

	ticketsService := api.NewTicketsService(client)
	result := processor.ProcessTicketsUpdate(context.Background(), reader, ticketsService.Update)

	out := getOutputFormat()
	if out == "json" {
		return output.PrintJSON(result)
	}

	fmt.Printf("Batch operation completed:\n")
	fmt.Printf("  Total: %d\n", result.Total)
	fmt.Printf("  Succeeded: %d\n", result.Succeeded)
	fmt.Printf("  Failed: %d\n", result.Failed)

	if result.Failed > 0 {
		fmt.Println("\nFailed items:")
		for _, r := range result.Results {
			if r.Status == "failed" {
				fmt.Printf("  Index %d: %s\n", r.Index, r.Error)
			}
		}
	}

	return nil
}

func runBatchClose() error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	reader, err := getBatchReader()
	if err != nil {
		return err
	}
	if closer, ok := reader.(io.Closer); ok {
		defer closer.Close()
	}

	processor := batch.NewProcessor()
	processor.SetRateLimitDelay(100 * time.Millisecond)

	ticketsService := api.NewTicketsService(client)
	result := processor.ProcessTicketsClose(context.Background(), reader, func(ctx context.Context, ticketID string) error {
		return ticketsService.Close(ctx, ticketID, "")
	})

	out := getOutputFormat()
	if out == "json" {
		return output.PrintJSON(result)
	}

	fmt.Printf("Batch operation completed:\n")
	fmt.Printf("  Total: %d\n", result.Total)
	fmt.Printf("  Succeeded: %d\n", result.Succeeded)
	fmt.Printf("  Failed: %d\n", result.Failed)

	if result.Failed > 0 {
		fmt.Println("\nFailed items:")
		for _, r := range result.Results {
			if r.Status == "failed" {
				fmt.Printf("  Index %d: %s\n", r.Index, r.Error)
			}
		}
	}

	return nil
}