package api

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
)

type ArticlesService struct {
	client *Client
}

func NewArticlesService(client *Client) *ArticlesService {
	return &ArticlesService{client: client}
}

func (s *ArticlesService) List(ctx context.Context, params map[string]string) ([]models.Article, error) {
	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	data, err := s.client.Get(ctx, "/articles", query)
	if err != nil {
		return nil, err
	}

	var resp models.ArticleListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *ArticlesService) Get(ctx context.Context, articleID string) (*models.Article, error) {
	data, err := s.client.Get(ctx, "/articles/"+articleID, nil)
	if err != nil {
		return nil, err
	}

	var modelsarticle models.Article
	if err := unmarshalData(data, &modelsarticle); err != nil {
		return nil, err
	}
	return &modelsarticle, nil
}

func (s *ArticlesService) Search(ctx context.Context, query string) ([]models.Article, error) {
	params := map[string]string{"query": query}
	return s.List(ctx, params)
}

func (s *ArticlesService) Create(ctx context.Context, req models.ArticleCreateRequest) (*models.Article, error) {
	data, err := s.client.Post(ctx, "/articles", req)
	if err != nil {
		return nil, err
	}

	var modelsarticle models.Article
	if err := unmarshalData(data, &modelsarticle); err != nil {
		return nil, err
	}
	return &modelsarticle, nil
}

func (s *ArticlesService) Update(ctx context.Context, articleID string, req models.ArticleUpdateRequest) (*models.Article, error) {
	data, err := s.client.Patch(ctx, "/articles/"+articleID, req)
	if err != nil {
		return nil, err
	}

	var modelsarticle models.Article
	if err := unmarshalData(data, &modelsarticle); err != nil {
		return nil, err
	}
	return &modelsarticle, nil
}

func (s *ArticlesService) Delete(ctx context.Context, articleID string) error {
	return s.client.Delete(ctx, "/articles/"+articleID)
}