package controllers

import (
	"log"
	"net/http"

	grpcclient "github.com/AnoopKV/GoExercise23/gRPCClient"
	"github.com/AnoopKV/GoExercise23/gRPCClient/proto/output/proto"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *grpcclient.GRPCCLientService
	//userService interfaces.IUser
}

func InitUserController(userService *grpcclient.GRPCCLientService) *UserController {
	return &UserController{userService: userService}
}

/*func InitUserController(userService interfaces.IUser) *UserController {
	//return &UserController{userService: userService}
}*/

func (u *UserController) HandleRegister(c *gin.Context) {
	var user proto.User
	if err := c.BindJSON(&user); err != nil { //Convert json into struct user
		log.Println("HandleRegister BindJSON Exception::" + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if res, err := u.userService.Register(&user); err != nil {
		log.Println("HandleRegister userService Exception::" + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if res != nil {
			c.IndentedJSON(http.StatusCreated, res)
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Already Registered"})
		}
	}
}

func (u *UserController) HandleLogin(c *gin.Context) {
	var user *proto.LoginRequest
	if err := c.BindJSON(&user); err != nil {
		log.Println("HandleLogin BindJSON() Exception::" + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	if res, err := u.userService.Login(user); err != nil {
		log.Println("HandleRegister Exception::" + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, res)
	}
}

func (u *UserController) HandleLogout(c *gin.Context) {
	c.JSON(200, gin.H{"data": "You have been logged out"})
}
