package repository

import (
	"context"
	"order-ms/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) domain.OrderRepository {
	return &mongoOrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *mongoOrderRepository) CreateOrder(order *domain.Order) (string, *domain.AppError) {
	doc := toOrderDocument(order)
	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return "", domain.InternalErr("Failed to create order")
	}
	return result.InsertedID.(string), nil
}

func (r *mongoOrderRepository) GetOrderById(id string) (*domain.Order, *domain.AppError) {
	var doc orderDocument
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&doc)
	if err != nil {
		return nil, domain.NotFoundErr("Order not found")
	}
	return toDomainOrder(&doc), nil
}

func (r *mongoOrderRepository) GetOrdersByUserId(userId string) ([]*domain.Order, *domain.AppError) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userId})
	if err != nil {
		return nil, domain.InternalErr("Failed to get orders")
	}
	defer cursor.Close(context.Background())

	orders := make([]*domain.Order, 0)
	for cursor.Next(context.Background()) {
		var doc orderDocument
		if err := cursor.Decode(&doc); err != nil {
			return nil, domain.InternalErr("Failed to decode order")
		}
		orders = append(orders, toDomainOrder(&doc))
	}
	return orders, nil
}

func (r *mongoOrderRepository) UpdateOrderStatus(id string, status string) *domain.AppError {
	result, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return domain.InternalErr("Failed to update order status")
	}
	if result.MatchedCount == 0 {
		return domain.NotFoundErr("Order not found")
	}
	return nil
}

func (r *mongoOrderRepository) DeleteOrder(id string) *domain.AppError {
	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return domain.InternalErr("Failed to delete order")
	}
	if result.DeletedCount == 0 {
		return domain.NotFoundErr("Order not found")
	}
	return nil
}