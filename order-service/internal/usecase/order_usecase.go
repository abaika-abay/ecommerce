package usecase

import (
	"context"
	"errors"

	"order-service/internal/domain"
	"order-service/internal/repository"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrder(ctx context.Context, id string) (*domain.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status domain.OrderStatus) (*domain.Order, error)
	ListUserOrders(ctx context.Context, userID string, page, limit int) ([]*domain.Order, int, error)
}

type orderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) OrderUsecase {
	return &orderUsecase{repo: repo}
}

// CreateOrder processes order creation.
// Corresponds to: rpc CreateOrder(CreateOrderRequest) returns (OrderResponse)
func (uc *orderUsecase) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	// Calculate total
	var total float64
	for _, item := range order.Items {
		total += item.Price * float64(item.Quantity)
	}
	order.Total = total
	order.Status = domain.OrderStatusPending

	if err := uc.repo.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrder fetches an order by ID.
// Corresponds to: rpc GetOrderByID(GetOrderRequest) returns (OrderResponse)
func (uc *orderUsecase) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	order, err := uc.repo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// UpdateOrderStatus changes the status of an existing order.
// Corresponds to: rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (OrderResponse)
func (uc *orderUsecase) UpdateOrderStatus(ctx context.Context, id string, status domain.OrderStatus) (*domain.Order, error) {
	order, err := uc.repo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	// Optional: validate transition rules here
	if order.Status == status {
		return order, nil
	}

	if err := uc.repo.UpdateOrderStatus(id, string(status)); err != nil {
		return nil, err
	}

	order.Status = status
	return order, nil
}

// ListUserOrders returns paginated orders of a specific user.
// Corresponds to: rpc ListUserOrders(ListOrdersRequest) returns (ListOrdersResponse)
func (uc *orderUsecase) ListUserOrders(ctx context.Context, userID string, page, limit int) ([]*domain.Order, int, error) {
	orders, err := uc.repo.ListUserOrders(userID)
	if err != nil {
		return nil, 0, err
	}

	total := len(orders)
	start := (page - 1) * limit
	if start > total {
		return []*domain.Order{}, total, nil
	}

	end := start + limit
	if end > total {
		end = total
	}

	return orders[start:end], total, nil
}
