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

type CommentsService struct {
	client *Client
}

func NewCommentsService(client *Client) *CommentsService {
	return &CommentsService{client: client}
}

func (s *CommentsService) List(ctx context.Context, ticketID string) ([]models.Comment, error) {
	data, err := s.client.Get(ctx, "/tickets/"+ticketID+"/comments", nil)
	if err != nil {
		return nil, err
	}

	var resp models.CommentListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *CommentsService) Get(ctx context.Context, ticketID string, commentID string) (*models.Comment, error) {
	data, err := s.client.Get(ctx, "/tickets/"+ticketID+"/comments/"+commentID, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Comment `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *CommentsService) Create(ctx context.Context, ticketID string, req models.CommentCreateRequest) (*models.Comment, error) {
	data, err := s.client.Post(ctx, "/tickets/"+ticketID+"/comments", req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Comment `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *CommentsService) Reply(ctx context.Context, ticketID string, content string, isPublic bool) (*models.Comment, error) {
	req := models.CommentCreateRequest{
		Content:  content,
		IsPublic: isPublic,
	}
	return s.Create(ctx, ticketID, req)
}

func (s *CommentsService) AddNote(ctx context.Context, ticketID string, content string) (*models.Comment, error) {
	req := models.CommentCreateRequest{
		Content:  content,
		IsPublic: false,
	}
	return s.Create(ctx, ticketID, req)
}

func (s *CommentsService) Delete(ctx context.Context, ticketID string, commentID string) error {
	return s.client.Delete(ctx, "/tickets/"+ticketID+"/comments/"+commentID)
}

func (s *CommentsService) UploadAttachment(ctx context.Context, ticketID string, filePath string) (*models.AttachmentUploadResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	if err := s.client.ensureToken(ctx); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/tickets/%s/attachments", ticketID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Zoho-oauthtoken "+s.client.tokenCache.AccessToken)
	req.Header.Set("orgId", s.client.config.OrgID)

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("upload failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	var result models.AttachmentUploadResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (s *CommentsService) DownloadAttachment(ctx context.Context, ticketID string, attachmentID string, outputPath string) error {
	endpoint := fmt.Sprintf("/tickets/%s/attachments/%s", ticketID, attachmentID)
	
	resp, err := s.client.GetRaw(ctx, endpoint)
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