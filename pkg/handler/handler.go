package handler

import (
	"fmt"
	"github.com/MrClean-code/wbtech/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	fmt.Println("NewHandler 111")
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	api := routes.Group("/api")
	{
		api.POST("/new/order", h.createOrder)
		api.GET("/orders", h.getAllOrders)
		logrus.Print("handlers working")
	}
	return routes
}
