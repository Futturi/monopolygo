package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("error")
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("bad id")
	}
	return idInt, nil
}
