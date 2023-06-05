package items

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Id          primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
}
