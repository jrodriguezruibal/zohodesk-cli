package api

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type DepartmentsService struct {
	client *Client
}

func NewDepartmentsService(client *Client) *DepartmentsService {
	return &DepartmentsService{client: client}
}

func (s *DepartmentsService) List(ctx context.Context, params map[string]string) ([]models.Department, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/departments", query)
	if err != nil {
		return nil, err
	}

	var resp models.DepartmentListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *DepartmentsService) Get(ctx context.Context, departmentID string) (*models.Department, error) {
	data, err := s.client.Get(ctx, "/departments/"+departmentID, nil)
	if err != nil {
		return nil, err
	}

	var modelsdepartment models.Department
	if err := unmarshalData(data, &modelsdepartment); err != nil {
		return nil, err
	}
	return &modelsdepartment, nil
}