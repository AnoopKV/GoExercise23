package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AnoopKV/GoExercise23/entities"
	"github.com/AnoopKV/GoExercise23/interfaces"
	"github.com/AnoopKV/GoExercise23/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	User *mongo.Collection
}

/*
******************************************************
Initiate User Service
******************************************************
*/
func InitUserService(userCollection *mongo.Collection) interfaces.IUser {
	return &UserService{User: userCollection}
}

/*
******************************************************
Register User Functionality
******************************************************
*/
func (u *UserService) Register(user *entities.User) (*entities.UserResponse, error) {
	user.Id = primitive.NewObjectID()

	if user.Password != user.ConfirmPassword {
		return nil, fmt.Errorf("password & confirmPassword should match")
	}

	if pwd, err := utils.BcryptString(user.Password); err != nil {
		log.Println("Error in UserService, BecryptString:: " + err.Error())
		return nil, err
	} else {
		user.Password, user.ConfirmPassword = pwd, pwd
	}
	curTime := primitive.NewDateTimeFromTime(time.Now())
	user.CreatedAt, user.UpdatedAt = curTime, curTime

	if res, _err := u.User.InsertOne(context.Background(), user); _err != nil {
		log.Println("Exception in UserService, Register():: " + _err.Error())
		return nil, _err
	} else {
		recordId := utils.GetId(res)
		if premtiveId, __err := utils.SetId(recordId); __err != nil {
			log.Println("Exception in UserService, setId():: " + __err.Error())
			return nil, __err
		} else {
			return &entities.UserResponse{Id: *premtiveId, FirstName: user.FirstName, LastName: user.LastName, Age: user.Age, Email: user.Email, User_Type: user.User_Type}, nil
		}
	}
}

/*
******************************************************
Login Functionality
******************************************************
*/
func (u *UserService) Login(Login *entities.Login) (*entities.LoginResponse, error) {
	//Get user info by email-id
	log.Println("Email ID::" + Login.Email)
	filter := bson.M{"email": bson.M{"$eq": Login.Email}}
	res := u.User.FindOne(context.Background(), filter)
	var _user *entities.User
	if res != nil {
		if err := res.Decode(&_user); err != nil {
			log.Println("UserService, Decode Error:: " + err.Error())
			return nil, err
		}
		//authenticate password
		if _err := utils.ComparePassword(_user.Password, Login.Password); _err == nil {
			//if success, generate jwt token
			if token, __err := utils.GenerateToken(_user.Email, _user.FirstName, _user.LastName, _user.Id.Hex(), _user.User_Type); __err != nil {
				fmt.Println("UserService, JWT Token Generation Exception:: " + __err.Error())
				return nil, __err
			} else {
				//return jwt token
				return &entities.LoginResponse{TokenId: token, Error: ""}, nil
			}
		} else {
			log.Println("UserService, PasswordCompare Error:: " + _err.Error())
			return nil, _err
		}
	}
	return nil, (fmt.Errorf("No Record Found!"))
}

/*
******************************************************
Logout Functionality
******************************************************
*/
func (u *UserService) Logout() error {
	return nil
}
