package models

import (
	"encoding/json"
	"testing"
)

func TestTagJSON(t *testing.T) {
	tag := Tag{
		ID:   "123",
		Name: "urgent",
	}

	jsonData, err := json.Marshal(tag)
	if err != nil {
		t.Errorf("Failed to marshal Tag: %v", err)
	}

	var unmarshaled Tag
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Tag: %v", err)
	}

	if unmarshaled.Name != tag.Name {
		t.Errorf("Expected Name %s, got %s", tag.Name, unmarshaled.Name)
	}
}

func TestTagListResponseJSON(t *testing.T) {
	response := TagListResponse{
		Data: []Tag{
			{ID: "1", Name: "urgent"},
			{ID: "2", Name: "bug"},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal TagListResponse: %v", err)
	}

	var unmarshaled TagListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TagListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d tags, got %d", len(response.Data), len(unmarshaled.Data))
	}
}

func TestTicketTagsRequestJSON(t *testing.T) {
	req := TicketTagsRequest{
		Tags: []string{"urgent", "bug", "customer"},
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal TicketTagsRequest: %v", err)
	}

	var unmarshaled TicketTagsRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TicketTagsRequest: %v", err)
	}

	if len(unmarshaled.Tags) != len(req.Tags) {
		t.Errorf("Expected %d tags, got %d", len(req.Tags), len(unmarshaled.Tags))
	}
}

func TestTicketMergeRequestJSON(t *testing.T) {
	req := TicketMergeRequest{
		TargetTicketID: "456",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal TicketMergeRequest: %v", err)
	}

	var unmarshaled TicketMergeRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TicketMergeRequest: %v", err)
	}

	if unmarshaled.TargetTicketID != req.TargetTicketID {
		t.Errorf("Expected TargetTicketID %s, got %s", req.TargetTicketID, unmarshaled.TargetTicketID)
	}
}

func TestFollowRequestJSON(t *testing.T) {
	req := FollowRequest{Follow: true}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal FollowRequest: %v", err)
	}

	var unmarshaled FollowRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal FollowRequest: %v", err)
	}

	if unmarshaled.Follow != req.Follow {
		t.Errorf("Expected Follow %v, got %v", req.Follow, unmarshaled.Follow)
	}
}

func TestTicketWithTags(t *testing.T) {
	ticket := Ticket{
		ID:     "123",
		Subject: "Test",
		Tags:   []string{"urgent", "bug"},
		CustomFields: map[string]interface{}{
			"field1": "value1",
			"field2": 123,
		},
	}

	jsonData, err := json.Marshal(ticket)
	if err != nil {
		t.Errorf("Failed to marshal Ticket with tags: %v", err)
	}

	var unmarshaled Ticket
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Ticket with tags: %v", err)
	}

	if len(unmarshaled.Tags) != len(ticket.Tags) {
		t.Errorf("Expected %d tags, got %d", len(ticket.Tags), len(unmarshaled.Tags))
	}
	if unmarshaled.CustomFields["field1"] != ticket.CustomFields["field1"] {
		t.Errorf("Expected custom field field1 = %v, got %v", ticket.CustomFields["field1"], unmarshaled.CustomFields["field1"])
	}
}