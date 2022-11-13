package data_test

import (
	"testing"

	"github.com/Jason5Lee/simple-blog/core/data"
	"github.com/Jason5Lee/simple-blog/core/errors"
	"github.com/stretchr/testify/assert"
)

func Test_EmptyTitle(t *testing.T) {
	_, err := data.NewArticleTitle("")
	assert.Equal(t, errors.ErrTitleEmpty, err)
}

func Test_LongTitle(t *testing.T) {
	longTitle := make([]byte, data.MAX_ARTICLE_TITLE_LENGTH+1)
	_, err := data.NewArticleTitle(string(longTitle))
	assert.Equal(t, errors.ErrTitleTooLong, err)
}

func Test_EmptyContent(t *testing.T) {
	_, err := data.NewArticleContent("")
	assert.Equal(t, errors.ErrContentEmpty, err)
}

func Test_LongContent(t *testing.T) {
	longContent := make([]byte, data.MAX_ARTICLE_CONTENT_LENGTH+1)
	_, err := data.NewArticleContent(string(longContent))
	assert.Equal(t, errors.ErrContentTooLong, err)
}

func Test_EmptyAuthor(t *testing.T) {
	_, err := data.NewArticleAuthor("")
	assert.Equal(t, errors.ErrAuthorEmpty, err)
}

func Test_LongAuthor(t *testing.T) {
	longAuthor := make([]byte, data.MAX_ARTICLE_AUTHOR_LENGTH+1)
	_, err := data.NewArticleAuthor(string(longAuthor))
	assert.Equal(t, errors.ErrAuthorTooLong, err)
}
