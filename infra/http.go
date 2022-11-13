package infra

import (
	"github.com/Jason5Lee/simple-blog/core/repository"
	"github.com/Jason5Lee/simple-blog/infra/controller"
	"github.com/gin-gonic/gin"
)

// StartHttpServer starts the HTTP server.
func StartHttpServer(articleRepo repository.ArticleRepository, listenAddr string) error {
	r := gin.Default()
	r.POST("/articles", controller.NewCreateArticleController(articleRepo))
	r.GET("/articles/:article_id", controller.NewGetArticleByIDController(articleRepo))
	r.GET("/articles", controller.NewGetAllArticlesController(articleRepo))
	return r.Run(listenAddr)
}
