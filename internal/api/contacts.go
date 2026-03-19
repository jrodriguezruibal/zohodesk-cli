package api

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type ContactsService struct {
	client *Client
}

func NewContactsService(client *Client) *ContactsService {
	return &ContactsService{client: client}
}

func (s *ContactsService) List(ctx context.Context, params map[string]string) ([]models.Contact, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/contacts", query)
	if err != nil {
		return nil, err
	}

	var resp models.ContactListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *ContactsService) Get(ctx context.Context, contactID string) (*models.Contact, error) {
	data, err := s.client.Get(ctx, "/contacts/"+contactID, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Contact `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *ContactsService) Search(ctx context.Context, email string, name string) ([]models.Contact, error) {
	contacts, err := s.List(ctx, map[string]string{
		"limit": "100",
	})
	if err != nil {
		return nil, err
	}

	var filtered []models.Contact
	for _, c := range contacts {
		if email != "" && !strings.Contains(strings.ToLower(c.Email), strings.ToLower(email)) {
			continue
		}
		if name != "" {
			fullName := strings.ToLower(c.FirstName + " " + c.LastName)
			if !strings.Contains(fullName, strings.ToLower(name)) {
				continue
			}
		}
		filtered = append(filtered, c)
	}

	return filtered, nil
}

func (s *ContactsService) SearchByEmail(ctx context.Context, email string) (*models.Contact, error) {
	contacts, err := s.Search(ctx, email, "")
	if err != nil {
		return nil, err
	}

	if len(contacts) == 0 {
		return nil, nil
	}

	return &contacts[0], nil
}