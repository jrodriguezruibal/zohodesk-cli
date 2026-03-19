package api

import (
	"context"
	"encoding/json"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type AccountsService struct {
	client *Client
}

func NewAccountsService(client *Client) *AccountsService {
	return &AccountsService{client: client}
}

func (s *AccountsService) List(ctx context.Context) ([]models.Account, error) {
	data, err := s.client.Get(ctx, "/accounts", nil)
	if err != nil {
		return nil, err
	}

	var resp models.AccountListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *AccountsService) Get(ctx context.Context, accountID string) (*models.Account, error) {
	data, err := s.client.Get(ctx, "/accounts/"+accountID, nil)
	if err != nil {
		return nil, err
	}

	var modelsaccount models.Account
	if err := unmarshalData(data, &modelsaccount); err != nil {
		return nil, err
	}
	return &modelsaccount, nil
}

func (s *AccountsService) Create(ctx context.Context, req models.AccountCreateRequest) (*models.Account, error) {
	data, err := s.client.Post(ctx, "/accounts", req)
	if err != nil {
		return nil, err
	}

	var modelsaccount models.Account
	if err := unmarshalData(data, &modelsaccount); err != nil {
		return nil, err
	}
	return &modelsaccount, nil
}

func (s *AccountsService) Update(ctx context.Context, accountID string, req models.AccountUpdateRequest) (*models.Account, error) {
	data, err := s.client.Patch(ctx, "/accounts/"+accountID, req)
	if err != nil {
		return nil, err
	}

	var modelsaccount models.Account
	if err := unmarshalData(data, &modelsaccount); err != nil {
		return nil, err
	}
	return &modelsaccount, nil
}

func (s *AccountsService) Delete(ctx context.Context, accountID string) error {
	return s.client.Delete(ctx, "/accounts/"+accountID)
}