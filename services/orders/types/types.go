package types

import (
	"context"

	"github.com/ArdiDev1/Order-Management-System/services/common/genproto/orders/v1"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
