package handler

import (
	"github.com/MrClean-code/wbtech"
	"github.com/MrClean-code/wbtech/pkg/exception"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createOrder(c *gin.Context) {
	var req wbtech.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		exception.NewErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	orderReq := wbtech.Order{
		TrackNumber:       req.TrackNumber,
		Entry:             req.Entry,
		Locale:            req.Locale,
		InternalSignature: req.InternalSignature,
		CustomerId:        req.CustomerId,
		DeliveryService:   req.DeliveryService,
		ShardKey:          req.ShardKey,
		SmId:              req.SmId,
		DateCreated:       req.DateCreated,
		OofShard:          req.OofShard,
		Delivery:          req.Delivery,
		Payment:           req.Payment,
		Item:              req.Item,
	}

	order, err := h.services.CreateOrders(orderReq)
	if err != nil {
		exception.NewErrorResponse(c, http.StatusInternalServerError, "Failed to get orders")
		return
	}
	c.JSON(http.StatusOK, order)

}

func (h *Handler) getAllOrders(c *gin.Context) {
	orders, err := h.services.GetOrderAll(c)
	if err != nil {
		exception.NewErrorResponse(c, http.StatusInternalServerError, "Failed to get orders")
		return
	}
	c.JSON(http.StatusOK, orders) // Возвращаем заказы в формате JSON
}

func (h *Handler) getOrderByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		exception.NewErrorResponse(c, http.StatusBadRequest, "Invalid 'id' parameter")
		return
	}

	order, err := h.services.GetOrderByID(id)
	if err != nil {
		exception.NewErrorResponse(c, http.StatusInternalServerError, "Failed to get order by ID")
		return
	}

	c.JSON(http.StatusOK, order)
}
