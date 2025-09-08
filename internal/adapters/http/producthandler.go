package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hexagonal-product-api/internal/core/domain"
	"hexagonal-product-api/internal/core/ports"
	"net/http"
)

type ProductHandler struct {
	productService ports.ProductService
}

func NewProductHandler(service ports.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}

}

func (h *ProductHandler) Create(c *gin.Context) {
	var requestBody struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	product, err := h.productService.Create(requestBody.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusCreated, product)

}

func (h *ProductHandler) ListAll(c *gin.Context) {
	products, err := h.productService.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, product)

}

func (h *ProductHandler) GetProductDetail(c *gin.Context) {
	id := c.Param("id")

	detail, err := h.productService.GetProductDetail(id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product details"})
		return
	}

	c.JSON(http.StatusOK, detail)
}
