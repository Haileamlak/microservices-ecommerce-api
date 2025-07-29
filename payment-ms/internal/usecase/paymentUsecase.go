package usecase

import (
	"context"
	"payment-ms/internal/domain"
	"payment-ms/internal/infrastructure"
	orderpb "payment-ms/internal/infrastructure/client/pb"
)

type paymentUseCase struct {
	repo           domain.PaymentRepository
	paymentService *infrastructure.PaymentService
	orderClient    orderpb.OrderServiceClient
}

func NewPaymentUseCase(repo domain.PaymentRepository, paymentService *infrastructure.PaymentService, orderClient orderpb.OrderServiceClient) domain.PaymentUseCase {
	return &paymentUseCase{
		repo:           repo,
		paymentService: paymentService,
		orderClient:    orderClient,
	}
}

func (uc *paymentUseCase) StartPayment(orderID string, amount float64, currency string) (string, *domain.AppError) {

	order, err := uc.orderClient.GetOrder(context.Background(), &orderpb.GetOrderRequest{Id: orderID})

	if err != nil {
		return "", domain.InternalErr("either order not found or error in getting order")
	}

	if order.Order.Status != "pending" {
		return "", domain.BadRequestErr("Order is not pending")
	}

	paymentLink, err := uc.paymentService.CreatePaymentLink(orderID, amount, currency)
	if err != nil {
		return "", domain.InternalErr("Failed to create payment link")
	}

	return paymentLink, nil
}

func (uc *paymentUseCase) UpdatePaymentStatus(orderID string, status domain.PaymentStatus) *domain.AppError {
	err := uc.repo.UpdateByOrderID(orderID, status)
	if err != nil {
		return domain.InternalErr("Failed to mark payment as paid")
	}

	return nil
}
