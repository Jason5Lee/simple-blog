package errors

import "errors"

var ErrNotFound = errors.New("article not found")
var ErrTitleEmpty = errors.New("title is empty")
var ErrContentEmpty = errors.New("content is empty")
var ErrAuthorEmpty = errors.New("author is empty")
var ErrTitleTooLong = errors.New("title is too long")
var ErrContentTooLong = errors.New("content is too long")
var ErrAuthorTooLong = errors.New("author is too long")
