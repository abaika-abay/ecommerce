package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/ecommerce/protos/order"
	"google.golang.org/grpc"
)

type OrderController struct {
	client order.OrderServiceClient
}

func NewOrderController(conn *grpc.ClientConn) *OrderController {
	return &OrderController{
		client: order.NewOrderServiceClient(conn),
	}
}

// CreateOrder handles HTTP POST /orders
// Corresponds to: rpc CreateOrder(CreateOrderRequest) returns (OrderResponse)
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req order.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.client.CreateOrder(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// GetOrder handles HTTP GET /orders/:id
// Corresponds to: rpc GetOrderByID(GetOrderRequest) returns (OrderResponse)
func (c *OrderController) GetOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.client.GetOrderByID(ctx.Request.Context(), &order.GetOrderRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateOrderStatus handles HTTP PUT /orders/:id/status
// Corresponds to: rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (OrderResponse)
func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var reqBody struct {
		Status string `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.client.UpdateOrderStatus(ctx.Request.Context(), &order.UpdateOrderStatusRequest{
		Id:     id,
		Status: reqBody.Status,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListUserOrders handles HTTP GET /users/:user_id/orders?page=X&limit=Y
// Corresponds to: rpc ListUserOrders(ListOrdersRequest) returns (ListOrdersResponse)
func (c *OrderController) ListUserOrders(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	res, err := c.client.ListUserOrders(ctx.Request.Context(), &order.ListOrdersRequest{
		UserId: userID,
		Page:   int32(page),
		Limit:  int32(limit),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
