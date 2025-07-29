package repository

import (
    "order-ms/internal/domain"
    "time"
)

// Matches the MongoDB document structure
type orderDocument struct {
    ID          string    `bson:"_id"` // UUID stored as string
    UserID       string    `bson:"user_id"`
    ProductID string    `bson:"product_id"`
    TotalPrice float64   `bson:"total_price"`
    Status string `bson:"status"`
    CreatedAt   time.Time `bson:"created_at"`
    UpdatedAt   time.Time `bson:"updated_at"`
}

// Domain → MongoDB
func toOrderDocument(o *domain.Order) *orderDocument {
    return &orderDocument{
        ID:          o.ID,
        UserID:       o.UserID,
        ProductID: o.ProductID,
        TotalPrice: o.TotalPrice,
        Status: o.Status,
        CreatedAt:   o.CreatedAt,
        UpdatedAt:   o.UpdatedAt,
    }
}

// MongoDB → Domain
func toDomainOrder(doc *orderDocument) *domain.Order {
    return &domain.Order{
        ID:          doc.ID,
        UserID:       doc.UserID,
		ProductID: doc.ProductID,
        TotalPrice: doc.TotalPrice,
        Status: doc.Status,
        CreatedAt:   doc.CreatedAt,
        UpdatedAt:   doc.UpdatedAt,
    }
}
