package grpc

import (
	"context"
	"order-ms/internal/domain"
	pb "order-ms/internal/infrastructure/client/pb"
	"time"
)

type OrderHandler struct {	
	pb.UnimplementedOrderServiceServer
	usecase domain.OrderUsecase
}

func NewOrderHandler(uc domain.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: uc}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {		
	order, err := h.usecase.CreateOrder(&domain.Order{
		UserID: req.UserId,
		ProductID: req.ProductId,
		Status: "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: order}, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := h.usecase.GetOrderById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderResponse{Order: toProtoOrder(order)}, nil
}


func toProtoOrder(o *domain.Order) *pb.Order {
	return &pb.Order{
		Id: o.ID,
		UserId: o.UserID,
		ProductId: o.ProductID,
		Status: o.Status,
	}
}