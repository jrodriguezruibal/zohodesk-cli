package models

import (
	"encoding/json"
	"testing"
)

func TestTimeEntryJSON(t *testing.T) {
	entry := TimeEntry{
		ID:           "123",
		TicketID:     "456",
		AgentID:      "789",
		AgentName:    "John Doe",
		Hours:        2,
		Minutes:      30,
		TotalMinutes: 150,
		Description:  "Investigated the issue",
		CreatedTime:  "2024-01-15T10:30:00Z",
		ExecutedTime: "2024-01-15T09:00:00Z",
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		t.Errorf("Failed to marshal TimeEntry: %v", err)
	}

	var unmarshaled TimeEntry
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TimeEntry: %v", err)
	}

	if unmarshaled.ID != entry.ID {
		t.Errorf("Expected ID %s, got %s", entry.ID, unmarshaled.ID)
	}
	if unmarshaled.AgentName != entry.AgentName {
		t.Errorf("Expected AgentName %s, got %s", entry.AgentName, unmarshaled.AgentName)
	}
	if unmarshaled.Hours != entry.Hours {
		t.Errorf("Expected Hours %d, got %d", entry.Hours, unmarshaled.Hours)
	}
	if unmarshaled.Minutes != entry.Minutes {
		t.Errorf("Expected Minutes %d, got %d", entry.Minutes, unmarshaled.Minutes)
	}
	if unmarshaled.TotalMinutes != entry.TotalMinutes {
		t.Errorf("Expected TotalMinutes %d, got %d", entry.TotalMinutes, unmarshaled.TotalMinutes)
	}
}

func TestTimeEntryCreateRequestJSON(t *testing.T) {
	req := TimeEntryCreateRequest{
		Hours:       1,
		Minutes:     30,
		Description: "Working on ticket",
		AgentID:     "agent123",
		ExecutedAt:  "2024-01-15",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal TimeEntryCreateRequest: %v", err)
	}

	var unmarshaled TimeEntryCreateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TimeEntryCreateRequest: %v", err)
	}

	if unmarshaled.Hours != req.Hours {
		t.Errorf("Expected Hours %d, got %d", req.Hours, unmarshaled.Hours)
	}
	if unmarshaled.Minutes != req.Minutes {
		t.Errorf("Expected Minutes %d, got %d", req.Minutes, unmarshaled.Minutes)
	}
	if unmarshaled.Description != req.Description {
		t.Errorf("Expected Description %s, got %s", req.Description, unmarshaled.Description)
	}
}

func TestTimeEntryListResponseJSON(t *testing.T) {
	response := TimeEntryListResponse{
		Data: []TimeEntry{
			{ID: "1", AgentName: "Agent 1", Hours: 1, Minutes: 30},
			{ID: "2", AgentName: "Agent 2", Hours: 2, Minutes: 0},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal TimeEntryListResponse: %v", err)
	}

	var unmarshaled TimeEntryListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TimeEntryListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d entries, got %d", len(response.Data), len(unmarshaled.Data))
	}
	if unmarshaled.Data[0].AgentName != response.Data[0].AgentName {
		t.Errorf("Expected first agent name %s, got %s", response.Data[0].AgentName, unmarshaled.Data[0].AgentName)
	}
}

func TestTimeReportJSON(t *testing.T) {
	report := TimeReport{
		TotalTime:    480,
		TotalHours:   8.0,
		EntriesCount: 10,
		EntriesByAgent: []TimeEntryByAgent{
			{AgentID: "1", AgentName: "Agent 1", TotalTime: 300, Hours: 5.0, Count: 5},
			{AgentID: "2", AgentName: "Agent 2", TotalTime: 180, Hours: 3.0, Count: 5},
		},
	}

	jsonData, err := json.Marshal(report)
	if err != nil {
		t.Errorf("Failed to marshal TimeReport: %v", err)
	}

	var unmarshaled TimeReport
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TimeReport: %v", err)
	}

	if unmarshaled.TotalHours != report.TotalHours {
		t.Errorf("Expected TotalHours %.1f, got %.1f", report.TotalHours, unmarshaled.TotalHours)
	}
	if unmarshaled.EntriesCount != report.EntriesCount {
		t.Errorf("Expected EntriesCount %d, got %d", report.EntriesCount, unmarshaled.EntriesCount)
	}
	if len(unmarshaled.EntriesByAgent) != len(report.EntriesByAgent) {
		t.Errorf("Expected %d agents, got %d", len(report.EntriesByAgent), len(unmarshaled.EntriesByAgent))
	}
}

func TestTimeEntryByTicketJSON(t *testing.T) {
	entry := TimeEntryByTicket{
		TicketID:     "123",
		TicketNumber: "TKT-001",
		Subject:      "Login issue",
		TotalTime:    120,
		Hours:        2.0,
		Count:        3,
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		t.Errorf("Failed to marshal TimeEntryByTicket: %v", err)
	}

	var unmarshaled TimeEntryByTicket
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TimeEntryByTicket: %v", err)
	}

	if unmarshaled.TicketNumber != entry.TicketNumber {
		t.Errorf("Expected TicketNumber %s, got %s", entry.TicketNumber, unmarshaled.TicketNumber)
	}
	if unmarshaled.Hours != entry.Hours {
		t.Errorf("Expected Hours %.1f, got %.1f", entry.Hours, unmarshaled.Hours)
	}
}