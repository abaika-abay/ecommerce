package usecase

import (
	"context"
	"errors"
	"inventory-service/internal/domain"
	"inventory-service/internal/repository"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	ListProducts(ctx context.Context, page, limit int, categoryID string) ([]*domain.Product, int, error)
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (uc *productUsecase) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	// Validate required fields
	if product.Name == "" {
		return nil, errors.New("product name is required")
	}
	if product.Price <= 0 {
		return nil, errors.New("product price must be positive")
	}
	if product.Stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}

	// Generate ID if not provided
	if product.ID == "" {
		product.ID = generateID() // Implement your ID generation logic
	}

	err := uc.repo.Create(product)
	if err != nil {
		return nil, err
	}

	// Return the created product with populated fields (like timestamps)
	return uc.repo.FindByID(product.ID)
}

func (uc *productUsecase) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	if id == "" {
		return nil, errors.New("product ID is required")
	}

	product, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *productUsecase) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	if product.ID == "" {
		return nil, errors.New("product ID is required")
	}

	// Verify the product exists
	existing, err := uc.repo.FindByID(product.ID)
	if err != nil {
		return nil, err
	}

	// Apply updates
	if product.Name != "" {
		existing.Name = product.Name
	}
	if product.Description != "" {
		existing.Description = product.Description
	}
	if product.Price > 0 {
		existing.Price = product.Price
	}
	if product.Stock >= 0 {
		existing.Stock = product.Stock
	}
	if product.CategoryID != "" {
		existing.CategoryID = product.CategoryID
	}

	err = uc.repo.Update(existing)
	if err != nil {
		return nil, err
	}

	// Return the updated product
	return uc.repo.FindByID(product.ID)
}

func (uc *productUsecase) DeleteProduct(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("product ID is required")
	}

	// Verify the product exists
	_, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}

	return uc.repo.Delete(id)
}

func (uc *productUsecase) ListProducts(ctx context.Context, page, limit int, categoryID string) ([]*domain.Product, int, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	products, total, err := uc.repo.List(page, limit, categoryID)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// Helper function to generate IDs (replace with your preferred ID generation)
func generateID() string {
	return primitive.NewObjectID().Hex() // Using MongoDB's ObjectID generation
}
