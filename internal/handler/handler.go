package handler

import (
	"awesomeProject/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	cfg     service.ConfigEmail
}

func NewHandler(Service *service.Service, cfg service.ConfigEmail) *Handler {
	return &Handler{service: Service, cfg: cfg}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.Login)
		auth.GET("/:token", h.VerifyEmail)
		auth.POST("/refreshtoken", h.RefreshToken)
	}
	api := router.Group("/api", h.UserIdentity)
	{
		hub := api.Group("/hub")
		{
			hub.GET("/", h.AllServers)
			hub.GET("/:room_id", h.GetServerById)
			hub.POST("/create", h.CreateServer)
		}
	}
	return router
}
