package api

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type AgentsService struct {
	client *Client
}

func NewAgentsService(client *Client) *AgentsService {
	return &AgentsService{client: client}
}

func (s *AgentsService) List(ctx context.Context, params map[string]string) ([]models.Agent, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/agents", query)
	if err != nil {
		return nil, err
	}

	var resp models.AgentListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *AgentsService) Get(ctx context.Context, agentID string) (*models.Agent, error) {
	data, err := s.client.Get(ctx, "/agents/"+agentID, nil)
	if err != nil {
		return nil, err
	}

	var modelsagent models.Agent
	if err := unmarshalData(data, &modelsagent); err != nil {
		return nil, err
	}
	return &modelsagent, nil
}

func (s *AgentsService) Search(ctx context.Context, email string, name string) ([]models.Agent, error) {
	agents, err := s.List(ctx, map[string]string{
		"limit": "100",
	})
	if err != nil {
		return nil, err
	}

	var filtered []models.Agent
	for _, a := range agents {
		if email != "" && a.Email != email {
			continue
		}
		if name != "" && a.Name != name {
			continue
		}
		filtered = append(filtered, a)
	}

	return filtered, nil
}