package handler

import (
	"context"

	"github.com/ArdiDev1/Order-Management-System/services/common/genproto/orders/v1"
	ordertypes "github.com/ArdiDev1/Order-Management-System/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService ordertypes.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, serv ordertypes.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: serv,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId:    42,
		CustomerId: 5,
		ProductId:  3,
		Quantity:   1,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "Success",
	}

	return res, nil
}
