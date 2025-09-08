package main

import (
	"github.com/gin-gonic/gin"
	"hexagonal-product-api/internal/adapters/http"
	"hexagonal-product-api/internal/adapters/storage"
	"hexagonal-product-api/internal/core/services"
	"log"
)

func main() {
	productRepo := storage.NewInMemoryStorage()

	productService := services.NewProductService(productRepo)

	productHandler := http.NewProductHandler(productService)

	router := gin.Default()

	router.POST("api/products", productHandler.Create)
	router.GET("api/products", productHandler.ListAll)
	router.GET("api/products/:id", productHandler.GetByID)
	router.GET("api/products/:id/details", productHandler.GetProductDetail)

	log.Println("Server started on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
