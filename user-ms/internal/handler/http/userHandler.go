package http

import (
	"encoding/json"
	"log"
	"net/http"
	"user-ms/internal/domain"
	"user-ms/internal/dto"
	"user-ms/pkg"
)

type UserHandler struct {
	usecase domain.UserUsecase
}	

func NewUserHandler(usecase domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

// @Summary Register a user
// @Description Register a user
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "User"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pkg.ValidateRequest(req); err != nil {
		WriteError(w, err)
		return
	}

	id, err := h.usecase.RegisterUser(req.ToDomain())
	if err != nil {
		WriteError(w, err)
		return
	}	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// @Summary Login a user
// @Description Login a user
// @Accept json
// @Produce json
// @Param request body dto.LoginUserRequest true "User"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.usecase.LoginUser(req.Email, req.Password)
	if err != nil {
		WriteError(w, err)
		return
	}
	log.Println(token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// func (h *UserHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
// 	token := r.Header.Get("Authorization")
// 	if token == "" {
// 		WriteError(w, domain.UnauthorizedErr("Token is required"))
// 		return
// 	}

// 	id, err := h.usecase.VerifyUser(token)
// 	if err != nil {
// 		WriteError(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"id": id})
// }

func WriteError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*domain.AppError); ok {
		http.Error(w, appErr.Message, appErr.StatusCode)
	} else {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
