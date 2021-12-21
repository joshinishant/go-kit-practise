package orderservice

import (
	"go-kit/orderrepository/entity"

	"golang.org/x/net/context"
)

type OderService interface {
	Create(ctx context.Context, order entity.Order) (string, error)
	GetById(ctx context.Context, id string) (entity.Order, error)
	ChangeStatus(ctx context.Context, id, status string) error
}
