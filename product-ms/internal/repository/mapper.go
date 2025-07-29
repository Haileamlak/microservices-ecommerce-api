package repository

import (
    "product-ms/internal/domain"
    "time"
)

// Matches the MongoDB document structure
type productDocument struct {
    ID          string    `bson:"_id"` // UUID stored as string
    Title       string    `bson:"title"`
    Description string    `bson:"description"`
    Price       float64   `bson:"price"`
    CreatedAt   time.Time `bson:"created_at"`
    UpdatedAt   time.Time `bson:"updated_at"`
}

// Domain → MongoDB
func toProductDocument(p *domain.Product) *productDocument {
    return &productDocument{
        ID:          p.ID,
        Title:       p.Title,
        Description: p.Description,
        Price:       p.Price,
        CreatedAt:   p.CreatedAt,
        UpdatedAt:   p.UpdatedAt,
    }
}

// MongoDB → Domain
func toDomainProduct(doc *productDocument) *domain.Product {
    return &domain.Product{
        ID:          doc.ID,
        Title:       doc.Title,
        Description: doc.Description,
        Price:       doc.Price,
        CreatedAt:   doc.CreatedAt,
        UpdatedAt:   doc.UpdatedAt,
    }
}
