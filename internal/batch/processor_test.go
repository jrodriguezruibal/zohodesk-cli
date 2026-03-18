package batch

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

func TestProcessTicketsCreate(t *testing.T) {
	processor := NewProcessor()
	processor.SetRateLimitDelay(0) // Disable delay for tests

	tickets := []models.TicketCreateRequest{
		{Subject: "Test Ticket 1", Email: "user1@example.com", Priority: "High"},
		{Subject: "Test Ticket 2", Email: "user2@example.com", Priority: "Medium"},
	}

	jsonData, _ := json.Marshal(tickets)
	reader := bytes.NewReader(jsonData)

	var createdCount int
	mockCreate := func(ctx context.Context, req models.TicketCreateRequest) (*models.Ticket, error) {
		createdCount++
		return &models.Ticket{
			ID:     "123" + string(rune(createdCount)),
			Subject: req.Subject,
			Status: "Open",
		}, nil
	}

	result := processor.ProcessTicketsCreate(context.Background(), reader, mockCreate)

	if result.Total != 2 {
		t.Errorf("Expected total 2, got %d", result.Total)
	}
	if result.Succeeded != 2 {
		t.Errorf("Expected succeeded 2, got %d", result.Succeeded)
	}
	if result.Failed != 0 {
		t.Errorf("Expected failed 0, got %d", result.Failed)
	}
	if len(result.Results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(result.Results))
	}
}

func TestProcessTicketsCreateInvalidJSON(t *testing.T) {
	processor := NewProcessor()

	invalidJSON := []byte(`not valid json`)
	reader := bytes.NewReader(invalidJSON)

	mockCreate := func(ctx context.Context, req models.TicketCreateRequest) (*models.Ticket, error) {
		return nil, nil
	}

	result := processor.ProcessTicketsCreate(context.Background(), reader, mockCreate)

	if result.Total != 0 {
		t.Errorf("Expected total 0 for invalid JSON, got %d", result.Total)
	}
	if len(result.Errors) == 0 {
		t.Error("Expected error for invalid JSON")
	}
}

func TestProcessTicketsUpdate(t *testing.T) {
	processor := NewProcessor()
	processor.SetRateLimitDelay(0)

	updates := []struct {
		ID        string `json:"id"`
		Status    string `json:"status,omitempty"`
		Priority  string `json:"priority,omitempty"`
	}{
		{ID: "123", Status: "Closed"},
		{ID: "456", Priority: "High"},
		{ID: "789", Status: "On Hold", Priority: "Medium"},
	}

	jsonData, _ := json.Marshal(updates)
	reader := bytes.NewReader(jsonData)

	mockUpdate := func(ctx context.Context, ticketID string, req models.TicketUpdateRequest) (*models.Ticket, error) {
		return &models.Ticket{
			ID:       ticketID,
			Status:   req.Status,
			Priority: req.Priority,
		}, nil
	}

	result := processor.ProcessTicketsUpdate(context.Background(), reader, mockUpdate)

	if result.Total != 3 {
		t.Errorf("Expected total 3, got %d", result.Total)
	}
	if result.Succeeded != 3 {
		t.Errorf("Expected succeeded 3, got %d", result.Succeeded)
	}
}

func TestProcessTicketsClose(t *testing.T) {
	processor := NewProcessor()
	processor.SetRateLimitDelay(0)

	ticketIDs := []string{"123", "456", "789"}

	jsonData, _ := json.Marshal(ticketIDs)
	reader := bytes.NewReader(jsonData)

	mockClose := func(ctx context.Context, ticketID string) error {
		return nil
	}

	result := processor.ProcessTicketsClose(context.Background(), reader, mockClose)

	if result.Total != 3 {
		t.Errorf("Expected total 3, got %d", result.Total)
	}
	if result.Succeeded != 3 {
		t.Errorf("Expected succeeded 3, got %d", result.Succeeded)
	}
	if result.Failed != 0 {
		t.Errorf("Expected failed 0, got %d", result.Failed)
	}
}

func TestProcessTicketsCloseWithError(t *testing.T) {
	processor := NewProcessor()
	processor.SetRateLimitDelay(0)

	ticketIDs := []string{"123", "456", "789"}

	jsonData, _ := json.Marshal(ticketIDs)
	reader := bytes.NewReader(jsonData)

	callCount := 0
	mockClose := func(ctx context.Context, ticketID string) error {
		callCount++
		if callCount == 2 {
			return context.DeadlineExceeded
		}
		return nil
	}

	result := processor.ProcessTicketsClose(context.Background(), reader, mockClose)

	if result.Total != 3 {
		t.Errorf("Expected total 3, got %d", result.Total)
	}
	if result.Succeeded != 2 {
		t.Errorf("Expected succeeded 2, got %d", result.Succeeded)
	}
	if result.Failed != 1 {
		t.Errorf("Expected failed 1, got %d", result.Failed)
	}
}

func TestReadStdin(t *testing.T) {
	// This test would require stdin manipulation, skip for now
	// The function is tested indirectly through batch command tests
}

func TestBatchResultJSON(t *testing.T) {
	result := &models.BatchResponse{
		Total:     3,
		Succeeded: 2,
		Failed:    1,
		Results: []models.BatchResult{
			{Index: 0, Status: "success"},
			{Index: 1, Status: "success"},
			{Index: 2, Status: "failed", Error: "API error"},
		},
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		t.Errorf("Failed to marshal BatchResponse: %v", err)
	}

	var unmarshaled models.BatchResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal BatchResponse: %v", err)
	}

	if unmarshaled.Total != result.Total {
		t.Errorf("Expected Total %d, got %d", result.Total, unmarshaled.Total)
	}
}