package controller

import (
	"github.com/Jason5Lee/simple-blog/core/errors"
	"github.com/gin-gonic/gin"
)

// respond is a helper function to respond.
func respond(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

// getStatusCode gets the status code from error.
func getStatusCode(err error) int {
	switch err {
	case errors.ErrNotFound:
		return 404
	case errors.ErrAuthorEmpty, errors.ErrAuthorTooLong, errors.ErrContentEmpty, errors.ErrContentTooLong, errors.ErrTitleEmpty, errors.ErrTitleTooLong:
		return 400
	}
	return 500
}

// respondErr is a helper function to respond error.
func respondErr(c *gin.Context, err error) {
	respond(c, getStatusCode(err), err.Error(), nil)
}
