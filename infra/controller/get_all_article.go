package controller

import (
	"github.com/Jason5Lee/simple-blog/core/repository"
	"github.com/Jason5Lee/simple-blog/core/usecase"
	"github.com/gin-gonic/gin"
)

// NewGetAllArticlesController creates a new controller for getting all articles.
func NewGetAllArticlesController(articleRepo repository.ArticleRepository) func(*gin.Context) {
	return func(c *gin.Context) {
		article, err := usecase.GetAllArticles(c, articleRepo)
		if err != nil {
			respondErr(c, err)
			return
		}
		response := make([]gin.H, len(article))
		for i, a := range article {
			response[i] = gin.H{
				"id":      a.ID,
				"title":   string(a.Title),
				"content": string(a.Content),
				"author":  string(a.Author),
			}
		}
		respond(c, 200, "Success", response)
	}
}
