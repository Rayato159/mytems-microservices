package users

import (
	"github.com/Rayato159/mytems-microservices/modules/items"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"id" json:"id"`
	Username string             `bson:"username" json:"username"`
	Level    int                `bson:"level,omitempty" json:"level"`
	Items    []*items.Item      `bson:"items,omitempty" json:"items"`
}
