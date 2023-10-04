package service

import (
	"github.com/MrClean-code/wbtech"
	"github.com/MrClean-code/wbtech/pkg/repository"
)

type OrderList interface {
	CreateOrders(order wbtech.Order) (int, error)
	GetOrderAll() ([]wbtech.Order, error)
}

type Service struct {
	OrderList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		OrderList: NewPostListPostgres(repos.OrderList),
	}
}
