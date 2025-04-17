package service

import (
	"context"
	"errors"
	"time"

	"github.com/abaika-abay/ecommerce/protos/inventory"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
)

type InventoryServer struct {
	inventory.UnimplementedInventoryServiceServer
	productUsecase usecase.ProductUsecase
}

func NewInventoryServer(productUsecase usecase.ProductUsecase) *InventoryServer {
	return &InventoryServer{
		productUsecase: productUsecase,
	}
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *inventory.CreateProductRequest) (*inventory.ProductResponse, error) {
	product := &domain.Product{
		ID:          generateID(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int(req.Stock),
		CategoryID:  req.CategoryId,
	}

	createdProduct, err := s.productUsecase.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &inventory.ProductResponse{
		Product: s.domainToProto(createdProduct),
	}, nil
}

func (s *InventoryServer) GetProductByID(ctx context.Context, req *inventory.GetProductRequest) (*inventory.ProductResponse, error) {
	if req.Id == "" {
		return nil, errors.New("product ID is required")
	}

	product, err := s.productUsecase.GetProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &inventory.ProductResponse{
		Product: s.domainToProto(product),
	}, nil
}

func (s *InventoryServer) UpdateProduct(ctx context.Context, req *inventory.UpdateProductRequest) (*inventory.ProductResponse, error) {
	product := &domain.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int(req.Stock),
		CategoryID:  req.CategoryId,
	}

	updatedProduct, err := s.productUsecase.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &inventory.ProductResponse{
		Product: s.domainToProto(updatedProduct),
	}, nil
}

func (s *InventoryServer) DeleteProduct(ctx context.Context, req *inventory.DeleteProductRequest) (*inventory.Empty, error) {
	if req.Id == "" {
		return nil, errors.New("product ID is required")
	}

	err := s.productUsecase.DeleteProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &inventory.Empty{}, nil
}

func (s *InventoryServer) ListProducts(ctx context.Context, req *inventory.ListProductsRequest) (*inventory.ListProductsResponse, error) {
	// Validate and set default pagination parameters
	page := int(req.Page)
	if page < 1 {
		page = 1
	}

	limit := int(req.Limit)
	if limit < 1 || limit > 100 {
		limit = 10
	}

	products, total, err := s.productUsecase.ListProducts(ctx, page, limit, req.CategoryId)
	if err != nil {
		return nil, err
	}

	// Convert domain products to protobuf products
	protoProducts := make([]*inventory.Product, len(products))
	for i, product := range products {
		protoProducts[i] = s.domainToProto(product)
	}

	return &inventory.ListProductsResponse{
		Products: protoProducts,
		Total:    int32(total),
		Page:     int32(page),
		Limit:    int32(limit),
	}, nil
}

func (s *InventoryServer) domainToProto(product *domain.Product) *inventory.Product {
	return &inventory.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		CategoryId:  product.CategoryID,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}
}

// Helper function to generate IDs
func generateID() string {
	return primitive.NewObjectID().Hex()
}
