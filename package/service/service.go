package service

import (
	retail "testTask_retail"
	"testTask_retail/package/repository"
)

type OrderDetails interface {
	GetOrderDetails() ([]retail.OrderDetails, error)
}

type Service struct {
	OrderDetails
}

// Service acces database
func NewService(repos *repository.Repository) *Service {
	return &Service{
		OrderDetails: NewOrderDetailService(repos.OrderDetails),
	}
}
