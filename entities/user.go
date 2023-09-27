package entities

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	FirstName       string             `json:"firstName" bson:"firstName,required"`
	LastName        string             `json:"lastName" bson:"lastName,required"`
	Age             int                `json:"age" bson:"age,required"`
	Email           string             `json:"email" bson:"email,required"`
	Password        string             `json:"password" bson:"password,required"`
	ConfirmPassword string             `json:"confirmPassword" bson:"confirmPassword,required"`
	CreatedAt       primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt       primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type SignedDetails struct {
    Email      string
    First_name string
    Last_name  string
    Uid        string
    User_type  string
    jwt.StandardClaims
}

type Login struct {
	Email    string `json:"email" bson:"email,required"`
	Password string `json:"password" bson:"password,required"`
}

type LoginResponse struct {
	TokenId string `json:"tokenId"`
	Error   string `json:"error"`
}

type CustomerResponse struct {
	Id      primitive.ObjectID `json:"id"`
	TokenId string             `json:"tokenId"`
	Error   string             `json:"error"`
}

type UserResponse struct {
	Id        primitive.ObjectID `json:"id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Age       int                `json:"age"`
	Email     string             `json:"email"`
	CreatedAt primitive.DateTime `json:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt"`
}
