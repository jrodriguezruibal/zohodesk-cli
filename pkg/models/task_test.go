package models

import (
	"encoding/json"
	"testing"
)

func TestTaskJSON(t *testing.T) {
	task := Task{
		ID:          "123",
		TicketID:    "456",
		Title:       "Review code changes",
		Description: "Review the pull request",
		Status:      "Open",
		Priority:    "High",
		OwnerID:     "789",
		OwnerName:   "John Doe",
		DueDate:     "2024-01-20",
		CreatedTime: "2024-01-15T10:30:00Z",
	}

	jsonData, err := json.Marshal(task)
	if err != nil {
		t.Errorf("Failed to marshal Task: %v", err)
	}

	var unmarshaled Task
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Task: %v", err)
	}

	if unmarshaled.Title != task.Title {
		t.Errorf("Expected Title %s, got %s", task.Title, unmarshaled.Title)
	}
	if unmarshaled.Status != task.Status {
		t.Errorf("Expected Status %s, got %s", task.Status, unmarshaled.Status)
	}
	if unmarshaled.Priority != task.Priority {
		t.Errorf("Expected Priority %s, got %s", task.Priority, unmarshaled.Priority)
	}
}

func TestTaskCreateRequestJSON(t *testing.T) {
	req := TaskCreateRequest{
		Title:       "New task",
		Description: "Task description",
		Priority:    "Medium",
		OwnerID:     "123",
		DueDate:     "2024-01-20",
		TicketID:    "456",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal TaskCreateRequest: %v", err)
	}

	var unmarshaled TaskCreateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TaskCreateRequest: %v", err)
	}

	if unmarshaled.Title != req.Title {
		t.Errorf("Expected Title %s, got %s", req.Title, unmarshaled.Title)
	}
	if unmarshaled.Priority != req.Priority {
		t.Errorf("Expected Priority %s, got %s", req.Priority, unmarshaled.Priority)
	}
}

func TestTaskUpdateRequestJSON(t *testing.T) {
	req := TaskUpdateRequest{
		Title:   "Updated title",
		Status:  "Completed",
		Priority: "Low",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal TaskUpdateRequest: %v", err)
	}

	var unmarshaled TaskUpdateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TaskUpdateRequest: %v", err)
	}

	if unmarshaled.Title != req.Title {
		t.Errorf("Expected Title %s, got %s", req.Title, unmarshaled.Title)
	}
	if unmarshaled.Status != req.Status {
		t.Errorf("Expected Status %s, got %s", req.Status, unmarshaled.Status)
	}
}

func TestTaskListResponseJSON(t *testing.T) {
	response := TaskListResponse{
		Data: []Task{
			{ID: "1", Title: "Task 1", Status: "Open"},
			{ID: "2", Title: "Task 2", Status: "Completed"},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal TaskListResponse: %v", err)
	}

	var unmarshaled TaskListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TaskListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d tasks, got %d", len(response.Data), len(unmarshaled.Data))
	}
	if unmarshaled.Data[0].Title != response.Data[0].Title {
		t.Errorf("Expected first task title %s, got %s", response.Data[0].Title, unmarshaled.Data[0].Title)
	}
}