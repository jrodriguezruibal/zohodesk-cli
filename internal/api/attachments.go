package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type AttachmentsService struct {
	client *Client
}

func NewAttachmentsService(client *Client) *AttachmentsService {
	return &AttachmentsService{client: client}
}

func (s *AttachmentsService) List(ctx context.Context, ticketID string) ([]models.Attachment, error) {
	data, err := s.client.Get(ctx, "/tickets/"+ticketID+"/attachments", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data []models.Attachment `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *AttachmentsService) Upload(ctx context.Context, ticketID string, filePath string) (*models.Attachment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	resp, err := s.client.doMultipartRequest(ctx, "/tickets/"+ticketID+"/attachments", &body, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data models.Attachment `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result.Data, nil
}

func (s *AttachmentsService) Download(ctx context.Context, attachmentID string, outputPath string) error {
	resp, err := s.client.doRequest(ctx, http.MethodGet, "/attachments/"+attachmentID+"/download", nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (s *AttachmentsService) Delete(ctx context.Context, attachmentID string) error {
	return s.client.Delete(ctx, "/attachments/"+attachmentID)
}