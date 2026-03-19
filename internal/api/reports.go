package api

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type ReportsService struct {
	client *Client
}

func NewReportsService(client *Client) *ReportsService {
	return &ReportsService{client: client}
}

func (s *ReportsService) GetTicketStats(ctx context.Context, params map[string]string) (*models.TicketStats, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/reports/tickets", query)
	if err != nil {
		return nil, err
	}

	var stats models.TicketStats
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

func (s *ReportsService) GetAgentStats(ctx context.Context, params map[string]string) (*models.AgentStatsList, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/reports/agents", query)
	if err != nil {
		return nil, err
	}

	var resp models.AgentStatsList
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ReportsService) GetSLAStats(ctx context.Context, params map[string]string) (*models.SLAStats, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/reports/sla", query)
	if err != nil {
		return nil, err
	}

	var stats models.SLAStats
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}