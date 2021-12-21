package orderrepository

import (
	"context"
	"errors"
	"go-kit/orderrepository/entity"
	"go-kit/storage"
	"log"
)

type OrderRepoImpl struct {
	log log.Logger
}

func (r *OrderRepoImpl) CreateOrder(ctx context.Context, order entity.Order) (string, error) {
	index := len(storage.OrderStore) + 1

	if index >= storage.OrderStoreMaxSize {
		return order.Id, errors.New("Storage Full!!")
	}

	if storage.OrderStore[order.Id] != nil {
		return order.Id, errors.New("Duplicate Order!!")
	}

	storage.OrderStore[order.Id] = order

	return order.Id, nil
}

func (r *OrderRepoImpl) GetOrderById(ctx context.Context, id string) (entity.Order, error) {
	if order := storage.OrderStore[id]; order != nil {
		return order, nil
	} else {
		return entity.Order{}, errors.New("Order doesnt exist!!!")
	}

}

func (r *OrderRepoImpl) ChangeOrderStatus(ctx context.Context, id, status string) error {

	var value entity.Order
	var found bool

	if value, found = storage.OrderStore[id]; found {
		value.Status = status
		return nil
	} else {
		return errors.New("Order doesnt exist!!!")
	}
}
