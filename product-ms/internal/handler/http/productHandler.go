package http

import (
	"encoding/json"
	"net/http"
	"product-ms/internal/domain"
	"product-ms/internal/dto"
	"product-ms/pkg"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	usecase   domain.ProductUsecase
}

func NewProductHandler(uc domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: uc}
}

// @Summary Create a product
// @Description Create a product
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param request body dto.CreateProductRequest true "Product"
// @Security BearerAuth
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pkg.ValidateRequest(req); err != nil {
		WriteError(w, err)
		return
	}

	product, err := h.usecase.CreateProduct(req.ToDomain())

	if err != nil {
		WriteError(w, err)
		return
	}

	response := dto.ProductResponseFromDomain(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// @Summary Get a product by ID
// @Description Get a product by ID
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param id path string true "Product ID"
// @Security BearerAuth
// @Success 200 {object} domain.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")

	if !pkg.ValidateID(ID) {
		WriteError(w, domain.BadRequestErr("Invalid product id"))
		return
	}

	product, err := h.usecase.GetProductByID(ID)
	if err != nil {
		WriteError(w, err)
		return
	}

	response := dto.ProductResponseFromDomain(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// @Summary Get all products
// @Description Get all products
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Security BearerAuth
// @Success 200 {object} []domain.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.usecase.GetAllProducts()
	if err != nil {
		WriteError(w, err)
		return
	}

	var response []dto.ProductResponse
	for _, product := range products {
		response = append(response, dto.ProductResponseFromDomain(product))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// @Summary Update a product
// @Description Update a product
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param id path string true "Product ID"
// @Param request body dto.UpdateProductRequest true "Product"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pkg.ValidateRequest(req); err != nil {
		WriteError(w, err)
		return
	}

	product, err := h.usecase.UpdateProduct(req.ToDomain())
	if err != nil {
		WriteError(w, err)
		return
	}

	response := dto.ProductResponseFromDomain(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// @Summary Delete a product
// @Description Delete a product
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param id path string true "Product ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")

	if !pkg.ValidateID(ID) {
		WriteError(w, domain.BadRequestErr("Invalid product id"))
		return
	}

	if err := h.usecase.DeleteProduct(ID); err != nil {
		WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

func WriteError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*domain.AppError); ok {
		http.Error(w, appErr.Message, appErr.StatusCode)
	} else {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
