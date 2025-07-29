package domain

type ProductRepository interface {
    Create(product *Product) (*Product, *AppError)
    GetByID(id string) (*Product, *AppError)
    GetAll() ([]*Product, *AppError)
    Update(product *Product) (*Product, *AppError)
    Delete(id string) *AppError
}
