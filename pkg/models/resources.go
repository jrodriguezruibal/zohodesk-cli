package models

type Department struct {
	ID           string `json:"id" yaml:"id"`
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description,omitempty" yaml:"description,omitempty"`
	IsVisible    bool   `json:"isVisible" yaml:"isVisible"`
	CreatedTime  string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
	ModifiedTime string `json:"modifiedTime,omitempty" yaml:"modifiedTime,omitempty"`
}

type DepartmentListResponse struct {
	Data []Department `json:"data"`
}

type Agent struct {
	ID           string `json:"id" yaml:"id"`
	Name         string `json:"name" yaml:"name"`
	Email        string `json:"email" yaml:"email"`
	Role         string `json:"role,omitempty" yaml:"role,omitempty"`
	DepartmentID string `json:"departmentId,omitempty" yaml:"departmentId,omitempty"`
	IsActive     bool   `json:"isActive" yaml:"isActive"`
	Phone        string `json:"phone,omitempty" yaml:"phone,omitempty"`
	Mobile       string `json:"mobile,omitempty" yaml:"mobile,omitempty"`
	PhotoURL     string `json:"photoUrl,omitempty" yaml:"photoUrl,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
}

type AgentListResponse struct {
	Data []Agent `json:"data"`
}

type TicketContext struct {
	Ticket        *Ticket     `json:"ticket" yaml:"ticket"`
	Contact       *Contact    `json:"contact,omitempty" yaml:"contact,omitempty"`
	Department    *Department `json:"department,omitempty" yaml:"department,omitempty"`
	Assignee      *Agent      `json:"assignee,omitempty" yaml:"assignee,omitempty"`
	SLA           *SLA        `json:"sla,omitempty" yaml:"sla,omitempty"`
	Threads       []Thread    `json:"threads,omitempty" yaml:"threads,omitempty"`
	ThreadCount   int         `json:"threadCount" yaml:"threadCount"`
	Tags         []string    `json:"tags,omitempty" yaml:"tags,omitempty"`
	RecentActivity []Activity `json:"recentActivity,omitempty" yaml:"recentActivity,omitempty"`
}

type SLA struct {
	ID            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	DueDate       string `json:"dueDate,omitempty" yaml:"dueDate,omitempty"`
	ResponseDue   string `json:"responseDue,omitempty" yaml:"responseDue,omitempty"`
	Status        string `json:"status,omitempty" yaml:"status,omitempty"`
	RemainingTime string `json:"remainingTime,omitempty" yaml:"remainingTime,omitempty"`
}

type Activity struct {
	ID          string `json:"id" yaml:"id"`
	TicketID    string `json:"ticketId" yaml:"ticketId"`
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	PerformedBy string `json:"performedBy,omitempty" yaml:"performedBy,omitempty"`
	PerformedAt string `json:"performedAt" yaml:"performedAt"`
}