package usecase

import (
	"context"

	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/repository"
)

// GetArticleByID gets an article by ID.
func GetArticleByID(ctx context.Context, repo repository.ArticleRepository, id data.ArticleID) (*data.Article, error) {
	return repo.GetByID(ctx, id)
}
