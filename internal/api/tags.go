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

	var resp models.TagListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *TagsService) Add(ctx context.Context, ticketID string, tags []string) error {
	req := models.TicketTagsRequest{Tags: tags}
	_, err := s.client.Post(ctx, "/tickets/"+ticketID+"/tags", req)
	return err
}

func (s *TagsService) Remove(ctx context.Context, ticketID string, tag string) error {
	return s.client.Delete(ctx, "/tickets/"+ticketID+"/tags/"+tag)
}