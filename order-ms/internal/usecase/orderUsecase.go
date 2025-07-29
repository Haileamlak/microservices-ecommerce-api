package usecase

import (
	"context"
	"order-ms/internal/domain"
	"order-ms/internal/infrastructure/client/pb"
	"order-ms/pkg"
	"time"
)

type orderUsecase struct {
	repo domain.OrderRepository
	productServiceClient pb.ProductServiceClient
}

func NewOrderUsecase(r domain.OrderRepository, productServiceClient pb.ProductServiceClient) domain.OrderUsecase {
	return &orderUsecase{repo: r, productServiceClient: productServiceClient}
}

func (uc *orderUsecase) CreateOrder(o *domain.Order) (string, *domain.AppError) {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()

	o.ID = pkg.GenerateID()

	product, err := uc.productServiceClient.GetProductByID(context.Background(), &pb.GetProductByIDRequest{Id: o.ProductID})

	if err != nil {
		return "", domain.NewAppError(err.Error(), "Failed to get product", 500)
	}
	o.TotalPrice = product.Product.Price

	return uc.repo.CreateOrder(o)
}

func (uc *orderUsecase) GetOrderById(id string) (*domain.Order, *domain.AppError) {
	return uc.repo.GetOrderById(id)
}

func (uc *orderUsecase) GetOrdersByUserId(userId string) ([]*domain.Order, *domain.AppError) {
	return uc.repo.GetOrdersByUserId(userId)
}

func (uc *orderUsecase) UpdateOrderStatus(id string, status string) *domain.AppError {
	return uc.repo.UpdateOrderStatus(id, status)
}

func (uc *orderUsecase) DeleteOrder(id string) *domain.AppError {
	return uc.repo.DeleteOrder(id)
}