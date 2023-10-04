package repository

import (
	"github.com/MrClean-code/wbtech"
	"github.com/jmoiron/sqlx"
)

type OrderList interface {
	CreateOrder(order wbtech.Order) (int, error)
	GetOrders() ([]wbtech.Order, error)
}

type Repository struct {
	OrderList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrderList: NewOrderPostgres(db),
	}
}
