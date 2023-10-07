package handler

import (
	"github.com/MrClean-code/wbtech/pkg/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createOrder(c *gin.Context) {

}

func (h *Handler) getAllOrders(c *gin.Context) {
	orders, err := h.services.GetOrderAll(c)
	if err != nil {
		exception.NewErrorResponse(c, http.StatusInternalServerError, "Failed to get orders")
		return
	}
	c.JSON(http.StatusOK, orders) // Возвращаем заказы в формате JSON
}
