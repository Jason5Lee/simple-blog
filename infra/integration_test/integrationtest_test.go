//go:build integration
// +build integration

package integrationtest_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Jason5Lee/simple-blog/core/repository"
	"github.com/Jason5Lee/simple-blog/infra"
	infra_repository "github.com/Jason5Lee/simple-blog/infra/repository"
	"github.com/stretchr/testify/suite"
)

type integrationTestSuite struct {
	suite.Suite
	port       int
	repo       repository.ArticleRepository
	httpClient *http.Client
	onTearDown func()
}

// Making a request to the testing server.
func (s *integrationTestSuite) request(method string, path string, body string, resp interface{}) error {
	req, err := http.NewRequest(method, fmt.Sprintf("http://localhost:%d%s", s.port, path), strings.NewReader(body))
	if err != nil {
		return err
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	response, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	respText, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respText, resp)
	return err
}

func TestIntegration(t *testing.T) {
	suite.Run(t, &integrationTestSuite{})
}

func (s *integrationTestSuite) SetupSuite() {
	config, err := infra.LoadConfig()
	s.Require().NoError(err)
	repo, err := infra_repository.NewArticleRepositoryMongoDB(config.MongoDBUri)
	s.Require().NoError(err)

	s.port = 8080
	s.repo = repo
	err = repo.Drop()
	s.Require().NoError(err)

	s.onTearDown = func() {
		_ = repo.Drop()
		_ = repo.Close()
	}
	go infra.StartHttpServer(repo, "localhost:8080")
	s.httpClient = &http.Client{}

	// Wait for the http server to start.
	time.Sleep(1 * time.Second)
}

func (s *integrationTestSuite) TearDownSuite() {
	if s.onTearDown != nil {
		s.onTearDown()
	}
}

type ErrorResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    *int   `json:"data"`
}

func (s *integrationTestSuite) Test_CreateArticle_Invalid() {
	s.Run("NoTitle", func() {
		resp := ErrorResp{}
		err := s.request("POST", "/articles", `{"content": "content", "author": "author"}`, &resp)
		s.Require().NoError(err)
		s.Equal(400, resp.Status)
		s.Equal("title is required", resp.Message)
		s.Nil(resp.Data)
	})
	s.Run("NoContent", func() {
		resp := ErrorResp{}
		err := s.request("POST", "/articles", `{"title": "title", "author": "author"}`, &resp)
		s.Require().NoError(err)
		s.Equal(400, resp.Status)
		s.Equal("content is required", resp.Message)
		s.Nil(resp.Data)
	})
	s.Run("NoAuthor", func() {
		resp := ErrorResp{}
		err := s.request("POST", "/articles", `{"title": "title", "content": "content"}`, &resp)
		s.Require().NoError(err)
		s.Equal(400, resp.Status)
		s.Equal("author is required", resp.Message)
		s.Nil(resp.Data)
	})
	s.Run("EmptyContent", func() {
		resp := ErrorResp{}
		err := s.request("POST", "/articles", `{"title": "title", "content": "", "author": "author"}`, &resp)
		s.Require().NoError(err)
		s.Equal(400, resp.Status)
		s.Equal("content is empty", resp.Message)
		s.Nil(resp.Data)
	})
	s.Run("EmptyAuthor", func() {
		resp := ErrorResp{}
		err := s.request("POST", "/articles", `{"title": "title", "content": "content", "author": ""}`, &resp)
		s.Require().NoError(err)
		s.Equal(400, resp.Status)
		s.Equal("author is empty", resp.Message)
		s.Nil(resp.Data)
	})
	s.Run("EmptyTitle", func() {
		resp := ErrorResp{}
		err := s.request("POST", "/articles", `{"title": "", "content": "content", "author": "author"}`, &resp)
		s.Require().NoError(err)
		s.Equal(400, resp.Status)
		s.Equal("title is empty", resp.Message)
		s.Nil(resp.Data)
	})
}

type CreateArticleResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID string `json:"id"`
	}
}
type GetArticleResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
	} `json:"data"`
}

func (s *integrationTestSuite) Test_Two_Articles() {
	createResp := CreateArticleResp{}
	err := s.request("POST", "/articles", `{"title": "title1", "content": "content1", "author": "author1"}`, &createResp)
	s.Require().NoError(err)
	s.Equal(201, createResp.Status)
	s.Equal("Success", createResp.Message)
	s.NotEmpty(createResp.Data.ID)
	id1 := createResp.Data.ID

	createResp = CreateArticleResp{}
	err = s.request("POST", "/articles", `{"title": "title2", "content": "content2", "author": "author2"}`, &createResp)
	s.Require().NoError(err)
	s.Equal(201, createResp.Status)
	s.Equal("Success", createResp.Message)
	s.NotEmpty(createResp.Data.ID)
	id2 := createResp.Data.ID

	getResp := GetArticleResp{}
	err = s.request("GET", "/articles/"+id1, "", &getResp)
	s.Require().NoError(err)
	s.Equal(200, getResp.Status)
	s.Equal("Success", getResp.Message)
	s.Equal(1, len(getResp.Data))
	s.Equal(id1, getResp.Data[0].ID)
	s.Equal("title1", getResp.Data[0].Title)
	s.Equal("content1", getResp.Data[0].Content)
	s.Equal("author1", getResp.Data[0].Author)

	getResp = GetArticleResp{}
	err = s.request("GET", "/articles/"+id2, "", &getResp)
	s.Require().NoError(err)
	s.Equal(200, getResp.Status)
	s.Equal("Success", getResp.Message)
	s.Equal(1, len(getResp.Data))
	s.Equal(id2, getResp.Data[0].ID)
	s.Equal("title2", getResp.Data[0].Title)
	s.Equal("content2", getResp.Data[0].Content)
	s.Equal("author2", getResp.Data[0].Author)

	getResp = GetArticleResp{}
	err = s.request("GET", "/articles/"+id1+id2, "", &getResp)
	s.Require().NoError(err)
	s.Equal(404, getResp.Status)
	s.Equal("article not found", getResp.Message)
	s.Nil(getResp.Data)

	getResp = GetArticleResp{}
	err = s.request("GET", "/articles", "", &getResp)
	s.Require().NoError(err)
	s.Equal(200, getResp.Status)
	s.Equal("Success", getResp.Message)
	s.Equal(2, len(getResp.Data))

	// Get all articles doesn't guarantee the order of articles
	if getResp.Data[0].ID == id2 {
		getResp.Data[0], getResp.Data[1] = getResp.Data[1], getResp.Data[0]
	}
	s.Equal(id1, getResp.Data[0].ID)
	s.Equal("title1", getResp.Data[0].Title)
	s.Equal("content1", getResp.Data[0].Content)
	s.Equal("author1", getResp.Data[0].Author)
	s.Equal(id2, getResp.Data[1].ID)
	s.Equal("title2", getResp.Data[1].Title)
	s.Equal("content2", getResp.Data[1].Content)
	s.Equal("author2", getResp.Data[1].Author)
}
