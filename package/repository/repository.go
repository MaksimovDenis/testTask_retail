package repository

import (
	retail "testTask_retail"

	"github.com/jmoiron/sqlx"
)

type OrderDetails interface {
	GetOrderDetails() ([]retail.OrderDetails, error)
}

type Repository struct {
	OrderDetails
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrderDetails: NewRetailPostgres(db),
	}
}
