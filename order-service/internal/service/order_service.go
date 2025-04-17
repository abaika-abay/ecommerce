package service

import (
	"context"
	"time"

	"github.com/yourusername/ecommerce/protos/order"
	"order-service/internal/domain"
	"order-service/internal/usecase"
)

type OrderServer struct {
	order.UnimplementedOrderServiceServer
	orderUsecase usecase.OrderUsecase
}

func NewOrderServer(orderUsecase usecase.OrderUsecase) *OrderServer {
	return &OrderServer{
		orderUsecase: orderUsecase,
	}
}

// CreateOrder handles order creation.
// Corresponds to: rpc CreateOrder(CreateOrderRequest) returns (OrderResponse)
func (s *OrderServer) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.OrderResponse, error) {
	orderItems := make([]domain.OrderItem, len(req.Items))
	for i, item := range req.Items {
		orderItems[i] = domain.OrderItem{
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity),
			Price:     item.Price,
		}
	}

	newOrder := &domain.Order{
		ID:     generateID(), // Implement your ID generation
		UserID: req.UserId,
		Items:  orderItems,
	}

	createdOrder, err := s.orderUsecase.CreateOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}

	return &order.OrderResponse{
		Order: s.domainToProto(createdOrder),
	}, nil
}

// GetOrderByID returns an order by its ID.
// Corresponds to: rpc GetOrderByID(GetOrderRequest) returns (OrderResponse)
func (s *OrderServer) GetOrderByID(ctx context.Context, req *order.GetOrderRequest) (*order.OrderResponse, error) {
	foundOrder, err := s.orderUsecase.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &order.OrderResponse{
		Order: s.domainToProto(foundOrder),
	}, nil
}

// UpdateOrderStatus changes an order's status.
// Corresponds to: rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (OrderResponse)
func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.OrderResponse, error) {
	updatedOrder, err := s.orderUsecase.UpdateOrderStatus(ctx, req.Id, domain.OrderStatus(req.Status))
	if err != nil {
		return nil, err
	}

	return &order.OrderResponse{
		Order: s.domainToProto(updatedOrder),
	}, nil
}

// ListUserOrders returns all orders for a user with pagination.
// Corresponds to: rpc ListUserOrders(ListOrdersRequest) returns (ListOrdersResponse)
func (s *OrderServer) ListUserOrders(ctx context.Context, req *order.ListOrdersRequest) (*order.ListOrdersResponse, error) {
	orders, total, err := s.orderUsecase.ListUserOrders(ctx, req.UserId, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, err
	}

	protoOrders := make([]*order.Order, len(orders))
	for i, o := range orders {
		protoOrders[i] = s.domainToProto(o)
	}

	return &order.ListOrdersResponse{
		Orders: protoOrders,
		Total:  int32(total),
	}, nil
}

// Helper to convert domain.Order to proto.Order
func (s *OrderServer) domainToProto(order *domain.Order) *order.Order {
	items := make([]*order.OrderItem, len(order.Items))
	for i, item := range order.Items {
		items[i] = &order.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		}
	}

	return &order.Order{
		Id:        order.ID,
		UserId:    order.UserID,
		Items:     items,
		Total:     order.Total,
		Status:    string(order.Status),
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
	}
}
