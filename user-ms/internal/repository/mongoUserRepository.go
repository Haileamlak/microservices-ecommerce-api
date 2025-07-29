package repository

import (
	"context"
	"time"
	"user-ms/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) domain.UserRepository {
	return &mongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *mongoUserRepository) CreateUser(user *domain.User) (string, *domain.AppError) {	
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	doc := toUserDocument(user)
	result, err := r.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", domain.InternalErr("Failed to create user")
	}

	user.ID = result.InsertedID.(string)

	return user.ID, nil
}

func (r *mongoUserRepository) GetUserByID(id string) (*domain.User, *domain.AppError) {
	var doc UserDocument
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return nil, domain.NotFoundErr("User not found")
	}
	if err != nil {
		return nil, domain.InternalErr("Failed to get user by ID")
	}

	return toDomainUser(&doc), nil
}

func (r *mongoUserRepository) GetUserByEmail(email string) (*domain.User, *domain.AppError) {
	var doc UserDocument
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&doc)

	if err == mongo.ErrNoDocuments {
		return nil, domain.NotFoundErr("User not found")
	}
	
	if err != nil {
		return nil, domain.InternalErr("Failed to get user by email")
	}


	return toDomainUser(&doc), nil
}

func (r *mongoUserRepository) UpdateUser(user *domain.User) *domain.AppError {
	user.UpdatedAt = time.Now()
	doc := toUserDocument(user)

	result, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": doc})
	if err != nil {
		return domain.InternalErr("Failed to update user")
	}

	if result.MatchedCount == 0 {
		return domain.NotFoundErr("User not found")
	}

	return nil
}

func (r *mongoUserRepository) DeleteUser(id string) *domain.AppError {
	result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return domain.InternalErr("Failed to delete user")
	}

	if result.DeletedCount == 0 {
		return domain.NotFoundErr("User not found")
	}

	return nil
}