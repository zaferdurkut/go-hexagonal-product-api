package services

import (
	"fmt"
	"github.com/google/uuid"
	"hexagonal-product-api/internal/core/domain"
	"hexagonal-product-api/internal/core/ports"
	"time"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Create(name string) (*domain.Product, error) {

	if name == "" {
		return nil, domain.ErrProductNameInvalid

	}

	product := &domain.Product{
		ID:   uuid.NewString(),
		Name: name,
	}
	err := s.repo.Save(product)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (s *ProductService) ListAll() ([]domain.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(id string) (*domain.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err

	}
	return product, nil

}

func (s *ProductService) GetProductDetail(id string) (*domain.ProductDetail, error) {
	startTime := time.Now()

	type ProductResult struct {
		Product *domain.Product
		Err     error
	}

	productChan := make(chan ProductResult, 1)

	type StockResult struct {
		Stock int
	}

	stockChan := make(chan StockResult, 1)

	go func() {
		product, err := s.repo.GetByID(id)
		productChan <- ProductResult{Product: product, Err: err}
	}()

	go func() {
		stock := s.fetchStockInfo(id)
		stockChan <- StockResult{Stock: stock}
	}()

	var productDetail domain.ProductDetail
	var productErr error

	for i := 0; i < 2; i++ {
		select {

		case res := <-productChan:
			if res.Err != nil {
				productErr = res.Err
			} else {
				productDetail.ID = res.Product.ID
				productDetail.Name = res.Product.Name
			}
		case res := <-stockChan:
			productDetail.Stock = res.Stock
		}
	}

	if productErr != nil {
		return nil, productErr
	}

	duration := time.Since(startTime)
	productDetail.FetchTime = fmt.Sprintf("%dms", duration.Milliseconds())

	return &productDetail, nil
}

func (s *ProductService) fetchStockInfo(productID string) int {
	time.Sleep(500 * time.Millisecond)
	return len(productID) * 10
}
