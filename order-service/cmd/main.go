package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/yourusername/ecommerce/protos/order"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/internal/usecase"
)

func main() {
	// Initialize MongoDB connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("ecommerce")

	// Initialize repository
	orderRepo := repository.NewOrderRepository(db)

	// Initialize usecase
	orderUsecase := usecase.NewOrderUsecase(orderRepo)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	orderServer := service.NewOrderServer(orderUsecase)
	order.RegisterOrderServiceServer(grpcServer, orderServer)

	// Start server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Order service started on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
