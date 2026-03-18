package models

type Task struct {
	ID           string `json:"id" yaml:"id"`
	TicketID     string `json:"ticketId,omitempty" yaml:"ticketId,omitempty"`
	Title        string `json:"title" yaml:"title"`
	Description  string `json:"description,omitempty" yaml:"description,omitempty"`
	Status       string `json:"status,omitempty" yaml:"status,omitempty"`
	Priority     string `json:"priority,omitempty" yaml:"priority,omitempty"`
	OwnerID      string `json:"ownerId,omitempty" yaml:"ownerId,omitempty"`
	OwnerName    string `json:"ownerName,omitempty" yaml:"ownerName,omitempty"`
	DueDate      string `json:"dueDate,omitempty" yaml:"dueDate,omitempty"`
	CompletedAt  string `json:"completedAt,omitempty" yaml:"completedAt,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
	ModifiedTime string `json:"modifiedTime,omitempty" yaml:"modifiedTime,omitempty"`
}

type TaskListResponse struct {
	Data []Task `json:"data"`
}

type TaskCreateRequest struct {
	TicketID    string `json:"ticketId,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Priority    string `json:"priority,omitempty"`
	OwnerID     string `json:"ownerId,omitempty"`
	DueDate     string `json:"dueDate,omitempty"`
}

type TaskUpdateRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	Priority    string `json:"priority,omitempty"`
	OwnerID     string `json:"ownerId,omitempty"`
	DueDate     string `json:"dueDate,omitempty"`
}