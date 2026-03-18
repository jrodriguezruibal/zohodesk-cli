package models

type Ticket struct {
	ID                string `json:"id" yaml:"id"`
	TicketNumber      string `json:"ticketNumber" yaml:"ticketNumber"`
	Subject           string `json:"subject" yaml:"subject"`
	Description       string `json:"description" yaml:"description"`
	Status            string `json:"status" yaml:"status"`
	Priority          string `json:"priority" yaml:"priority"`
	DepartmentID      string `json:"departmentId,omitempty" yaml:"departmentId,omitempty"`
	DepartmentName    string `json:"departmentName,omitempty" yaml:"departmentName,omitempty"`
	ContactID         string `json:"contactId,omitempty" yaml:"contactId,omitempty"`
	ContactName       string `json:"contactName,omitempty" yaml:"contactName,omitempty"`
	ContactEmail      string `json:"contactEmail,omitempty" yaml:"contactEmail,omitempty"`
	AssigneeID        string `json:"assigneeId,omitempty" yaml:"assigneeId,omitempty"`
	AssigneeName      string `json:"assigneeName,omitempty" yaml:"assigneeName,omitempty"`
	CreatedTime       string `json:"createdTime" yaml:"createdTime"`
	ModifiedTime      string `json:"modifiedTime" yaml:"modifiedTime"`
	DueDate           string `json:"dueDate,omitempty" yaml:"dueDate,omitempty"`
	Classification    string `json:"classification,omitempty" yaml:"classification,omitempty"`
	ClosedTime        string `json:"closedTime,omitempty" yaml:"closedTime,omitempty"`
	Resolution        string `json:"resolution,omitempty" yaml:"resolution,omitempty"`
	WebURL            string `json:"webUrl,omitempty" yaml:"webUrl,omitempty"`
}

type TicketListResponse struct {
	Data []Ticket `json:"data"`
}

type TicketCreateRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description,omitempty"`
	Priority    string `json:"priority,omitempty"`
	DepartmentID string `json:"departmentId,omitempty"`
	ContactID   string `json:"contactId,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

type TicketUpdateRequest struct {
	Status       string `json:"status,omitempty"`
	Priority     string `json:"priority,omitempty"`
	AssigneeID   string `json:"assigneeId,omitempty"`
	DepartmentID string `json:"departmentId,omitempty"`
	Resolution   string `json:"resolution,omitempty"`
}

type Contact struct {
	ID          string `json:"id" yaml:"id"`
	FirstName   string `json:"firstName,omitempty" yaml:"firstName"`
	LastName    string `json:"lastName,omitempty" yaml:"lastName"`
	Email       string `json:"email,omitempty" yaml:"email"`
	Phone       string `json:"phone,omitempty" yaml:"phone"`
	Mobile      string `json:"mobile,omitempty" yaml:"mobile"`
	Department  string `json:"department,omitempty" yaml:"department"`
	Title       string `json:"title,omitempty" yaml:"title"`
	CreatedTime string `json:"createdTime,omitempty" yaml:"createdTime"`
}

type ContactListResponse struct {
	Data []Contact `json:"data"`
}

type Thread struct {
	ID          string `json:"id" yaml:"id"`
	TicketID    string `json:"ticketId" yaml:"ticketId"`
	From        string `json:"from,omitempty" yaml:"from"`
	To          string `json:"to,omitempty" yaml:"to"`
	CC          string `json:"cc,omitempty" yaml:"cc"`
	Subject     string `json:"subject,omitempty" yaml:"subject"`
	Description string `json:"description,omitempty" yaml:"description"`
	CreatedTime string `json:"createdTime" yaml:"createdTime"`
	SenderEmail string `json:"senderEmail,omitempty" yaml:"senderEmail"`
	SenderName  string `json:"senderName,omitempty" yaml:"senderName"`
	ThreadType  string `json:"threadType,omitempty" yaml:"threadType"`
}

type ThreadListResponse struct {
	Data []Thread `json:"data"`
}

type FullTicket struct {
	Ticket           Ticket    `json:"ticket" yaml:"ticket"`
	Contact          *Contact  `json:"contact,omitempty" yaml:"contact,omitempty"`
	Threads          []Thread  `json:"threads,omitempty" yaml:"threads,omitempty"`
	ThreadCount      int       `json:"threadCount" yaml:"threadCount"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type SearchParams struct {
	Email  string `json:"email,omitempty"`
	Query  string `json:"query,omitempty"`
	Status string `json:"status,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
}