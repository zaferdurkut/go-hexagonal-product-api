package storage

import (
	"hexagonal-product-api/internal/core/domain"
	"sync"
)

type InMemoryStorage struct {
	products map[string]domain.Product
	mu       sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		products: make(map[string]domain.Product),
	}

}

func (s *InMemoryStorage) Save(product *domain.Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, value := range s.products {
		if value.Name == product.Name {
			return nil
		}
	}

	s.products[product.ID] = *product
	return nil
}

func (s *InMemoryStorage) FindAll() ([]domain.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]domain.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products, nil
}

func (s *InMemoryStorage) GetByID(id string) (*domain.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	product, ok := s.products[id]
	if !ok {
		return nil, domain.ErrProductNotFound
	}
	return &product, nil

}
