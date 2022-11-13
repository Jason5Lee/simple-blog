package usecase

import (
	"context"

	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/repository"
)

// CreateArticle creates a new article.
// Note that each field in the type `ArticleInfo` uses the type definition,
// which means they are validated.
func CreateArticle(ctx context.Context, repo repository.ArticleRepository, article *data.ArticleInfo) (data.ArticleID, error) {
	return repo.Create(ctx, article)
}
