package domain

type UserRepository interface {
	CreateUser(user *User) (string,*AppError)
	GetUserByID(id string) (*User, *AppError)
	GetUserByEmail(email string) (*User, *AppError)
	UpdateUser(user *User) *AppError
	DeleteUser(id string) *AppError
}
