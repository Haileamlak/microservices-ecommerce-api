package domain

import "time"

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

type Payment struct {
	ID string
	OrderID string
	Amount float64
	Status PaymentStatus
	Currency string
	PaymentLink string
	CreatedAt time.Time
	UpdatedAt time.Time
}	

type PaymentResult struct {
	PaymentID string
	Status PaymentStatus
}