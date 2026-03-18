package batch

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type Processor struct {
	rateLimitDelay time.Duration
}

func NewProcessor() *Processor {
	return &Processor{
		rateLimitDelay: 100 * time.Millisecond,
	}
}

func (p *Processor) SetRateLimitDelay(delay time.Duration) {
	p.rateLimitDelay = delay
}

func (p *Processor) ProcessTicketsCreate(ctx context.Context, reader io.Reader, createFn func(ctx context.Context, req models.TicketCreateRequest) (*models.Ticket, error)) *models.BatchResponse {
	var items []models.TicketCreateRequest
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&items); err != nil {
		return &models.BatchResponse{
			Total:  0,
			Errors: []string{fmt.Sprintf("failed to parse input: %v", err)},
		}
	}

	response := &models.BatchResponse{
		Total:   len(items),
		Results: make([]models.BatchResult, len(items)),
	}

	for i, item := range items {
		time.Sleep(p.rateLimitDelay)

		result := models.BatchResult{
			Index: i,
		}

		ticket, err := createFn(ctx, item)
		if err != nil {
			result.Status = "failed"
			result.Error = err.Error()
			response.Failed++
		} else {
			result.Status = "success"
			result.Data = ticket
			response.Succeeded++
		}

		response.Results[i] = result
	}

	return response
}

func (p *Processor) ProcessTicketsUpdate(ctx context.Context, reader io.Reader, updateFn func(ctx context.Context, ticketID string, req models.TicketUpdateRequest) (*models.Ticket, error)) *models.BatchResponse {
	var items []struct {
		ID      string                  `json:"id"`
		Status  string                  `json:"status,omitempty"`
		Priority string                 `json:"priority,omitempty"`
		AssigneeID string                `json:"assigneeId,omitempty"`
	}

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&items); err != nil {
		return &models.BatchResponse{
			Total:  0,
			Errors: []string{fmt.Sprintf("failed to parse input: %v", err)},
		}
	}

	response := &models.BatchResponse{
		Total:   len(items),
		Results: make([]models.BatchResult, len(items)),
	}

	for i, item := range items {
		time.Sleep(p.rateLimitDelay)

		result := models.BatchResult{
			Index: i,
		}

		req := models.TicketUpdateRequest{
			Status:     item.Status,
			Priority:   item.Priority,
			AssigneeID: item.AssigneeID,
		}

		ticket, err := updateFn(ctx, item.ID, req)
		if err != nil {
			result.Status = "failed"
			result.Error = err.Error()
			response.Failed++
		} else {
			result.Status = "success"
			result.Data = ticket
			response.Succeeded++
		}

		response.Results[i] = result
	}

	return response
}

func (p *Processor) ProcessTicketsClose(ctx context.Context, reader io.Reader, closeFn func(ctx context.Context, ticketID string) error) *models.BatchResponse {
	var ticketIDs []string

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&ticketIDs); err != nil {
		return &models.BatchResponse{
			Total:  0,
			Errors: []string{fmt.Sprintf("failed to parse input: %v", err)},
		}
	}

	response := &models.BatchResponse{
		Total:   len(ticketIDs),
		Results: make([]models.BatchResult, len(ticketIDs)),
	}

	for i, ticketID := range ticketIDs {
		time.Sleep(p.rateLimitDelay)

		result := models.BatchResult{
			Index: i,
		}

		if err := closeFn(ctx, ticketID); err != nil {
			result.Status = "failed"
			result.Error = err.Error()
			response.Failed++
		} else {
			result.Status = "success"
			result.Data = map[string]string{"ticketId": ticketID, "status": "closed"}
			response.Succeeded++
		}

		response.Results[i] = result
	}

	return response
}

func (p *Processor) ReadStdin() ([]byte, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, fmt.Errorf("no data on stdin. Pipe JSON data or use --batch with a file")
	}

	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read stdin: %w", err)
	}

	return data, nil
}