package services

import (
	"context"
	"log"
	"time"

	"github.com/AnoopKV/GoExercise23/entities"
	"github.com/AnoopKV/GoExercise23/interfaces"
	"github.com/AnoopKV/GoExercise23/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
}

func InitProductService(productCollection *mongo.Collection) interfaces.IProduct {
	return &ProductService{Product: productCollection}
}

// AddProduct implements interfaces.IProduct.
func (p *ProductService) AddProduct(product *entities.Product) (*primitive.ObjectID, error) {

	product.Id = primitive.NewObjectID()
	curTime := primitive.NewDateTimeFromTime(time.Now())
	product.CreatedAt, product.UpdatedAt = curTime, curTime

	if res, _err := p.Product.InsertOne(context.Background(), product); _err != nil {
		log.Println("Exception in ProductService, AddProduct():: " + _err.Error())
		return nil, _err
	} else {
		recordId := utils.GetId(res)
		if premtiveId, __err := utils.SetId(recordId); __err != nil {
			log.Println("Exception in ProductService, setId():: " + __err.Error())
			return nil, __err
		} else {
			return premtiveId, nil
		}
	}
}

// GetProductById implements interfaces.IProduct.
func (p *ProductService) GetProductById(objectId primitive.ObjectID) (*entities.Product, error) {
	var filter interface{}

	filter = bson.M{"_id": objectId} //bson.M{"_id": bson.M{"$eq": objectId}}
	//option = bson.D{{"_id", 0}}
	res := p.Product.FindOne(context.Background(), filter)
	var product = &entities.Product{}
	if res != nil {
		if err := res.Decode(product); err != nil {
			log.Println("Decode Error: " + err.Error())
		}
	}
	return product, nil
}

// SearchProducts implements interfaces.IProduct.
func (p *ProductService) SearchProducts(val string) (*[]entities.Product, error) {

	var filter interface{} //option
	filter = bson.M{"name": bson.M{"$eq": val}}
	var products []entities.Product
	if cursor, err := p.Product.Find(context.Background(), filter); err == nil {
		for cursor.Next(context.TODO()) {
			var result = &entities.Product{}
			if err := cursor.Decode(result); err != nil {
				log.Println(err)
				return nil, err
			} else {
				products = append(products, *result)
			}
		}
		return &products, nil
	} else {
		log.Println(err)
		return nil, err
	}
}
