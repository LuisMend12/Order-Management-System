package service

import (
	"context"

	"github.com/ArdiDev1/Order-Management-System/services/common/genproto/orders/v1"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	//store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	order.OrderId = int32(len(ordersDb) + 1)
	ordersDb = append(ordersDb, order)
	return nil
}
