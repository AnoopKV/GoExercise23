package interfaces

import (
	"github.com/AnoopKV/GoExercise23/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProduct interface {
	AddProduct(p *entities.Product) (*primitive.ObjectID, error)
	GetProductById(id primitive.ObjectID) (*entities.Product, error)
	SearchProducts(val string) (*[]entities.Product, error)
}
