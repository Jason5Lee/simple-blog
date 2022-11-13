package repository

import (
	"context"
	"fmt"

	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/errors"
	"github.com/Jason5Lee/simple-blog/core/repository"
)

// In-memory implementation of ArticleRepository.
type ArticleRepositoryInMemory struct {
	articles map[data.ArticleID]*data.ArticleInfo
}

func NewArticleRepositoryInMemory() *ArticleRepositoryInMemory {
	return &ArticleRepositoryInMemory{
		articles: make(map[data.ArticleID]*data.ArticleInfo),
	}
}

func (r *ArticleRepositoryInMemory) Create(ctx context.Context, article *data.ArticleInfo) (data.ArticleID, error) {
	id := data.ArticleID(fmt.Sprint(len(r.articles) + 1))
	r.articles[id] = article
	return id, nil
}

func (r *ArticleRepositoryInMemory) GetByID(ctx context.Context, id data.ArticleID) (*data.Article, error) {
	article, ok := r.articles[id]
	if !ok {
		return nil, errors.ErrNotFound
	}
	return &data.Article{
		ID:          id,
		ArticleInfo: *article,
	}, nil
}

func (r *ArticleRepositoryInMemory) GetAll(ctx context.Context) ([]*data.Article, error) {
	articles := make([]*data.Article, 0, len(r.articles))
	for id, article := range r.articles {
		articles = append(articles, &data.Article{
			ID:          id,
			ArticleInfo: *article,
		})
	}
	return articles, nil
}

var _ repository.ArticleRepository = (*ArticleRepositoryInMemory)(nil)
