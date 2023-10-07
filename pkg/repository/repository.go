package repository

import (
	"github.com/MrClean-code/wbtech"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type OrderList interface {
	CreateOrder(order wbtech.Order) (int, error)
	GetOrders(c *gin.Context) ([]wbtech.Order, error)
}

type Repository struct {
	OrderList
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		OrderList: NewOrderPostgres(db),
	}
}
