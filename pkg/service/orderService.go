package service

import (
	"fmt"
	"github.com/MrClean-code/wbtech"
	"github.com/MrClean-code/wbtech/pkg/repository"
	"github.com/gin-gonic/gin"
)

type OrderListPostgres struct {
	repos repository.OrderList
	cache *repository.Cache
}

func NewPostListPostgres(repos repository.OrderList) *OrderListPostgres {
	return &OrderListPostgres{
		repos: repos,
		cache: &repository.Cache{Data: make(map[string]interface{})},
	}
}

func (s *OrderListPostgres) GetOrderByID(id int) (wbtech.Order, error) {
	cachedData, found := s.cache.Get(fmt.Sprintf("order_%d", id))
	if found {
		if order, ok := cachedData.(wbtech.Order); ok {
			return order, nil
		}
	}

	order, err := s.repos.GetOrderByID(id)
	if err != nil {
		return wbtech.Order{}, err
	}

	s.cache.Add(fmt.Sprintf("order_%d", id), order)

	return order, nil
}

func (p *OrderListPostgres) CreateOrders(order wbtech.Order) (int, error) {
	return p.repos.CreateOrder(order)
}

func (p *OrderListPostgres) GetOrderAll(с *gin.Context) ([]wbtech.Order, error) {
	return p.repos.GetOrders(с)
}
