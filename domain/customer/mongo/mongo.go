// Mongo is a mongo implementation of the Customer Repository
package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/percybolmer/tavern/domain/customer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db *mongo.Database
	// customer is used to store customers
	customer *mongo.Collection
}

// Create a new mongodb repository
func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// Find Metabot DB
	db := client.Database("ddd")
	customers := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"person.id": id})

	var c customer.Customer
	err := result.Decode(&c)
	if err != nil {
		return customer.Customer{}, err
	}
	return c, nil
}

func (mr *MongoRepository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := mr.customer.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MongoRepository) Update(c customer.Customer) error {
	panic("to implement")
}
