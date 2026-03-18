package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/spf13/cobra"
)

var attachmentsCmd = &cobra.Command{
	Use:   "attachments",
	Short: "Manage ticket attachments",
	Long:  `Upload, download, list, and delete attachments on tickets.`,
}

var attachmentsListCmd = &cobra.Command{
	Use:   "list <ticket-id>",
	Short: "List attachments for a ticket",
	Long:  `List all attachments attached to a specific ticket.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAttachmentsList,
}

var attachmentsUploadCmd = &cobra.Command{
	Use:   "upload <ticket-id> <file>",
	Short: "Upload an attachment to a ticket",
	Long:  `Upload a file as an attachment to a ticket.`,
	Args:  cobra.ExactArgs(2),
	RunE:  runAttachmentsUpload,
}

var attachmentsDownloadCmd = &cobra.Command{
	Use:   "download <attachment-id>",
	Short: "Download an attachment",
	Long:  `Download an attachment by ID.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAttachmentsDownload,
}

var attachmentsDeleteCmd = &cobra.Command{
	Use:   "delete <attachment-id>",
	Short: "Delete an attachment",
	Long:  `Delete an attachment by ID.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAttachmentsDelete,
}

var (
	flagOutputPath string
)

func init() {
	rootCmd.AddCommand(attachmentsCmd)
	attachmentsCmd.AddCommand(attachmentsListCmd)
	attachmentsCmd.AddCommand(attachmentsUploadCmd)
	attachmentsCmd.AddCommand(attachmentsDownloadCmd)
	attachmentsCmd.AddCommand(attachmentsDeleteCmd)

	attachmentsDownloadCmd.Flags().StringVarP(&flagOutputPath, "output", "o", "", "output file path (default: attachment name)")
}

func runAttachmentsList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]

	attachments, err := api.NewAttachmentsService(client).List(context.Background(), ticketID)
	if err != nil {
		return fmt.Errorf("failed to list attachments: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAttachments(attachments, out)
}

func runAttachmentsUpload(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	ticketID := args[0]
	filePath := args[1]

	attachment, err := api.NewAttachmentsService(client).Upload(context.Background(), ticketID, filePath)
	if err != nil {
		return fmt.Errorf("failed to upload attachment: %w", err)
	}

	out := getOutputFormat()
	return output.PrintAttachment(attachment, out)
}

func runAttachmentsDownload(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	attachmentID := args[0]

	outputPath := flagOutputPath
	if outputPath == "" {
		outputPath = attachmentID
	}

	if err := api.NewAttachmentsService(client).Download(context.Background(), attachmentID, outputPath); err != nil {
		return fmt.Errorf("failed to download attachment: %w", err)
	}

	fmt.Printf("Attachment downloaded to: %s\n", outputPath)
	return nil
}

func runAttachmentsDelete(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	attachmentID := args[0]

	if err := api.NewAttachmentsService(client).Delete(context.Background(), attachmentID); err != nil {
		return fmt.Errorf("failed to delete attachment: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "attachmentId": "%s", "message": "Attachment deleted successfully"}`, attachmentID)
	} else {
		fmt.Printf("Attachment %s deleted successfully\n", attachmentID)
	}
	return nil
}