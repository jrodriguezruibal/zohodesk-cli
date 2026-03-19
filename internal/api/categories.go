package api

import (
	"context"
	"encoding/json"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type CategoriesService struct {
	client *Client
}

func NewCategoriesService(client *Client) *CategoriesService {
	return &CategoriesService{client: client}
}

func (s *CategoriesService) List(ctx context.Context) ([]models.Category, error) {
	data, err := s.client.Get(ctx, "/kb/categories", nil)
	if err != nil {
		return nil, err
	}

	var resp models.CategoryListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *CategoriesService) Get(ctx context.Context, categoryID string) (*models.Category, error) {
	data, err := s.client.Get(ctx, "/kb/categories/"+categoryID, nil)
	if err != nil {
		return nil, err
	}

	var modelscategory models.Category
	if err := unmarshalData(data, &modelscategory); err != nil {
		return nil, err
	}
	return &modelscategory, nil
}