package handler

import (
	"awesomeProject/internal/models"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("error with input: %s", err.Error())
		return
	}
	res, err := h.service.SignUp(input, h.cfg)
	if err != nil {
		log.Fatalf("error while creating user: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": res,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var input models.SignInInput
	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("error with input: %s", err.Error())
	}
	token, err := h.service.Token(input)
	if err != nil {
		log.Fatalf("error while creating token: %s", err.Error())
	}
	refresh, err := h.service.GetRefresh(input)
	c.JSON(http.StatusOK, map[string]interface{}{
		"accesstoken":   token,
		"refresh_token": refresh,
	})
}

func (h *Handler) VerifyEmail(c *gin.Context) {
	token := c.Param("token")
	user, err := h.service.VerifyEmail(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": user,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var token models.RefreshToken
	if err := c.BindJSON(&token); err != nil {
		log.Fatalf("error with token: %s", err.Error())
	}
	result, err := h.service.RefreshToken(token.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": result,
	})
}
