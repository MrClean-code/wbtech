package service

import (
	"fmt"
	"github.com/MrClean-code/wbtech"
	"github.com/MrClean-code/wbtech/pkg/repository"
	"github.com/gin-gonic/gin"
)

type OrderListPostgres struct {
	repos repository.OrderList
}

func NewPostListPostgres(repos repository.OrderList) *OrderListPostgres {
	fmt.Println("NewPostListPostgres 111")
	return &OrderListPostgres{repos: repos}
}

func (p *OrderListPostgres) CreateOrders(order wbtech.Order) (int, error) {
	return p.repos.CreateOrder(order)
}

func (p *OrderListPostgres) GetOrderAll(с *gin.Context) ([]wbtech.Order, error) {
	return p.repos.GetOrders(с)
}
