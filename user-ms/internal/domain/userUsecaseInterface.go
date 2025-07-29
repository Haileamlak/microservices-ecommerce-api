package domain

type UserUsecase interface { 		
	RegisterUser(user *User) (string, *AppError)
	LoginUser(email, password string) (string, *AppError)
	VerifyUser(token string) (string, *AppError)
}
