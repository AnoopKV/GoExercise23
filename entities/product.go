package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name,required"`
	Category  string             `json:"category" bson:"category,required"`
	Quantity  int                `json:"quantity" bson:"quantity,required"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
