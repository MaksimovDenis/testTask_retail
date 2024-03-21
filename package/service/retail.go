package service

import (
	retail "testTask_retail"
	"testTask_retail/package/repository"
)

type OrderDetailsService struct {
	repo repository.OrderDetails
}

func NewOrderDetailService(repo repository.OrderDetails) *OrderDetailsService {
	return &OrderDetailsService{repo: repo}
}

func (r *OrderDetailsService) GetOrderDetails() ([]retail.OrderDetails, error) {
	return r.repo.GetOrderDetails()
}
