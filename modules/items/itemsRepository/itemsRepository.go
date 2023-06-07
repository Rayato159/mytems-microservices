package itemsRepository

import "go.mongodb.org/mongo-driver/mongo"

type IItemsRepository interface{}

type itemsRepository struct {
	db *mongo.Client
}

func NewItemsRepository(db *mongo.Client) IItemsRepository {
	return &itemsRepository{db}
}
