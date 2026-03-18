package models

import (
	"encoding/json"
	"testing"
)

func TestSLAJSON(t *testing.T) {
	sla := SLA{
		ID:            "123",
		Name:          "Premium SLA",
		DueDate:       "2024-01-20T10:00:00Z",
		ResponseDue:   "2024-01-19T10:00:00Z",
		Status:        "Active",
		RemainingTime: "24 hours",
	}

	jsonData, err := json.Marshal(sla)
	if err != nil {
		t.Errorf("Failed to marshal SLA: %v", err)
	}

	var unmarshaled SLA
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal SLA: %v", err)
	}

	if unmarshaled.Name != sla.Name {
		t.Errorf("Expected Name %s, got %s", sla.Name, unmarshaled.Name)
	}
	if unmarshaled.Status != sla.Status {
		t.Errorf("Expected Status %s, got %s", sla.Status, unmarshaled.Status)
	}
}

func TestActivityJSON(t *testing.T) {
	activity := Activity{
		ID:          "123",
		TicketID:    "456",
		Type:        "StatusChange",
		Description: "Status changed from Open to In Progress",
		PerformedBy: "John Doe",
		PerformedAt: "2024-01-15T10:30:00Z",
	}

	jsonData, err := json.Marshal(activity)
	if err != nil {
		t.Errorf("Failed to marshal Activity: %v", err)
	}

	var unmarshaled Activity
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Activity: %v", err)
	}

	if unmarshaled.Type != activity.Type {
		t.Errorf("Expected Type %s, got %s", activity.Type, unmarshaled.Type)
	}
	if unmarshaled.TicketID != activity.TicketID {
		t.Errorf("Expected TicketID %s, got %s", activity.TicketID, unmarshaled.TicketID)
	}
}

func TestTicketContextJSON(t *testing.T) {
	ticket := &Ticket{
		ID:          "123",
		Subject:     "Test Ticket",
		Status:      "Open",
		Priority:    "High",
		ContactID:   "456",
		DepartmentID: "789",
		AssigneeID:  "101",
	}

	contact := &Contact{
		ID:        "456",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	department := &Department{
		ID:          "789",
		Name:        "Support",
		Description: "Customer Support",
	}

	assignee := &Agent{
		ID:       "101",
		Name:     "Agent Smith",
		Email:    "agent@example.com",
		IsActive: true,
	}

	sla := &SLA{
		ID:     "202",
		Name:   "Premium SLA",
		Status: "Active",
	}

	threads := []Thread{
		{ID: "301", TicketID: "123", Description: "First message"},
		{ID: "302", TicketID: "123", Description: "Second message"},
	}

	ctx := TicketContext{
		Ticket:      ticket,
		Contact:     contact,
		Department:  department,
		Assignee:    assignee,
		SLA:         sla,
		Threads:     threads,
		ThreadCount: len(threads),
		Tags:        []string{"urgent", "bug"},
	}

	jsonData, err := json.Marshal(ctx)
	if err != nil {
		t.Errorf("Failed to marshal TicketContext: %v", err)
	}

	var unmarshaled TicketContext
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TicketContext: %v", err)
	}

	if unmarshaled.Ticket.ID != ticket.ID {
		t.Errorf("Expected Ticket ID %s, got %s", ticket.ID, unmarshaled.Ticket.ID)
	}
	if unmarshaled.Contact.FirstName != contact.FirstName {
		t.Errorf("Expected Contact FirstName %s, got %s", contact.FirstName, unmarshaled.Contact.FirstName)
	}
	if unmarshaled.Department.Name != department.Name {
		t.Errorf("Expected Department Name %s, got %s", department.Name, unmarshaled.Department.Name)
	}
	if unmarshaled.Assignee.Name != assignee.Name {
		t.Errorf("Expected Assignee Name %s, got %s", assignee.Name, unmarshaled.Assignee.Name)
	}
	if unmarshaled.ThreadCount != ctx.ThreadCount {
		t.Errorf("Expected ThreadCount %d, got %d", ctx.ThreadCount, unmarshaled.ThreadCount)
	}
}

func TestAgentWithOptionalFields(t *testing.T) {
	agent := Agent{
		ID:           "123",
		Name:         "John Doe",
		Email:        "john@example.com",
		Role:         "Agent",
		DepartmentID: "456",
		IsActive:     true,
		Phone:        "+1234567890",
		Mobile:       "+1234567891",
		PhotoURL:     "https://example.com/photo.jpg",
		CreatedTime:  "2024-01-15T10:00:00Z",
	}

	jsonData, err := json.Marshal(agent)
	if err != nil {
		t.Errorf("Failed to marshal Agent: %v", err)
	}

	var unmarshaled Agent
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Agent: %v", err)
	}

	if unmarshaled.Phone != agent.Phone {
		t.Errorf("Expected Phone %s, got %s", agent.Phone, unmarshaled.Phone)
	}
	if unmarshaled.PhotoURL != agent.PhotoURL {
		t.Errorf("Expected PhotoURL %s, got %s", agent.PhotoURL, unmarshaled.PhotoURL)
	}
}

func TestDepartmentWithOptionalFields(t *testing.T) {
	dept := Department{
		ID:           "123",
		Name:         "Support",
		Description:  "Customer Support Department",
		IsVisible:    true,
		CreatedTime:  "2024-01-01T00:00:00Z",
		ModifiedTime: "2024-01-15T10:00:00Z",
	}

	jsonData, err := json.Marshal(dept)
	if err != nil {
		t.Errorf("Failed to marshal Department: %v", err)
	}

	var unmarshaled Department
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Department: %v", err)
	}

	if unmarshaled.CreatedTime != dept.CreatedTime {
		t.Errorf("Expected CreatedTime %s, got %s", dept.CreatedTime, unmarshaled.CreatedTime)
	}
}

func TestAgentListResponseJSON(t *testing.T) {
	response := AgentListResponse{
		Data: []Agent{
			{ID: "1", Name: "Agent 1", Email: "agent1@example.com", IsActive: true},
			{ID: "2", Name: "Agent 2", Email: "agent2@example.com", IsActive: true},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal AgentListResponse: %v", err)
	}

	var unmarshaled AgentListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal AgentListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d agents, got %d", len(response.Data), len(unmarshaled.Data))
	}
	if unmarshaled.Data[0].Name != response.Data[0].Name {
		t.Errorf("Expected first agent name %s, got %s", response.Data[0].Name, unmarshaled.Data[0].Name)
	}
}

func TestDepartmentListResponseJSON(t *testing.T) {
	response := DepartmentListResponse{
		Data: []Department{
			{ID: "1", Name: "Support", IsVisible: true},
			{ID: "2", Name: "Sales", IsVisible: true},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal DepartmentListResponse: %v", err)
	}

	var unmarshaled DepartmentListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal DepartmentListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d departments, got %d", len(response.Data), len(unmarshaled.Data))
	}
	if unmarshaled.Data[0].Name != response.Data[0].Name {
		t.Errorf("Expected first department name %s, got %s", response.Data[0].Name, unmarshaled.Data[0].Name)
	}
}