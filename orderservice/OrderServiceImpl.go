package orderservice

import (
	"context"
	"go-kit/entity"
	"go-kit/orderrepository"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type OrderService struct {
	Log       log.Logger
	OrderRepo orderrepository.OrderRepository
}

func NewService(rep orderrepository.OrderRepository, log log.Logger) *OrderService {
	return &OrderService{Log: log, OrderRepo: rep}
}

func (s *OrderService) Create(ctx context.Context, order entity.Order) (string, error) {
	uuid, _ := uuid.NewV4()
	order.Id = uuid.String()

	order.CreatedOn = time.Now().Unix()
	order.Status = "Pending"

	orderId, error := s.OrderRepo.CreateOrder(ctx, order)
	return orderId, error
}
func (s *OrderService) GetById(ctx context.Context, id string) (entity.Order, error) {
	return s.OrderRepo.GetOrderById(ctx, id)
}
func (s *OrderService) ChangeStatus(ctx context.Context, id, status string) error {
	return s.OrderRepo.ChangeOrderStatus(ctx, id, status)
}
