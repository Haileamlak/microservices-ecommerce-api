package dto

import (
	"product-ms/internal/domain"
	"time"
)

type CreateProductRequest struct {
	Title       string  `json:"title" validate:"required,min=3,max=100"`
	Description string  `json:"description" validate:"required,min=5,max=500"`
	Price       float64 `json:"price" validate:"required,gt=0"`
}

func (r *CreateProductRequest) ToDomain() *domain.Product {
	return &domain.Product{
		Title:       r.Title,
		Description: r.Description,
		Price:       r.Price,
	}
}

type UpdateProductRequest struct {
	ID          string  `json:"id" validate:"required,uuid4"`
	Title       string  `json:"title" validate:"required,min=3,max=100"`
	Description string  `json:"description" validate:"required,min=5,max=500"`
	Price       float64 `json:"price" validate:"required,gt=0"`
}

func (r *UpdateProductRequest) ToDomain() *domain.Product {
	return &domain.Product{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Price:       r.Price,
	}
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
}

func ProductResponseFromDomain(product *domain.Product) ProductResponse {
	p := ProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
	}
	return p
}
