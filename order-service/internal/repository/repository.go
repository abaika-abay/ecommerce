package repository

import "order-service/internal/domain"

type OrderRepository interface {
	Create(order *domain.Order) error
	FindByID(id string) (*domain.Order, error)
	UpdateStatus(id string, status domain.OrderStatus) error
	ListByUser(userID string, page, limit int) ([]*domain.Order, int, error)
}
