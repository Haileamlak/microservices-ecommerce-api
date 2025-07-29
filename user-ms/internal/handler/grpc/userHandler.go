package grpc

import (
	"context"
	"user-ms/internal/domain"
	pb "user-ms/internal/infrastructure/client/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase domain.UserUsecase
}	

func NewUserHandler(uc domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
		token, err := h.usecase.LoginUser(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

func (h *UserHandler) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	id, err := h.usecase.VerifyUser(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyTokenResponse{Valid: true, UserId: id}, nil
}