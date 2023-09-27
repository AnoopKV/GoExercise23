package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect2DB(_connectionString string) (*mongo.Client, error) {
	log.Println("Initiating Connection to Mongo DB")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoConnection := options.Client().ApplyURI(_connectionString)

	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println("Connected to database")
	return mongoClient, nil
}

func GetCollection(client *mongo.Client, collectionName string, dbName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
