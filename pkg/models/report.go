package models

type TicketStats struct {
	Total      int            `json:"total" yaml:"total"`
	Open       int            `json:"open" yaml:"open"`
	Closed     int            `json:"closed" yaml:"closed"`
	OnHold     int            `json:"onHold" yaml:"onHold"`
	InProgress int            `json:"inProgress" yaml:"inProgress"`
	ByStatus   []StatusCount  `json:"byStatus" yaml:"byStatus"`
	ByPriority []PriorityCount `json:"byPriority" yaml:"byPriority"`
	ByDepartment []DepartmentCount `json:"byDepartment" yaml:"byDepartment"`
}

type StatusCount struct {
	Status string `json:"status" yaml:"status"`
	Count  int    `json:"count" yaml:"count"`
}

type PriorityCount struct {
	Priority string `json:"priority" yaml:"priority"`
	Count    int    `json:"count" yaml:"count"`
}

type DepartmentCount struct {
	DepartmentID   string `json:"departmentId" yaml:"departmentId"`
	DepartmentName string `json:"departmentName" yaml:"departmentName"`
	Count          int    `json:"count" yaml:"count"`
}

type AgentStats struct {
	AgentID         string  `json:"agentId" yaml:"agentId"`
	AgentName       string  `json:"agentName" yaml:"agentName"`
	TicketsAssigned int     `json:"ticketsAssigned" yaml:"ticketsAssigned"`
	TicketsResolved int     `json:"ticketsResolved" yaml:"ticketsResolved"`
	AvgResponseTime string  `json:"avgResponseTime,omitempty" yaml:"avgResponseTime,omitempty"`
	AvgResolutionTime string `json:"avgResolutionTime,omitempty" yaml:"avgResolutionTime,omitempty"`
	SatisfactionScore float64 `json:"satisfactionScore,omitempty" yaml:"satisfactionScore,omitempty"`
}

type AgentStatsList struct {
	Data []AgentStats `json:"data"`
}

type SLAStats struct {
	TotalTickets    int     `json:"totalTickets" yaml:"totalTickets"`
	Complied        int     `json:"complied" yaml:"complied"`
	Violated        int     `json:"violated" yaml:"violated"`
	ComplianceRate  float64 `json:"complianceRate" yaml:"complianceRate"`
	ByPriority      []SLAByPriority `json:"byPriority" yaml:"byPriority"`
}

type SLAByPriority struct {
	Priority string  `json:"priority" yaml:"priority"`
	Total    int     `json:"total" yaml:"total"`
	Complied int     `json:"complied" yaml:"complied"`
	Rate     float64 `json:"rate" yaml:"rate"`
}