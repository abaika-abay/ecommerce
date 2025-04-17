package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/ecommerce/protos/inventory"
	"google.golang.org/grpc"
)

type InventoryController struct {
	client inventory.InventoryServiceClient
}

func NewInventoryController(conn *grpc.ClientConn) *InventoryController {
	return &InventoryController{
		client: inventory.NewInventoryServiceClient(conn),
	}
}

// ListProducts handles GET /products?page=X&limit=Y
// Corresponds to: rpc ListProducts(ListProductsRequest) returns (ListProductsResponse)
func (c *InventoryController) ListProducts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := &inventory.ListProductsRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	res, err := c.client.ListProducts(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetProduct handles GET /products/:id
// Corresponds to: rpc GetProductByID(GetProductRequest) returns (ProductResponse)
func (c *InventoryController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.client.GetProductByID(ctx.Request.Context(), &inventory.GetProductRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// CreateProduct handles POST /products
// Corresponds to: rpc CreateProduct(CreateProductRequest) returns (ProductResponse)
func (c *InventoryController) CreateProduct(ctx *gin.Context) {
	var req inventory.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.client.CreateProduct(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// UpdateProductStock handles PUT /products/:id/stock
// Corresponds to: rpc UpdateProductStock(UpdateProductStockRequest) returns (ProductResponse)
func (c *InventoryController) UpdateProductStock(ctx *gin.Context) {
	id := ctx.Param("id")
	var reqBody struct {
		Stock int32 `json:"stock"`
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := &inventory.UpdateProductStockRequest{
		Id:    id,
		Stock: reqBody.Stock,
	}

	res, err := c.client.UpdateProductStock(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
