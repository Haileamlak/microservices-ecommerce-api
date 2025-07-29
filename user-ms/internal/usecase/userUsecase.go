package usecase

import (
	"user-ms/internal/domain"
	"user-ms/pkg"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uc *userUsecase) RegisterUser(user *domain.User) (string, *domain.AppError) {

	existingUser, existError := uc.userRepo.GetUserByEmail(user.Email)

	if existingUser != nil {
		return "", domain.BadRequestErr("User already exists")
	}

	if existError.Code != "NOT_FOUND" {
		return "", existError
	}

	user.ID = pkg.GenerateID()
	hashedPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		return "", domain.InternalErr("Failed to hash password")
	}

	user.Password = hashedPassword
	return uc.userRepo.CreateUser(user)
}

func (uc *userUsecase) LoginUser(email, password string) (string, *domain.AppError) {
	user, err := uc.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !pkg.ComparePassword(user.Password, password) {
		return "", domain.UnauthorizedErr("Invalid password")
	}

	token, err := pkg.GenerateToken(user)
	if err != nil {
		return "", domain.InternalErr("Failed to generate token")
	}

	return token, nil
}

func (uc *userUsecase) VerifyUser(token string) (string, *domain.AppError) {
	id, err := pkg.VerifyToken(token)
	if err != nil {
		return "", domain.UnauthorizedErr("Invalid token")
	}

	return id, nil
}
