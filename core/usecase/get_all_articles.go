package usecase

import (
	"context"

	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/repository"
)

// GetAllArticles gets all articles.
func GetAllArticles(ctx context.Context, repo repository.ArticleRepository) ([]*data.Article, error) {
	return repo.GetAll(ctx)
}
