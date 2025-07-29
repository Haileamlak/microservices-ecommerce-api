package dto

import (
	"user-ms/internal/domain"
	"time"
)

type RegisterUserRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

func (r *RegisterUserRequest) ToDomain() *domain.User {
	return &domain.User{
		Username:  r.Username,
		Email:     r.Email,
		Password:  r.Password,
	}
}

type LoginUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}


func (r *LoginUserRequest) ToDomain() *domain.User {
	return &domain.User{
		Email:     r.Email,
		Password:  r.Password,
	}
}


type UpdateUserRequest struct {
	ID          string  `json:"id" validate:"required,uuid4"`
	Username  string `json:"username" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

func (r *UpdateUserRequest) ToDomain() *domain.User {
	return &domain.User{
		ID:        r.ID,
		Username:  r.Username,
		Email:     r.Email,
		Password:  r.Password,
	}
}

type UserResponse struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	CreatedAt   string  `json:"created_at"`
}

func UserResponseFromDomain(user *domain.User) UserResponse {
	p := UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
	}
	return p
}
