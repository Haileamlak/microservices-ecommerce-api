package repository

import (
	"context"
	"payment-ms/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	)

type mongoPaymentRepository struct {
	collection *mongo.Collection
}

func NewMongoPaymentRepository(db *mongo.Database) domain.PaymentRepository {
	return &mongoPaymentRepository{collection: db.Collection("payments")}
}

func (r *mongoPaymentRepository) Create(payment *domain.Payment) (string, *domain.AppError) {
	doc := toPaymentDocument(payment)

	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return "", domain.InternalErr("Failed to create payment")
	}

	return result.InsertedID.(string), nil
}


func (r *mongoPaymentRepository) Update(id string, status domain.PaymentStatus) *domain.AppError {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}

	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return domain.InternalErr("Failed to update payment status")
	}

	if result.MatchedCount == 0 {
			return domain.NotFoundErr("Payment not found")
	}

	return nil
}

func (r *mongoPaymentRepository) UpdateByOrderID(orderID string, status domain.PaymentStatus) *domain.AppError {
	filter := bson.M{"order_id": orderID}
	update := bson.M{"$set": bson.M{"status": status}}

	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return domain.InternalErr("Failed to update payment status")
	}

	if result.MatchedCount == 0 {
		return domain.NotFoundErr("Payment not found")
	}

	return nil
}


func (r *mongoPaymentRepository) GetByID(id string) (*domain.Payment, *domain.AppError) {		
	filter := bson.M{"_id": id}

	var doc paymentDocument
	err := r.collection.FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		return nil, domain.NotFoundErr("Payment not found")
	}

	return toDomainPayment(&doc), nil
}