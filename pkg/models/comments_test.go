package models

import (
	"encoding/json"
	"testing"
)

func TestCommentJSON(t *testing.T) {
	comment := Comment{
		ID:          "123",
		TicketID:    "456",
		Content:     "Test comment",
		AuthorName:  "John Doe",
		AuthorEmail: "john@example.com",
		IsPublic:    true,
		CreatedTime: "2024-01-15T10:30:00Z",
	}

	jsonData, err := json.Marshal(comment)
	if err != nil {
		t.Errorf("Failed to marshal Comment: %v", err)
	}

	var unmarshaled Comment
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Comment: %v", err)
	}

	if unmarshaled.ID != comment.ID {
		t.Errorf("Expected ID %s, got %s", comment.ID, unmarshaled.ID)
	}
	if unmarshaled.Content != comment.Content {
		t.Errorf("Expected Content %s, got %s", comment.Content, unmarshaled.Content)
	}
	if unmarshaled.IsPublic != comment.IsPublic {
		t.Errorf("Expected IsPublic %v, got %v", comment.IsPublic, unmarshaled.IsPublic)
	}
}

func TestCommentCreateRequestJSON(t *testing.T) {
	req := CommentCreateRequest{
		Content:  "Test comment",
		IsPublic: true,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal CommentCreateRequest: %v", err)
	}

	var unmarshaled CommentCreateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal CommentCreateRequest: %v", err)
	}

	if unmarshaled.Content != req.Content {
		t.Errorf("Expected Content %s, got %s", req.Content, unmarshaled.Content)
	}
}

func TestBatchResultJSON(t *testing.T) {
	result := BatchResult{
		Index:  0,
		Status: "success",
		Data:   map[string]string{"ticketId": "123"},
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		t.Errorf("Failed to marshal BatchResult: %v", err)
	}

	var unmarshaled BatchResult
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal BatchResult: %v", err)
	}

	if unmarshaled.Index != result.Index {
		t.Errorf("Expected Index %d, got %d", result.Index, unmarshaled.Index)
	}
	if unmarshaled.Status != result.Status {
		t.Errorf("Expected Status %s, got %s", result.Status, unmarshaled.Status)
	}
}

func TestBatchResponseJSON(t *testing.T) {
	response := BatchResponse{
		Total:     10,
		Succeeded: 8,
		Failed:    2,
		Results: []BatchResult{
			{Index: 0, Status: "success"},
			{Index: 1, Status: "failed", Error: "API error"},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal BatchResponse: %v", err)
	}

	var unmarshaled BatchResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal BatchResponse: %v", err)
	}

	if unmarshaled.Total != response.Total {
		t.Errorf("Expected Total %d, got %d", response.Total, unmarshaled.Total)
	}
	if unmarshaled.Succeeded != response.Succeeded {
		t.Errorf("Expected Succeeded %d, got %d", response.Succeeded, unmarshaled.Succeeded)
	}
	if unmarshaled.Failed != response.Failed {
		t.Errorf("Expected Failed %d, got %d", response.Failed, unmarshaled.Failed)
	}
	if len(unmarshaled.Results) != len(response.Results) {
		t.Errorf("Expected %d Results, got %d", len(response.Results), len(unmarshaled.Results))
	}
}

func TestAttachmentJSON(t *testing.T) {
	attachment := Attachment{
		ID:          "123",
		Name:        "document.pdf",
		Size:        1024,
		ContentType: "application/pdf",
		URL:         "https://example.com/document.pdf",
	}

	jsonData, err := json.Marshal(attachment)
	if err != nil {
		t.Errorf("Failed to marshal Attachment: %v", err)
	}

	var unmarshaled Attachment
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Attachment: %v", err)
	}

	if unmarshaled.Name != attachment.Name {
		t.Errorf("Expected Name %s, got %s", attachment.Name, unmarshaled.Name)
	}
	if unmarshaled.Size != attachment.Size {
		t.Errorf("Expected Size %d, got %d", attachment.Size, unmarshaled.Size)
	}
}

func TestAgentJSON(t *testing.T) {
	agent := Agent{
		ID:       "123",
		Name:     "John Doe",
		Email:    "john@example.com",
		Role:     "Agent",
		IsActive: true,
	}

	jsonData, err := json.Marshal(agent)
	if err != nil {
		t.Errorf("Failed to marshal Agent: %v", err)
	}

	var unmarshaled Agent
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Agent: %v", err)
	}

	if unmarshaled.Name != agent.Name {
		t.Errorf("Expected Name %s, got %s", agent.Name, unmarshaled.Name)
	}
}

func TestDepartmentJSON(t *testing.T) {
	dept := Department{
		ID:          "123",
		Name:        "Support",
		Description: "Customer Support Department",
		IsVisible:   true,
	}

	jsonData, err := json.Marshal(dept)
	if err != nil {
		t.Errorf("Failed to marshal Department: %v", err)
	}

	var unmarshaled Department
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Department: %v", err)
	}

	if unmarshaled.Name != dept.Name {
		t.Errorf("Expected Name %s, got %s", dept.Name, unmarshaled.Name)
	}
}