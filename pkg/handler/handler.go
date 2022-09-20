package handler

import (
	"example.com/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		deals := api.Group("/deals")
		{
			deals.POST("/", h.createUnclaimedDeal)
			deals.POST("/claim", h.createDeal)
			deals.GET("/", h.getAllDeals)
			deals.GET("/:id", h.getDealById)
			deals.GET("/user/:id", h.getDealsByUserId)
			deals.PUT("/:id", h.updateDeal)
			deals.DELETE("/:id", h.deleteDeal)
		}
		station := api.Group("/station")
		{
			station.POST("/make", h.VerifyDeal)
		}
	}

	admin := router.Group("/admin")
	{
		admin.POST("/get-users", h.getUsers)
	}

	return router
}
