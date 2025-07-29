package domain

type PaymentRepository interface {
	Create(payment *Payment) (string, *AppError)
	Update(paymentID string, status PaymentStatus) *AppError
	UpdateByOrderID(orderID string, status PaymentStatus) *AppError
	GetByID(paymentID string) (*Payment, *AppError)
}