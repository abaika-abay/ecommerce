package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/abaika-abay/ecommerce/protos/inventory"
	"google.golang.org/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	"inventory-service/internal/usecase"
)

func main() {
	// Initialize database connection
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	productRepo := repository.NewProductRepository(db)

	// Initialize usecase
	productUsecase := usecase.NewProductUsecase(productRepo)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	inventoryServer := service.NewInventoryServer(productUsecase)
	inventory.RegisterInventoryServiceServer(grpcServer, inventoryServer)

	// Start server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Inventory service started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
