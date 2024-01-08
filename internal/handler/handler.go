package handler

import (
	"awesomeProject/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(Service *service.Service) *Handler {
	return &Handler{service: Service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.Login)
	}
	//hub := router.Group("/hub")
	//{
	//hub.GET("/", h.AllServers)
	//hub.GET("/:id", h.GetServerById)
	//hub.POST("/", h.CreateServer)
	//}
	return router
}
