package usecase_test

import (
	"context"
	"sort"
	"testing"

	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/errors"
	"github.com/Jason5Lee/simple-blog/core/usecase"
	infra_repository "github.com/Jason5Lee/simple-blog/infra/repository"
	"github.com/stretchr/testify/assert"
)

// Because the GetAllArticles does not garantee the order,
// we need to sort the articles before comparing.
type articleSorterByAuthor struct {
	articles []*data.Article
}

func (s *articleSorterByAuthor) Len() int {
	return len(s.articles)
}
func (s *articleSorterByAuthor) Swap(i, j int) {
	s.articles[i], s.articles[j] = s.articles[j], s.articles[i]
}
func (s *articleSorterByAuthor) Less(i, j int) bool {
	return s.articles[i].Author < s.articles[j].Author
}

func Test_NoArticle(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	repo := infra_repository.NewArticleRepositoryInMemory()

	_, err := usecase.GetArticleByID(ctx, repo, data.ArticleID("24"))
	assert.Equal(errors.ErrNotFound, err, "get article from empty repo should return ErrNotFound")

	articles, err := usecase.GetAllArticles(ctx, repo)
	assert.Nil(err, "get all articles from empty repo should not return error")
	assert.Empty(articles, "get all articles from empty repo should return empty slice")
}

const testTitle = "Hello World"
const testContent = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
const testAuthor = "John"

func Test_SingleArticle(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	repo := infra_repository.NewArticleRepositoryInMemory()
	articleId, err := usecase.CreateArticle(ctx, repo, &data.ArticleInfo{
		Title:   testTitle,
		Content: testContent,
		Author:  testAuthor,
	})
	assert.Nil(err, "create article should not return error")

	article, err := usecase.GetArticleByID(ctx, repo, articleId)
	assert.Nil(err, "get article from the created ID should not return error")
	assert.Equal(testTitle, string(article.Title), "get article from the created ID should return correct title")
	assert.Equal(testContent, string(article.Content), "get article from the created ID should return correct content")
	assert.Equal(testAuthor, string(article.Author), "get article from the created ID should return correct author")

	_, err = usecase.GetArticleByID(ctx, repo, data.ArticleID(string(articleId)+"e"))
	assert.Equal(errors.ErrNotFound, err, "get article from the wrong ID should return ErrNotFound")

	articles, err := usecase.GetAllArticles(ctx, repo)
	assert.Nil(err, "get all articles should not return error")
	assert.Len(articles, 1, "get all articles should return 1 article")
	assert.Equal(testTitle, string(articles[0].Title), "get all articles should return correct title")
	assert.Equal(testContent, string(articles[0].Content), "get all articles should return correct content")
	assert.Equal(testAuthor, string(articles[0].Author), "get all articles should return correct author")
}

func Test_TwoArticles(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	repo := infra_repository.NewArticleRepositoryInMemory()
	articleId1, err := usecase.CreateArticle(ctx, repo, &data.ArticleInfo{
		Title:   "title 1",
		Content: "content 1",
		Author:  "author 1",
	})
	assert.Nil(err, "create article 1 should not return error")

	articleId2, err := usecase.CreateArticle(ctx, repo, &data.ArticleInfo{
		Title:   "title 2",
		Content: "content 2",
		Author:  "author 2",
	})
	assert.Nil(err, "create article 2 should not return error")

	article, err := usecase.GetArticleByID(ctx, repo, articleId1)
	assert.Nil(err, "get article 1 from the created ID should not return error")
	assert.Equal("title 1", string(article.Title), "get article 1 from the created ID should return correct title")
	assert.Equal("content 1", string(article.Content), "get article 1 from the created ID should return correct content")
	assert.Equal("author 1", string(article.Author), "get article 1 from the created ID should return correct author")

	article, err = usecase.GetArticleByID(ctx, repo, articleId2)
	assert.Nil(err, "get article 2 from the created ID should not return error")
	assert.Equal("title 2", string(article.Title), "get article 2 from the created ID should return correct title")
	assert.Equal("content 2", string(article.Content), "get article 2 from the created ID should return correct content")
	assert.Equal("author 2", string(article.Author), "get article 2 from the created ID should return correct author")

	articles, err := usecase.GetAllArticles(ctx, repo)
	assert.Nil(err, "get all articles should not return error")
	assert.Len(articles, 2, "get all articles should return 2 articles")

	sorter := &articleSorterByAuthor{articles}
	sort.Sort(sorter)
	assert.Equal("title 1", string(sorter.articles[0].Title), "get all articles should return correct title")
	assert.Equal("content 1", string(sorter.articles[0].Content), "get all articles should return correct content")
	assert.Equal("author 1", string(sorter.articles[0].Author), "get all articles should return correct author")
	assert.Equal("title 2", string(sorter.articles[1].Title), "get all articles should return correct title")
	assert.Equal("content 2", string(sorter.articles[1].Content), "get all articles should return correct content")
	assert.Equal("author 2", string(sorter.articles[1].Author), "get all articles should return correct author")
}
