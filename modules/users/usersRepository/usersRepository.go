package usersRepository

import "go.mongodb.org/mongo-driver/mongo"

type IUsersRepository interface{}

type usersRepository struct {
	db *mongo.Client
}

func NewUsersRepository(db *mongo.Client) IUsersRepository {
	return &usersRepository{db}
}
