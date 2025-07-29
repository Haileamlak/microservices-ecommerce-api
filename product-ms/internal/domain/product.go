package domain

import "time"

type Product struct {
    ID          string
    Title       string
    Description string
    Price       float64
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
