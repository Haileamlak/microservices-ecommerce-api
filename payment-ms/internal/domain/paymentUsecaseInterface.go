package domain

type PaymentUseCase interface {
    StartPayment(orderID string, amount float64, currency string) (string, *AppError)
    UpdatePaymentStatus(orderID string, status PaymentStatus) *AppError
}