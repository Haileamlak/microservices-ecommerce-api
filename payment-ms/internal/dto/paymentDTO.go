package dto

import (
	"payment-ms/internal/domain"
)

type InitiatePaymentRequest struct {
	OrderID  string  `json:"order_id" validate:"required,uuid4"`
	Amount   float64 `json:"amount" validate:"required,gt=0"`
	Currency string  `json:"currency" validate:"required,oneof=usd eur"`
}

func (r *InitiatePaymentRequest) ToDomain() *domain.Payment {
	return &domain.Payment{
		OrderID:  r.OrderID,
		Amount:   r.Amount,
		Status:   domain.PaymentStatusPending,
		Currency: r.Currency,
	}
}

type UpdatePaymentRequest struct {
	ID     string               `json:"id" validate:"required,uuid4"`
	Status domain.PaymentStatus `json:"status" validate:"required,oneof=pending completed cancelled"`
}
