package http

import (
	"encoding/json"
	"net/http"
	"order-ms/internal/domain"
	"order-ms/internal/dto"
	"order-ms/pkg"

	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	usecase domain.OrderUsecase
}

func NewOrderHandler(uc domain.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: uc}
}

// @Summary Create an order
// @Description Create an order for a user
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param request body dto.CreateOrderRequest true "Order"
// @Security BearerAuth
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, domain.BadRequestErr("Invalid request body"))
		return
	}
	if err := pkg.ValidateRequest(&req); err != nil {
		WriteError(w, err)
		return
	}
	id, err := h.usecase.CreateOrder(req.ToDomain())
	if err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// @Summary Get an order by ID
// @Description Get an order by ID
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param id path string true "Order ID"
// @Security BearerAuth
// @Success 200 {object} domain.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || !pkg.ValidateID(id) {
		WriteError(w, domain.BadRequestErr("Invalid order ID"))
		return
	}
	order, err := h.usecase.GetOrderById(id)
	if err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// @Summary Get orders by user ID
// @Description Get orders by user ID
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param userId path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} []domain.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/user/{userId} [get]
func (h *OrderHandler) GetOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	if userId == "" || !pkg.ValidateID(userId) {
		WriteError(w, domain.BadRequestErr("Invalid user ID"))
		return
	}
	orders, err := h.usecase.GetOrdersByUserId(userId)
	if err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

// @Summary Update an order status
// @Description Update an order status
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param id path string true "Order ID"
// @Param request body dto.UpdateOrderRequest true "Status"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || !pkg.ValidateID(id) {
		WriteError(w, domain.BadRequestErr("Invalid order ID"))
		return
	}
	var req dto.UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, domain.BadRequestErr("Invalid request body"))
		return
	}
	if err := pkg.ValidateRequest(&req); err != nil {
		WriteError(w, err)
		return
	}
	err := h.usecase.UpdateOrderStatus(id, req.Status)
	if err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order status updated successfully"})
}

// @Summary Delete an order
// @Description Delete an order
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer token"
// @Param id path string true "Order ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || !pkg.ValidateID(id) {
		WriteError(w, domain.BadRequestErr("Invalid order ID"))
		return
	}
	err := h.usecase.DeleteOrder(id)
	if err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
}

func WriteError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*domain.AppError); ok {
		http.Error(w, appErr.Message, appErr.StatusCode)
	} else {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
