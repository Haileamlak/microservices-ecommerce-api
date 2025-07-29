package domain

type ProductUsecase interface {
    CreateProduct(p *Product) (*Product, *AppError)
    GetProductByID(id string) (*Product, *AppError)
    GetAllProducts() ([]*Product, *AppError)
    UpdateProduct(p *Product) (*Product, *AppError)
    DeleteProduct(id string) *AppError
}
