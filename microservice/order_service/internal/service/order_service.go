package service

import "order_service/internal/repository"

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(r *repository.OrderRepository) *OrderService {
	return &OrderService{repo: r}
}

func (s *OrderService) CreateOrder(userId int32, product string) (*repository.Order, error) {
	return s.repo.CreateOrder(userId, product)
}

func (s *OrderService) GetOrder(id int32) (*repository.Order, error) {
	return s.repo.GetOrder(id)
}
