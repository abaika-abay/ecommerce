package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"inventory-service/internal/domain"
)

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &productRepository{
		collection: db.Collection("products"),
	}
}

func (r *productRepository) Create(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate new ObjectID if not provided
	if product.ID == "" {
		product.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, bson.M{
		"_id":         product.ID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
		"category_id": product.CategoryID,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	})

	return err
}

func (r *productRepository) FindByID(id string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result struct {
		ID          string    `bson:"_id"`
		Name        string    `bson:"name"`
		Description string    `bson:"description"`
		Price       float64   `bson:"price"`
		Stock       int       `bson:"stock"`
		CategoryID  string    `bson:"category_id"`
		CreatedAt   time.Time `bson:"created_at"`
		UpdatedAt   time.Time `bson:"updated_at"`
	}

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &domain.Product{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		Stock:       result.Stock,
		CategoryID:  result.CategoryID,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (r *productRepository) Update(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"category_id": product.CategoryID,
			"updated_at":  time.Now(),
		},
	}

	result, err := r.collection.UpdateByID(ctx, product.ID, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *productRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *productRepository) List(page, limit int, categoryID string) ([]*domain.Product, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build filter
	filter := bson.M{}
	if categoryID != "" {
		filter["category_id"] = categoryID
	}

	// Get total count
	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Set up pagination options
	opts := options.Find().
		SetSkip(int64((page - 1) * limit)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	// Execute query
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	for cursor.Next(ctx) {
		var result struct {
			ID          string    `bson:"_id"`
			Name        string    `bson:"name"`
			Description string    `bson:"description"`
			Price       float64   `bson:"price"`
			Stock       int       `bson:"stock"`
			CategoryID  string    `bson:"category_id"`
			CreatedAt   time.Time `bson:"created_at"`
			UpdatedAt   time.Time `bson:"updated_at"`
		}

		if err := cursor.Decode(&result); err != nil {
			return nil, 0, err
		}

		products = append(products, &domain.Product{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
			Price:       result.Price,
			Stock:       result.Stock,
			CategoryID:  result.CategoryID,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		})
	}

	return products, int(total), nil
}
