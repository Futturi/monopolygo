package handler

import (
	"awesomeProject/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var (
	m = melody.New()
)

func (h *Handler) AllServers(c *gin.Context) {
	servers, err := h.service.AllServers()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, servers)
}

func (h *Handler) GetServerById(c *gin.Context) {
	var room models.Room
	id_room, err := getRoomId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	room, err = h.service.GetServerById(id_room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"server": room,
	})
}

func (h *Handler) CreateServer(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	room_id, err := h.service.CreateServer(user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusFound, "/room/"+strconv.Itoa(room_id))
}

func (h *Handler) Connect(c *gin.Context) {
	var user models.User
	room_id, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	room, err := h.service.Connect(room_id, user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err.Error(),
			"user_id": user,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"room": room,
	})

	m.BroadcastBinary([]byte(user.Username + " joined the room"))

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	c.Redirect(http.StatusFound, "/room/"+c.Param("room_id"))
}

func (h *Handler) Disconnect(c *gin.Context) {
	var username string

	room_id, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.BindJSON(&username)

	room, err := h.service.Disconnect(room_id, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, room)

}
