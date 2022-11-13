package controller

import (
	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/repository"
	"github.com/Jason5Lee/simple-blog/core/usecase"
	"github.com/gin-gonic/gin"
)

// NewGetArticleByIDController creates a controller for getting an article by ID.
func NewGetArticleByIDController(articleRepo repository.ArticleRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error

		id := c.Param("article_id")
		if id == "" {
			respond(c, 400, "article_id is required", nil)
			return
		}

		article, err := usecase.GetArticleByID(c, articleRepo, data.ArticleID(id))
		if err != nil {
			respondErr(c, err)
			return
		}
		respond(c, 200, "Success", []gin.H{
			{
				"id":      id,
				"title":   string(article.Title),
				"content": string(article.Content),
				"author":  string(article.Author),
			},
		})
	}
}
