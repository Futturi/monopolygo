package handler

import (
	"errors"

	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
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

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return
	}

	userId, err := h.service.Authentification.ParseToken(headerParts[1])
	if err != nil {
		return
	}
	c.Set(userCtx, userId)
}
