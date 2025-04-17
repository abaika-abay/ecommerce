package repository

import "inventory-service/internal/domain"

type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id string) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id string) error
	List(page, limit int, categoryID string) ([]*domain.Product, int, error)
}
