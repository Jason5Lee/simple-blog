package data

import "github.com/Jason5Lee/simple-blog/core/errors"

// Use type definition to represent the validated value.
type ArticleTitle string
type ArticleContent string
type ArticleAuthor string

type ArticleInfo struct {
	Title   ArticleTitle
	Content ArticleContent
	Author  ArticleAuthor
}

type ArticleID string

type Article struct {
	ID ArticleID
	ArticleInfo
}

const MAX_ARTICLE_TITLE_LENGTH = 1024
const MAX_ARTICLE_CONTENT_LENGTH = 4 * 1024 * 1024 // 4MB
const MAX_ARTICLE_AUTHOR_LENGTH = 1024

// NewArticleTitle returns a new ArticleTitle if the title is valid.
func NewArticleTitle(title string) (ArticleTitle, error) {
	if title == "" {
		return "", errors.ErrTitleEmpty
	}
	if len(title) > MAX_ARTICLE_TITLE_LENGTH {
		return "", errors.ErrTitleTooLong
	}
	return ArticleTitle(title), nil
}

// NewArticleContent returns a new ArticleContent if the content is valid.
func NewArticleContent(content string) (ArticleContent, error) {
	if content == "" {
		return "", errors.ErrContentEmpty
	}
	if len(content) > MAX_ARTICLE_CONTENT_LENGTH {
		return "", errors.ErrContentTooLong
	}
	return ArticleContent(content), nil
}

// NewArticleAuthor returns a new ArticleAuthor if the author is valid.
func NewArticleAuthor(author string) (ArticleAuthor, error) {
	if author == "" {
		return "", errors.ErrAuthorEmpty
	}
	if len(author) > MAX_ARTICLE_AUTHOR_LENGTH {
		return "", errors.ErrAuthorTooLong
	}
	return ArticleAuthor(author), nil
}
