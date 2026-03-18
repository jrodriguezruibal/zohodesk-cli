package api

import (
	"context"
	"encoding/json"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type TasksService struct {
	client *Client
}

func NewTasksService(client *Client) *TasksService {
	return &TasksService{client: client}
}

func (s *TasksService) List(ctx context.Context, ticketID string) ([]models.Task, error) {
	endpoint := "/tasks"
	if ticketID != "" {
		endpoint = "/tickets/" + ticketID + "/tasks"
	}

	data, err := s.client.Get(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var resp models.TaskListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *TasksService) Get(ctx context.Context, taskID string) (*models.Task, error) {
	data, err := s.client.Get(ctx, "/tasks/"+taskID, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Task `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *TasksService) Create(ctx context.Context, req models.TaskCreateRequest) (*models.Task, error) {
	data, err := s.client.Post(ctx, "/tasks", req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Task `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *TasksService) Update(ctx context.Context, taskID string, req models.TaskUpdateRequest) (*models.Task, error) {
	data, err := s.client.Patch(ctx, "/tasks/"+taskID, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Task `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *TasksService) Complete(ctx context.Context, taskID string) error {
	req := models.TaskUpdateRequest{
		Status: "Completed",
	}
	_, err := s.Update(ctx, taskID, req)
	return err
}

func (s *TasksService) Delete(ctx context.Context, taskID string) error {
	return s.client.Delete(ctx, "/tasks/"+taskID)
}