package ports

import "hexagonal-product-api/internal/core/domain"

type ProductService interface {
	Create(name string) (*domain.Product, error)
	ListAll() ([]domain.Product, error)
	GetByID(id string) (*domain.Product, error)
	GetProductDetail(id string) (*domain.ProductDetail, error)
}
