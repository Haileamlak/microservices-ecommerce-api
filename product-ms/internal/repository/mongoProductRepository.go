package repository

import (
	"context"
	"product-ms/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Database) domain.ProductRepository {
	return &mongoProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *mongoProductRepository) Create(p *domain.Product) (*domain.Product, *domain.AppError) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	doc := toProductDocument(p)

	result, err := r.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, domain.InternalErr("Failed to create product")
	}

	p.ID = result.InsertedID.(string)

	return p, nil
}

func (r *mongoProductRepository) GetByID(id string) (*domain.Product, *domain.AppError) {
	var doc productDocument
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return nil, domain.NotFoundErr("Product not found")
	}

	if err != nil {
		return nil, domain.InternalErr("Failed to get product by id")
	}

	return toDomainProduct(&doc), nil
}

func (r *mongoProductRepository) GetAll() ([]*domain.Product, *domain.AppError) {
	var docs []productDocument
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err == mongo.ErrNoDocuments {
		return nil, domain.NotFoundErr("No product found")
	}

	if err != nil {
		return nil, domain.InternalErr("Failed to get all products")
	}
	defer cursor.Close(context.TODO())

	if cursor.Err() != nil {
		return nil, domain.InternalErr("Failed to get all products")
	}

	if err := cursor.All(context.TODO(), &docs); err != nil {
		return nil, domain.InternalErr("Failed to get all products")
	}

	var products []*domain.Product
	for _, doc := range docs {
		products = append(products, toDomainProduct(&doc))
	}

	return products, nil
}

func (r *mongoProductRepository) Update(p *domain.Product) (*domain.Product, *domain.AppError) {
	p.UpdatedAt = time.Now()

	doc := toProductDocument(p)

	result, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": p.ID}, bson.M{"$set": doc})
	if err != nil {
		return nil, domain.InternalErr("Failed to update product")
	}

	if result.MatchedCount == 0 {
		return nil, domain.NotFoundErr("Product not found")
	}

	return p, nil
}

func (r *mongoProductRepository) Delete(id string) *domain.AppError {
	result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return domain.InternalErr("Failed to delete product")
	}

	if result.DeletedCount == 0 {
		return domain.NotFoundErr("Product not found")
	}

	return nil
}
