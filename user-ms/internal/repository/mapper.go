package repository

import (
	"time"
	"user-ms/internal/domain"
)

type UserDocument struct {
	ID        string `bson:"_id"`
	Username  string `bson:"username"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func toUserDocument(user *domain.User) *UserDocument {
	return &UserDocument{
		ID: user.ID,
		Username: user.Username,	
		Email: user.Email,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,

	}
}

func toDomainUser(doc *UserDocument) *domain.User {
	return &domain.User{
		ID: doc.ID,
		Username: doc.Username,
		Email: doc.Email,
		Password: doc.Password,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
	}
}