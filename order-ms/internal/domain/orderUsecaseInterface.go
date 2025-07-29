package domain

type OrderUsecase interface {
	CreateOrder(order *Order) (string, *AppError)
	GetOrderById(id string) (*Order, *AppError)
	GetOrdersByUserId(userId string) ([]*Order, *AppError)
	UpdateOrderStatus(id string, status string) *AppError
	DeleteOrder(id string) *AppError
}
