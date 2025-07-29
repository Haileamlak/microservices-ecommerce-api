package dto

import (
	"order-ms/internal/domain"
	"time"
)

type CreateOrderRequest struct {
	UserID string `json:"user_id" validate:"required,uuid4"`
	ProductID string `json:"product_id" validate:"required,uuid4"`
	TotalPrice float64 `json:"total_price" validate:"required,gt=0"`
}

func (r *CreateOrderRequest) ToDomain() *domain.Order {
	return &domain.Order{
		UserID: r.UserID,
		ProductID: r.ProductID,
		TotalPrice: r.TotalPrice,
	}
}

type UpdateOrderRequest struct {
	ID          string  `json:"id" validate:"required,uuid4"`
	Status string `json:"status" validate:"required,oneof=pending confirmed shipped delivered cancelled"`
}

func (r *UpdateOrderRequest) ToDomain() *domain.Order {
	return &domain.Order{
		ID:          r.ID,
		Status: r.Status,
	}
}

type OrderResponse struct {
	ID          string  `json:"id"`
	UserID string `json:"user_id"`
	ProductID string `json:"product_id"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt   string  `json:"created_at"`
}

func OrderResponseFromDomain(order *domain.Order) OrderResponse {
	o := OrderResponse{
		ID:          order.ID,
		UserID: order.UserID,
		ProductID: order.ProductID,
		TotalPrice: order.TotalPrice,
		CreatedAt:   order.CreatedAt.Format(time.RFC3339),
	}
	return o
}
