package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type TicketsService struct {
	client *Client
}

func NewTicketsService(client *Client) *TicketsService {
	return &TicketsService{client: client}
}

func (s *TicketsService) List(ctx context.Context, params map[string]string) ([]models.Ticket, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/tickets", query)
	if err != nil {
		return nil, err
	}

	var resp models.TicketListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *TicketsService) Get(ctx context.Context, ticketID string) (*models.Ticket, error) {
	data, err := s.client.Get(ctx, "/tickets/"+ticketID, nil)
	if err != nil {
		return nil, err
	}

	var ticket models.Ticket
	if err := json.Unmarshal(data, &ticket); err != nil {
		return nil, fmt.Errorf("failed to parse ticket response: %w", err)
	}

	if ticket.ID == "" {
		var resp struct {
			Data models.Ticket `json:"data"`
		}
		if err := json.Unmarshal(data, &resp); err == nil && resp.Data.ID != "" {
			return &resp.Data, nil
		}
	}

	return &ticket, nil
}

func (s *TicketsService) GetFull(ctx context.Context, ticketID string) (*models.FullTicket, error) {
	ticket, err := s.Get(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	threads, err := s.GetThreads(ctx, ticketID)
	if err != nil {
		threads = nil
	}

	var contact *models.Contact
	if ticket.ContactID != "" {
		contact, _ = s.GetContact(ctx, ticket.ContactID)
	}

	return &models.FullTicket{
		Ticket:      *ticket,
		Contact:     contact,
		Threads:     threads,
		ThreadCount: len(threads),
	}, nil
}

func (s *TicketsService) GetWithContext(ctx context.Context, ticketID string) (*models.TicketContext, error) {
	ticket, err := s.Get(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	context := &models.TicketContext{
		Ticket: ticket,
	}

	// Get contact
	if ticket.ContactID != "" {
		context.Contact, _ = s.GetContact(ctx, ticket.ContactID)
	}

	// Get threads
	context.Threads, _ = s.GetThreads(ctx, ticketID)
	if context.Threads != nil {
		context.ThreadCount = len(context.Threads)
	}

	// Get department
	if ticket.DepartmentID != "" {
		dept, err := NewDepartmentsService(s.client).Get(ctx, ticket.DepartmentID)
		if err == nil {
			context.Department = dept
		}
	}

	// Get assignee
	if ticket.AssigneeID != "" {
		agent, err := NewAgentsService(s.client).Get(ctx, ticket.AssigneeID)
		if err == nil {
			context.Assignee = agent
		}
	}

	return context, nil
}

func (s *TicketsService) Assign(ctx context.Context, ticketID string, agentID string, departmentID string) (*models.Ticket, error) {
	req := models.TicketUpdateRequest{
		AssigneeID:   agentID,
		DepartmentID: departmentID,
	}
	return s.Update(ctx, ticketID, req)
}

func (s *TicketsService) Merge(ctx context.Context, sourceID string, targetID string) error {
	req := models.TicketMergeRequest{
		TargetTicketID: targetID,
	}
	_, err := s.client.Post(ctx, "/tickets/"+sourceID+"/merge", req)
	return err
}

func (s *TicketsService) Follow(ctx context.Context, ticketID string) error {
	req := models.FollowRequest{Follow: true}
	_, err := s.client.Post(ctx, "/tickets/"+ticketID+"/follow", req)
	return err
}

func (s *TicketsService) Unfollow(ctx context.Context, ticketID string) error {
	return s.client.Delete(ctx, "/tickets/"+ticketID+"/follow")
}

func (s *TicketsService) Create(ctx context.Context, req models.TicketCreateRequest) (*models.Ticket, error) {
	data, err := s.client.Post(ctx, "/tickets", req)
	if err != nil {
		return nil, err
	}

	var ticket models.Ticket
	if err := json.Unmarshal(data, &ticket); err != nil {
		return nil, fmt.Errorf("failed to parse ticket response: %w", err)
	}

	if ticket.ID == "" {
		var resp struct {
			Data models.Ticket `json:"data"`
		}
		if err := json.Unmarshal(data, &resp); err == nil && resp.Data.ID != "" {
			return &resp.Data, nil
		}
	}

	return &ticket, nil
}

func (s *TicketsService) Update(ctx context.Context, ticketID string, req models.TicketUpdateRequest) (*models.Ticket, error) {
	data, err := s.client.Patch(ctx, "/tickets/"+ticketID, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data models.Ticket `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (s *TicketsService) Close(ctx context.Context, ticketID string, resolution string) error {
	req := models.TicketUpdateRequest{
		Status:     "Closed",
		Resolution: resolution,
	}
	_, err := s.Update(ctx, ticketID, req)
	return err
}

func (s *TicketsService) Search(ctx context.Context, params models.SearchParams) ([]models.Ticket, error) {
	tickets, err := s.List(ctx, map[string]string{
		"limit": "100",
	})
	if err != nil {
		return nil, err
	}

	var filtered []models.Ticket
	for _, t := range tickets {
		if params.Email != "" && t.ContactEmail != params.Email {
			continue
		}
		if params.Status != "" && t.Status != params.Status {
			continue
		}
		filtered = append(filtered, t)
	}

	return filtered, nil
}

func (s *TicketsService) GetThreads(ctx context.Context, ticketID string) ([]models.Thread, error) {
	data, err := s.client.Get(ctx, "/tickets/"+ticketID+"/threads", nil)
	if err != nil {
		return nil, err
	}

	var resp models.ThreadListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *TicketsService) GetContact(ctx context.Context, contactID string) (*models.Contact, error) {
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