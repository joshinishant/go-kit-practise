package orderrepository

import (
	"context"
	"go-kit/orderrepository/entity"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) (string, error)
	GetOrderById(ctx context.Context, id string) (entity.Order, error)
	ChangeOrderStatus(ctx context.Context, id, status string) error
}
