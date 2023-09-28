package interfaces

import "github.com/AnoopKV/GoExercise23/entities"

type IUser interface {
	Register(u *entities.User) (*entities.UserResponse, error)
	Login(u *entities.Login) (*entities.LoginResponse, error)
	Logout(token string) error
}
