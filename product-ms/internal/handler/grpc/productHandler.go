package grpc

import (
	"context"
	"product-ms/internal/domain"
	"product-ms/internal/dto"
	pb "product-ms/internal/infrastructure/client/pb"
	"product-ms/pkg"
	"time"
	
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	usecase   domain.ProductUsecase
}

func NewProductHandler(uc domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: uc}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	input := dto.CreateProductRequest{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
	}

	if err := pkg.ValidateRequest(input); err != nil {
		return nil, err
	}

	p := &domain.Product{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	p, err := h.usecase.CreateProduct(p)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{Product: toProtoProduct(p)}, nil
}

func (h *ProductHandler) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.GetProductByIDResponse, error) {
	p, err := h.usecase.GetProductByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetProductByIDResponse{Product: toProtoProduct(p)}, nil
}

func (h *ProductHandler) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	products, err := h.usecase.GetAllProducts()
	if err != nil {
		return nil, err
	}

	var pbProducts []*pb.Product
	for _, p := range products {
		pbProducts = append(pbProducts, toProtoProduct(p))
	}

	return &pb.GetAllProductsResponse{Products: pbProducts}, nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	input := dto.UpdateProductRequest{
		ID:          req.GetId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
	}

	if err := pkg.ValidateRequest(input); err != nil {
		return nil, err
	}

	p := &domain.Product{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		UpdatedAt:   time.Now(),
	}

	p, err := h.usecase.UpdateProduct(p)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductResponse{Product: toProtoProduct(p)}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	if err := h.usecase.DeleteProduct(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{Message: "product deleted successfully"}, nil
}

// Helper: domain → gRPC
func toProtoProduct(p *domain.Product) *pb.Product {
	return &pb.Product{
		Id:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
	}
}
