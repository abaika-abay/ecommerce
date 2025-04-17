package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"order-service/internal/domain"
)

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) OrderRepository {
	return &orderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *orderRepository) Create(order *domain.Order) error {
	_, err := r.collection.InsertOne(context.Background(), bson.M{
		"id":         order.ID,
		"user_id":    order.UserID,
		"items":      order.Items,
		"total":      order.Total,
		"status":     order.Status,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})
	return err
}

// GetOrderByID retrieves a single order by its ID from the database.
// Corresponds to: rpc GetOrderByID(GetOrderRequest) returns (OrderResponse)
func (r *orderRepository) GetOrderByID(id string) (*domain.Order, error) {
	var order domain.Order
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrderStatus updates the status of an order in the database.
// Corresponds to: rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (OrderResponse)
func (r *orderRepository) UpdateOrderStatus(id string, status string) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"id": id},
		bson.M{
			"$set": bson.M{
				"status":     status,
				"updated_at": time.Now(),
			},
		},
	)
	return err
}

// ListUserOrders fetches all orders for a given user ID.
// Corresponds to: rpc ListUserOrders(ListOrdersRequest) returns (ListOrdersResponse)
func (r *orderRepository) ListUserOrders(userID string) ([]*domain.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []*domain.Order
	for cursor.Next(context.Background()) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
