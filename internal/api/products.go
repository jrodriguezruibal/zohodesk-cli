package api

import (
	"context"
	"encoding/json"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type ProductsService struct {
	client *Client
}

func NewProductsService(client *Client) *ProductsService {
	return &ProductsService{client: client}
}

func (s *ProductsService) List(ctx context.Context) ([]models.Product, error) {
	data, err := s.client.Get(ctx, "/products", nil)
	if err != nil {
		return nil, err
	}

	var resp models.ProductListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *ProductsService) Get(ctx context.Context, productID string) (*models.Product, error) {
	data, err := s.client.Get(ctx, "/products/"+productID, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Product `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}