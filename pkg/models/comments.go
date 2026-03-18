package models

type Comment struct {
	ID            string `json:"id" yaml:"id"`
	TicketID      string `json:"ticketId,omitempty" yaml:"ticketId,omitempty"`
	Content       string `json:"content" yaml:"content"`
	Author        string `json:"author,omitempty" yaml:"author,omitempty"`
	AuthorID      string `json:"authorId,omitempty" yaml:"authorId,omitempty"`
	AuthorName    string `json:"authorName,omitempty" yaml:"authorName,omitempty"`
	AuthorEmail   string `json:"authorEmail,omitempty" yaml:"authorEmail,omitempty"`
	CommentType   string `json:"commentType,omitempty" yaml:"commentType,omitempty"`
	IsPublic      bool   `json:"isPublic" yaml:"isPublic"`
	ParentID      string `json:"parentId,omitempty" yaml:"parentId,omitempty"`
	CreatedTime   string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
	ModifiedTime  string `json:"modifiedTime,omitempty" yaml:"modifiedTime,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

type CommentCreateRequest struct {
	Content     string `json:"content"`
	IsPublic    bool   `json:"isPublic"`
	AuthorID    string `json:"authorId,omitempty"`
	ParentID    string `json:"parentId,omitempty"`
}

type CommentListResponse struct {
	Data []Comment `json:"data"`
}

type Attachment struct {
	ID           string `json:"id" yaml:"id"`
	Name         string `json:"name" yaml:"name"`
	Size         int64  `json:"size" yaml:"size"`
	ContentType  string `json:"contentType" yaml:"contentType"`
	URL          string `json:"url,omitempty" yaml:"url,omitempty"`
	DownloadURL  string `json:"downloadUrl,omitempty" yaml:"downloadUrl,omitempty"`
}

type AttachmentUploadResponse struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Size     int64  `json:"size"`
		MimeType string `json:"mimeType"`
	} `json:"data"`
}

type BatchRequest struct {
	Operation string                   `json:"operation"`
	Items     []map[string]interface{} `json:"items"`
}

type BatchResult struct {
	Index    int         `json:"index"`
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Error    string      `json:"error,omitempty"`
}

type BatchResponse struct {
	Total     int           `json:"total"`
	Succeeded int           `json:"succeeded"`
	Failed    int           `json:"failed"`
	Results   []BatchResult `json:"results"`
	Errors    []string      `json:"errors,omitempty"`
}

type Tag struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

type TagListResponse struct {
	Data []Tag `json:"data"`
}

type TicketTagRequest struct {
	Tags []string `json:"tags"`
}