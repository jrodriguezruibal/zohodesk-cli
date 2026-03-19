package models

import (
	"encoding/json"
	"testing"
)

func TestProductJSON(t *testing.T) {
	product := Product{
		ID:          "123",
		Name:        "Enterprise Plan",
		Description: "Full-featured plan",
		ProductCode: "ENT",
		Category:    "Subscription",
		IsActive:    true,
		CreatedTime: "2024-01-15T10:30:00Z",
	}

	jsonData, err := json.Marshal(product)
	if err != nil {
		t.Errorf("Failed to marshal Product: %v", err)
	}

	var unmarshaled Product
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Product: %v", err)
	}

	if unmarshaled.Name != product.Name {
		t.Errorf("Expected Name %s, got %s", product.Name, unmarshaled.Name)
	}
	if unmarshaled.IsActive != product.IsActive {
		t.Errorf("Expected IsActive %v, got %v", product.IsActive, unmarshaled.IsActive)
	}
}

func TestProductListResponseJSON(t *testing.T) {
	response := ProductListResponse{
		Data: []Product{
			{ID: "1", Name: "Product 1", IsActive: true},
			{ID: "2", Name: "Product 2", IsActive: false},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal ProductListResponse: %v", err)
	}

	var unmarshaled ProductListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal ProductListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d products, got %d", len(response.Data), len(unmarshaled.Data))
	}
}

func TestAccountJSON(t *testing.T) {
	account := Account{
		ID:         "123",
		Name:       "Acme Corp",
		Email:      "contact@acme.com",
		Phone:      "+1234567890",
		Website:    "https://acme.com",
		Type:       "Customer",
		Industry:   "Technology",
		OwnerID:    "456",
		OwnerName:  "John Doe",
		IsActive:   true,
		CreatedTime: "2024-01-15T10:30:00Z",
	}

	jsonData, err := json.Marshal(account)
	if err != nil {
		t.Errorf("Failed to marshal Account: %v", err)
	}

	var unmarshaled Account
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Account: %v", err)
	}

	if unmarshaled.Name != account.Name {
		t.Errorf("Expected Name %s, got %s", account.Name, unmarshaled.Name)
	}
	if unmarshaled.IsActive != account.IsActive {
		t.Errorf("Expected IsActive %v, got %v", account.IsActive, unmarshaled.IsActive)
	}
}

func TestAccountCreateRequestJSON(t *testing.T) {
	req := AccountCreateRequest{
		Name:     "New Account",
		Email:    "info@example.com",
		Phone:    "+1234567890",
		Website:  "https://example.com",
		Type:     "Prospect",
		Industry: "Finance",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal AccountCreateRequest: %v", err)
	}

	var unmarshaled AccountCreateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal AccountCreateRequest: %v", err)
	}

	if unmarshaled.Name != req.Name {
		t.Errorf("Expected Name %s, got %s", req.Name, unmarshaled.Name)
	}
}

func TestAccountUpdateRequestJSON(t *testing.T) {
	req := AccountUpdateRequest{
		Name:   "Updated Account",
		Type:   "Customer",
		IsActive: true,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal AccountUpdateRequest: %v", err)
	}

	var unmarshaled AccountUpdateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal AccountUpdateRequest: %v", err)
	}

	if unmarshaled.Name != req.Name {
		t.Errorf("Expected Name %s, got %s", req.Name, unmarshaled.Name)
	}
}

func TestAccountListResponseJSON(t *testing.T) {
	response := AccountListResponse{
		Data: []Account{
			{ID: "1", Name: "Account 1", Type: "Customer"},
			{ID: "2", Name: "Account 2", Type: "Prospect"},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal AccountListResponse: %v", err)
	}

	var unmarshaled AccountListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal AccountListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d accounts, got %d", len(response.Data), len(unmarshaled.Data))
	}
}