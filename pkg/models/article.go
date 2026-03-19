package models

type Article struct {
	ID            string   `json:"id" yaml:"id"`
	Title         string   `json:"title" yaml:"title"`
	Content       string   `json:"content,omitempty" yaml:"content,omitempty"`
	Summary       string   `json:"summary,omitempty" yaml:"summary,omitempty"`
	CategoryID    string   `json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	CategoryName  string   `json:"categoryName,omitempty" yaml:"categoryName,omitempty"`
	AuthorID      string   `json:"authorId,omitempty" yaml:"authorId,omitempty"`
	AuthorName    string   `json:"authorName,omitempty" yaml:"authorName,omitempty"`
	Status        string   `json:"status,omitempty" yaml:"status,omitempty"`
	Tags          []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	ViewCount     string   `json:"viewCount,omitempty" yaml:"viewCount,omitempty"`
	LikeCount     string   `json:"likeCount,omitempty" yaml:"likeCount,omitempty"`
	CreatedTime   string   `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
	ModifiedTime  string   `json:"modifiedTime,omitempty" yaml:"modifiedTime,omitempty"`
	PublishedTime string   `json:"publishedTime,omitempty" yaml:"publishedTime,omitempty"`
	URL           string   `json:"url,omitempty" yaml:"url,omitempty"`
}

type ArticleListResponse struct {
	Data []Article `json:"data"`
}

type ArticleCreateRequest struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Summary    string   `json:"summary,omitempty"`
	CategoryID string   `json:"categoryId,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

type ArticleUpdateRequest struct {
	Title      string   `json:"title,omitempty"`
	Content    string   `json:"content,omitempty"`
	Summary    string   `json:"summary,omitempty"`
	CategoryID string   `json:"categoryId,omitempty"`
	Status     string   `json:"status,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

type Category struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	ParentID    string `json:"parentId,omitempty" yaml:"parentId,omitempty"`
	ArticleCount int   `json:"articleCount,omitempty" yaml:"articleCount,omitempty"`
	CreatedTime string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`
}

type CategoryListResponse struct {
	Data []Category `json:"data"`
}