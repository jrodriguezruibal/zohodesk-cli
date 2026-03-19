package models

import (
	"encoding/json"
	"testing"
)

func TestTicketStatsJSON(t *testing.T) {
	stats := TicketStats{
		Total:      100,
		Open:       30,
		Closed:     50,
		OnHold:     10,
		InProgress: 10,
		ByStatus: []StatusCount{
			{Status: "Open", Count: 30},
			{Status: "Closed", Count: 50},
		},
		ByPriority: []PriorityCount{
			{Priority: "High", Count: 20},
			{Priority: "Medium", Count: 50},
		},
	}

	jsonData, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Failed to marshal TicketStats: %v", err)
	}

	var unmarshaled TicketStats
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal TicketStats: %v", err)
	}

	if unmarshaled.Total != stats.Total {
		t.Errorf("Expected Total %d, got %d", stats.Total, unmarshaled.Total)
	}
	if unmarshaled.Open != stats.Open {
		t.Errorf("Expected Open %d, got %d", stats.Open, unmarshaled.Open)
	}
	if len(unmarshaled.ByStatus) != len(stats.ByStatus) {
		t.Errorf("Expected %d status counts, got %d", len(stats.ByStatus), len(unmarshaled.ByStatus))
	}
}

func TestAgentStatsJSON(t *testing.T) {
	stats := AgentStats{
		AgentID:          "123",
		AgentName:        "John Doe",
		TicketsAssigned:  50,
		TicketsResolved:  45,
		AvgResponseTime:  "2h 30m",
		AvgResolutionTime: "24h",
		SatisfactionScore: 4.5,
	}

	jsonData, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Failed to marshal AgentStats: %v", err)
	}

	var unmarshaled AgentStats
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal AgentStats: %v", err)
	}

	if unmarshaled.AgentName != stats.AgentName {
		t.Errorf("Expected AgentName %s, got %s", stats.AgentName, unmarshaled.AgentName)
	}
	if unmarshaled.TicketsResolved != stats.TicketsResolved {
		t.Errorf("Expected TicketsResolved %d, got %d", stats.TicketsResolved, unmarshaled.TicketsResolved)
	}
}

func TestSLAStatsJSON(t *testing.T) {
	stats := SLAStats{
		TotalTickets:   100,
		Complied:       85,
		Violated:       15,
		ComplianceRate: 85.0,
		ByPriority: []SLAByPriority{
			{Priority: "High", Total: 30, Complied: 25, Rate: 83.3},
			{Priority: "Medium", Total: 70, Complied: 60, Rate: 85.7},
		},
	}

	jsonData, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Failed to marshal SLAStats: %v", err)
	}

	var unmarshaled SLAStats
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal SLAStats: %v", err)
	}

	if unmarshaled.ComplianceRate != stats.ComplianceRate {
		t.Errorf("Expected ComplianceRate %.1f, got %.1f", stats.ComplianceRate, unmarshaled.ComplianceRate)
	}
	if len(unmarshaled.ByPriority) != len(stats.ByPriority) {
		t.Errorf("Expected %d priority stats, got %d", len(stats.ByPriority), len(unmarshaled.ByPriority))
	}
}

func TestAgentStatsListJSON(t *testing.T) {
	list := AgentStatsList{
		Data: []AgentStats{
			{AgentID: "1", AgentName: "Agent 1", TicketsResolved: 10},
			{AgentID: "2", AgentName: "Agent 2", TicketsResolved: 20},
		},
	}

	jsonData, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Failed to marshal AgentStatsList: %v", err)
	}

	var unmarshaled AgentStatsList
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal AgentStatsList: %v", err)
	}

	if len(unmarshaled.Data) != len(list.Data) {
		t.Errorf("Expected %d agents, got %d", len(list.Data), len(unmarshaled.Data))
	}
}