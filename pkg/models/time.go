package models

type TimeEntry struct {
	ID           string `json:"id" yaml:"id"`
	TicketID     string `json:"ticketId" yaml:"ticketId"`
	AgentID      string `json:"agentId,omitempty" yaml:"agentId,omitempty"`
	AgentName    string `json:"agentName,omitempty" yaml:"agentName,omitempty"`
	Hours        int    `json:"hours" yaml:"hours"`
	Minutes      int    `json:"minutes" yaml:"minutes"`
	TotalMinutes int    `json:"totalMinutes" yaml:"totalMinutes"`
	Description  string `json:"description,omitempty" yaml:"description,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
	ExecutedTime string `json:"executedTime,omitempty" yaml:"executedTime,omitempty"`
	ParentTicket struct {
		ID            string `json:"id,omitempty" yaml:"id,omitempty"`
		TicketNumber  string `json:"ticketNumber,omitempty" yaml:"ticketNumber,omitempty"`
		Subject       string `json:"subject,omitempty" yaml:"subject,omitempty"`
	} `json:"parentTicket,omitempty" yaml:"parentTicket,omitempty"`
}

type TimeEntryListResponse struct {
	Data []TimeEntry `json:"data"`
}

type TimeEntryCreateRequest struct {
	Hours       int    `json:"hours,omitempty"`
	Minutes     int    `json:"minutes,omitempty"`
	Description string `json:"description,omitempty"`
	AgentID     string `json:"agentId,omitempty"`
	ExecutedAt  string `json:"executedAt,omitempty"`
}

type TimeReport struct {
	TotalTime       int                `json:"totalTime" yaml:"totalTime"`
	TotalHours      float64            `json:"totalHours" yaml:"totalHours"`
	EntriesCount    int                `json:"entriesCount" yaml:"entriesCount"`
	EntriesByAgent  []TimeEntryByAgent `json:"entriesByAgent,omitempty" yaml:"entriesByAgent,omitempty"`
	EntriesByTicket []TimeEntryByTicket `json:"entriesByTicket,omitempty" yaml:"entriesByTicket,omitempty"`
}

type TimeEntryByAgent struct {
	AgentID   string  `json:"agentId" yaml:"agentId"`
	AgentName string  `json:"agentName" yaml:"agentName"`
	TotalTime int     `json:"totalTime" yaml:"totalTime"`
	Hours     float64 `json:"hours" yaml:"hours"`
	Count     int     `json:"count" yaml:"count"`
}

type TimeEntryByTicket struct {
	TicketID     string `json:"ticketId" yaml:"ticketId"`
	TicketNumber string `json:"ticketNumber" yaml:"ticketNumber"`
	Subject      string `json:"subject" yaml:"subject"`
	TotalTime    int    `json:"totalTime" yaml:"totalTime"`
	Hours        float64 `json:"hours" yaml:"hours"`
	Count        int    `json:"count" yaml:"count"`
}

func (t *TimeEntry) DurationString() string {
	if t.Hours > 0 && t.Minutes > 0 {
		return string(rune(t.Hours)) + "h " + string(rune(t.Minutes)) + "m"
	} else if t.Hours > 0 {
		return string(rune(t.Hours)) + "h"
	} else if t.Minutes > 0 {
		return string(rune(t.Minutes)) + "m"
	}
	return "0m"
}