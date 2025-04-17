package app

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"
	"api-gateway/internal/server"
	"google.golang.org/grpc"
)

type App struct {
	server *server.Server
}

func NewApp() *App {
	s := server.NewServer()

	// Initialize controllers with gRPC connections
	inventoryController := controller.NewInventoryController(s.InventoryConn)
	orderController := controller.NewOrderController(s.OrderConn)

	// Setup routes with middleware
	router := s.GinEngine
	router.Use(middleware.AuthMiddleware())

	// Inventory routes
	inventory := router.Group("/inventory")
	{
		inventory.GET("/products", inventoryController.ListProducts)
		inventory.GET("/products/:id", inventoryController.GetProduct)
		// Add other inventory routes
	}

	// Order routes
	orders := router.Group("/orders")
	{
		orders.POST("/", orderController.CreateOrder)
		orders.GET("/:id", orderController.GetOrder)
		// Add other order routes
	}

	return &App{server: s}
}

func (a *App) Start() error {
	return a.server.Start()
}
