package repository

import (
	"context"

	"github.com/Jason5Lee/simple-blog/core/data"
)

type ArticleRepository interface {
	// Create creates a new article.
	Create(ctx context.Context, article *data.ArticleInfo) (data.ArticleID, error)
	// GetByID gets an article by ID.
	GetByID(ctx context.Context, id data.ArticleID) (*data.Article, error)
	// GetAll gets all articles.
	GetAll(ctx context.Context) ([]*data.Article, error)
}
