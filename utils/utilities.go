package utils

import (
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
**************************************
Extract ID from Premitivie Object ID
**************************************
*/
func GetId(result *mongo.InsertOneResult) string {
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex()
	}
	return ""
}

/*
**************************************
convert string to primitive ObjectID
**************************************
*/
func SetId(hexStr string) (*primitive.ObjectID, error) {
	if premitiveId, err := primitive.ObjectIDFromHex(hexStr); err != nil {
		log.Println(err.Error())
		return nil, err
	} else {
		return &premitiveId, nil
	}
}

/*
**********************************
Get Environment Variable values
**********************************
*/
func GetEnvVal(val string) string {
	return os.Getenv(val)
}

func ParseInt(val string) (int, error) {
	intVal, err := strconv.Atoi(val)
	return intVal, err
}
