package ports

import "hexagonal-product-api/internal/core/domain"

type ProductRepository interface {
	Save(product *domain.Product) error
	FindAll() ([]domain.Product, error)
	GetByID(id string) (*domain.Product, error)
}
