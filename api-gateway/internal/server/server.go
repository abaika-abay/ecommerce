package server

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	ginEngine     *gin.Engine
	inventoryConn *grpc.ClientConn
	orderConn     *grpc.ClientConn
}

func NewServer() *Server {
	return &Server{
		ginEngine: gin.Default(),
	}
}

func (s *Server) Start() error {
	if err := s.initGRPCClients(); err != nil {
		return err
	}

	s.setupRoutes()

	return s.ginEngine.Run(":8080")
}

func (s *Server) initGRPCClients() error {
	var err error

	// Initialize inventory service connection
	s.inventoryConn, err = grpc.Dial(
		"inventory-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	// Initialize order service connection
	s.orderConn, err = grpc.Dial(
		"order-service:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	return err
}

func (s *Server) setupRoutes() {
	api := s.ginEngine.Group("/api")

	// Inventory routes
	inventory := api.Group("/inventory")
	{
		inventory.GET("/products", s.listProducts)
		inventory.GET("/products/:id", s.getProduct)
		inventory.POST("/products", s.createProduct)
		inventory.PUT("/products/:id", s.updateProduct)
		inventory.DELETE("/products/:id", s.deleteProduct)
	}

	// Order routes
	orders := api.Group("/orders")
	{
		orders.POST("/", s.createOrder)
		orders.GET("/:id", s.getOrder)
		orders.PUT("/:id/status", s.updateOrderStatus)
		orders.GET("/user/:user_id", s.listUserOrders)
	}
}
