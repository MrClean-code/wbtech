package service

import (
	"github.com/MrClean-code/wbtech"
	"github.com/MrClean-code/wbtech/pkg/repository"
)

type OrderListPostgres struct {
	repos repository.OrderList
}

func NewPostListPostgres(repos repository.OrderList) *OrderListPostgres {
	return &OrderListPostgres{repos: repos}
}

func (p OrderListPostgres) CreateOrders(order wbtech.Order) (int, error) {
	panic("implement me")
}

func (p OrderListPostgres) GetOrderAll() ([]wbtech.Order, error) {
	return p.repos.GetOrders()
}
