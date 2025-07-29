package usecase

import (
	"product-ms/internal/domain"
	"product-ms/pkg"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{repo: r}
}
func (uc *productUsecase) CreateProduct(p *domain.Product) (*domain.Product, *domain.AppError) {
	if p.Title == "" || p.Price < 0 {
		return nil, domain.BadRequestErr("Invalid product")
	}

	p.ID = pkg.GenerateID()
	return uc.repo.Create(p)
}

func (uc *productUsecase) GetProductByID(id string) (*domain.Product, *domain.AppError) {
	return uc.repo.GetByID(id)
}

func (uc *productUsecase) GetAllProducts() ([]*domain.Product, *domain.AppError) {
	return uc.repo.GetAll()
}

func (uc *productUsecase) UpdateProduct(p *domain.Product) (*domain.Product, *domain.AppError) {
	if p.Title == "" || p.Price < 0 {
		return nil, domain.BadRequestErr("Invalid product")
	}
	return uc.repo.Update(p)
}

func (uc *productUsecase) DeleteProduct(id string) *domain.AppError {
	return uc.repo.Delete(id)
}
