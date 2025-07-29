package repository

import (
    "payment-ms/internal/domain"
    "time"
)

// Matches the MongoDB document structure
type paymentDocument struct {
	ID string `bson:"_id"`
	OrderID string `bson:"order_id"`
	Amount float64 `bson:"amount"`
	Status string `bson:"status"`
	PaymentLink string `bson:"payment_link"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// Domain → MongoDB
func toPaymentDocument(p *domain.Payment) *paymentDocument {
    return &paymentDocument{
        ID:          p.ID,
        OrderID:       p.OrderID,
        Amount: p.Amount,
        Status: string(p.Status),
        PaymentLink: p.PaymentLink,
        CreatedAt:   p.CreatedAt,
        UpdatedAt:   p.UpdatedAt,
    }
}

// MongoDB → Domain
func toDomainPayment(doc *paymentDocument) *domain.Payment {
    return &domain.Payment{
        ID:          doc.ID,
        OrderID:       doc.OrderID,
        Amount: doc.Amount,
        Status: domain.PaymentStatus(doc.Status),
        PaymentLink: doc.PaymentLink,
        CreatedAt:   doc.CreatedAt,
        UpdatedAt:   doc.UpdatedAt,
    }
}
