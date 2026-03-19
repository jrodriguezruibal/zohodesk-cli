package api

import (
	"context"
	"encoding/json"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type TagsService struct {
	client *Client
}

func NewTagsService(client *Client) *TagsService {
	return &TagsService{client: client}
}

func (s *TagsService) List(ctx context.Context, ticketID string) ([]models.Tag, error) {
	data, err := s.client.Get(ctx, "/tickets/"+ticketID+"/tags", nil)
	if err != nil {
		return nil, err
	}

	// Try { "tags": [...] } format
	var resp struct {
		Tags []models.Tag `json:"tags"`
	}
	if err := json.Unmarshal(data, &resp); err == nil && len(resp.Tags) > 0 {
		return resp.Tags, nil
	}

	// Try { "data": [...] } format
	var resp2 models.TagListResponse
	if err := json.Unmarshal(data, &resp2); err == nil && len(resp2.Data) > 0 {
		return resp2.Data, nil
	}

	// Try direct array
	var tags []models.Tag
	if err := json.Unmarshal(data, &tags); err == nil {
		return tags, nil
	}

	return []models.Tag{}, nil
}

func (s *TagsService) Add(ctx context.Context, ticketID string, tags []string) error {
	req := models.TicketTagsRequest{Tags: tags}
	_, err := s.client.Post(ctx, "/tickets/"+ticketID+"/tags", req)
	return err
}

func (s *TagsService) Remove(ctx context.Context, ticketID string, tag string) error {
	return s.client.Delete(ctx, "/tickets/"+ticketID+"/tags/"+tag)
}