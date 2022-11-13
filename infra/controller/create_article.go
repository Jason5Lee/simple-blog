package controller

import (
	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/repository"
	"github.com/Jason5Lee/simple-blog/core/usecase"
	"github.com/gin-gonic/gin"
)

// CreateArticleRequest is the request body for creating an article.
type CreateArticleRequest struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Author  *string `json:"author"`
}

// NewCreateArticleController creates a new controller for creating an article.
func NewCreateArticleController(articleRepo repository.ArticleRepository) func(*gin.Context) {
	return func(c *gin.Context) {
		var err error
		var req CreateArticleRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			respondErr(c, err)
			return
		}
		if req.Title == nil {
			respond(c, 400, "title is required", nil)
			return
		}
		if req.Content == nil {
			respond(c, 400, "content is required", nil)
			return
		}
		if req.Author == nil {
			respond(c, 400, "author is required", nil)
			return
		}

		article := &data.ArticleInfo{}
		article.Title, err = data.NewArticleTitle(*req.Title)
		if err != nil {
			respondErr(c, err)
			return
		}
		article.Content, err = data.NewArticleContent(*req.Content)
		if err != nil {
			respondErr(c, err)
			return
		}
		article.Author, err = data.NewArticleAuthor(*req.Author)
		if err != nil {
			respondErr(c, err)
			return
		}

		id, err := usecase.CreateArticle(c, articleRepo, article)
		if err != nil {
			respondErr(c, err)
			return
		}
		respond(c, 201, "Success", gin.H{"id": id})
	}
}
