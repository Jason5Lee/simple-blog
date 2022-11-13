package main

import (
	"github.com/Jason5Lee/simple-blog/infra"
	infra_repository "github.com/Jason5Lee/simple-blog/infra/repository"
)

func main() {
	config, err := infra.LoadConfig()
	if err != nil {
		panic(err)
	}
	repo, err := infra_repository.NewArticleRepositoryMongoDB(config.MongoDBUri)
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	err = infra.StartHttpServer(repo, config.Listen)
	if err != nil {
		panic(err)
	}
}
