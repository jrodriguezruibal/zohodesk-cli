package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Manage ticket tags",
	Long:  `Add, remove, and list tags on tickets.`,
}

var tagsListCmd = &cobra.Command{
	Use:   "list <ticket-id>",
	Short: "List tags on a ticket",
	Long:  `List all tags associated with a ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTagsList,
}

var tagsAddCmd = &cobra.Command{
	Use:   "add <ticket-id>",
	Short: "Add tags to a ticket",
	Long:  `Add one or more tags to a ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTagsAdd,
}

var tagsRemoveCmd = &cobra.Command{
	Use:   "remove <ticket-id> <tag>",
	Short: "Remove a tag from a ticket",
	Long:  `Remove a specific tag from a ticket.`,
	Args:  cobra.ExactArgs(2),
	RunE:  runTagsRemove,
}

var (
	flagTags []string
)

func init() {
	rootCmd.AddCommand(tagsCmd)
	tagsCmd.AddCommand(tagsListCmd)
	tagsCmd.AddCommand(tagsAddCmd)
	tagsCmd.AddCommand(tagsRemoveCmd)

	tagsAddCmd.Flags().StringSliceVarP(&flagTags, "tag", "t", []string{}, "tag to add (can be specified multiple times)")
}

func runTagsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	tags, err := api.NewTagsService(client).List(context.Background(), ticketID)
	if err != nil {
		return fmt.Errorf("failed to list tags: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTags(tags, out)
}

func runTagsAdd(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	if len(flagTags) == 0 {
		return fmt.Errorf("at least one --tag is required")
	}

	if err := api.NewTagsService(client).Add(context.Background(), ticketID, flagTags); err != nil {
		return fmt.Errorf("failed to add tags: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "ticketId": "%s", "tags": %v, "message": "Tags added successfully"}`, ticketID, flagTags)
	} else {
		fmt.Printf("Tags %v added to ticket %s successfully\n", flagTags, ticketID)
	}
	return nil
}

func runTagsRemove(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]
	tag := args[1]

	if err := api.NewTagsService(client).Remove(context.Background(), ticketID, tag); err != nil {
		return fmt.Errorf("failed to remove tag: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "ticketId": "%s", "tag": "%s", "message": "Tag removed successfully"}`, ticketID, tag)
	} else {
		fmt.Printf("Tag '%s' removed from ticket %s successfully\n", tag, ticketID)
	}
	return nil
}