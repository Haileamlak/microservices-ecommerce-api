package domain

import "time"

type Order struct {
	ID string
	UserID string
	ProductID string
	TotalPrice float64
	Status string // pending, confirmed, shipped, delivered, cancelled
	CreatedAt time.Time
	UpdatedAt time.Time
}