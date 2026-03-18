package api

import (
	"context"
	"encoding/json"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type TimeService struct {
	client *Client
}

func NewTimeService(client *Client) *TimeService {
	return &TimeService{client: client}
}

func (s *TimeService) List(ctx context.Context, ticketID string) ([]models.TimeEntry, error) {
	data, err := s.client.Get(ctx, "/tickets/"+ticketID+"/timeentries", nil)
	if err != nil {
		return nil, err
	}

	var resp models.TimeEntryListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *TimeService) Add(ctx context.Context, ticketID string, req models.TimeEntryCreateRequest) (*models.TimeEntry, error) {
	data, err := s.client.Post(ctx, "/tickets/"+ticketID+"/timeentries", req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.TimeEntry `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *TimeService) Delete(ctx context.Context, timeEntryID string) error {
	return s.client.Delete(ctx, "/timeentries/"+timeEntryID)
}

func (s *TimeService) GetReport(ctx context.Context, params map[string]string) (*models.TimeReport, error) {
	data, err := s.client.Get(ctx, "/reports/time", nil)
	if err != nil {
		return nil, err
	}

	var report models.TimeReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, err
	}

	return &report, nil
}