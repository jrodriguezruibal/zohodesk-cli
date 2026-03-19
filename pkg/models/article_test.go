package models

import (
	"encoding/json"
	"testing"
)

func TestArticleJSON(t *testing.T) {
	article := Article{
		ID:           "123",
		Title:        "How to reset password",
		Content:      "Step by step guide...",
		Summary:      "Password reset guide",
		CategoryID:   "456",
		CategoryName: "Help",
		AuthorID:     "789",
		AuthorName:   "John Doe",
		Status:       "Published",
		Tags:         []string{"password", "account"},
		ViewCount:    "100",
		LikeCount:    "10",
		CreatedTime:  "2024-01-15T10:30:00Z",
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		t.Errorf("Failed to marshal Article: %v", err)
	}

	var unmarshaled Article
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Article: %v", err)
	}

	if unmarshaled.Title != article.Title {
		t.Errorf("Expected Title %s, got %s", article.Title, unmarshaled.Title)
	}
	if unmarshaled.Status != article.Status {
		t.Errorf("Expected Status %s, got %s", article.Status, unmarshaled.Status)
	}
	if len(unmarshaled.Tags) != len(article.Tags) {
		t.Errorf("Expected %d tags, got %d", len(article.Tags), len(unmarshaled.Tags))
	}
}

func TestArticleCreateRequestJSON(t *testing.T) {
	req := ArticleCreateRequest{
		Title:      "New Article",
		Content:    "Article content here",
		Summary:    "Brief summary",
		CategoryID: "123",
		Tags:       []string{"tag1", "tag2"},
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal ArticleCreateRequest: %v", err)
	}

	var unmarshaled ArticleCreateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal ArticleCreateRequest: %v", err)
	}

	if unmarshaled.Title != req.Title {
		t.Errorf("Expected Title %s, got %s", req.Title, unmarshaled.Title)
	}
	if len(unmarshaled.Tags) != len(req.Tags) {
		t.Errorf("Expected %d tags, got %d", len(req.Tags), len(unmarshaled.Tags))
	}
}

func TestArticleUpdateRequestJSON(t *testing.T) {
	req := ArticleUpdateRequest{
		Title:   "Updated Title",
		Status:  "Draft",
		Content: "Updated content",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal ArticleUpdateRequest: %v", err)
	}

	var unmarshaled ArticleUpdateRequest
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal ArticleUpdateRequest: %v", err)
	}

	if unmarshaled.Title != req.Title {
		t.Errorf("Expected Title %s, got %s", req.Title, unmarshaled.Title)
	}
	if unmarshaled.Status != req.Status {
		t.Errorf("Expected Status %s, got %s", req.Status, unmarshaled.Status)
	}
}

func TestCategoryJSON(t *testing.T) {
	category := Category{
		ID:           "123",
		Name:         "Getting Started",
		Description:  "Articles for new users",
		ParentID:     "",
		ArticleCount: 25,
		CreatedTime:  "2024-01-01T00:00:00Z",
	}

	jsonData, err := json.Marshal(category)
	if err != nil {
		t.Errorf("Failed to marshal Category: %v", err)
	}

	var unmarshaled Category
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal Category: %v", err)
	}

	if unmarshaled.Name != category.Name {
		t.Errorf("Expected Name %s, got %s", category.Name, unmarshaled.Name)
	}
	if unmarshaled.ArticleCount != category.ArticleCount {
		t.Errorf("Expected ArticleCount %d, got %d", category.ArticleCount, unmarshaled.ArticleCount)
	}
}

func TestArticleListResponseJSON(t *testing.T) {
	response := ArticleListResponse{
		Data: []Article{
			{ID: "1", Title: "Article 1", Status: "Published"},
			{ID: "2", Title: "Article 2", Status: "Draft"},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal ArticleListResponse: %v", err)
	}

	var unmarshaled ArticleListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal ArticleListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d articles, got %d", len(response.Data), len(unmarshaled.Data))
	}
	if unmarshaled.Data[0].Title != response.Data[0].Title {
		t.Errorf("Expected first article title %s, got %s", response.Data[0].Title, unmarshaled.Data[0].Title)
	}
}

func TestCategoryListResponseJSON(t *testing.T) {
	response := CategoryListResponse{
		Data: []Category{
			{ID: "1", Name: "Category 1", ArticleCount: 10},
			{ID: "2", Name: "Category 2", ArticleCount: 20},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal CategoryListResponse: %v", err)
	}

	var unmarshaled CategoryListResponse
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Errorf("Failed to unmarshal CategoryListResponse: %v", err)
	}

	if len(unmarshaled.Data) != len(response.Data) {
		t.Errorf("Expected %d categories, got %d", len(response.Data), len(unmarshaled.Data))
	}
}