package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnoopKV/GoExercise23/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	//get token from header
	//if no token, throw error n abort
	//if found, validate and proceed to next method
	return func(c *gin.Context) {
		log.Println("Authenticate check in-progress")

		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Request.Header.Set("uid", claims.Uid)
		c.Request.Header.Set("user_type", claims.User_type)
		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	//extract claims and check whether user has particular role
	//if yes proceed to next metgin
	//if no return 403/unauthorized error
	return func(c *gin.Context) {
		log.Println("Authorization check in-progress")
		user_type := c.Request.Header.Get("user_type")
		if user_type == "user" {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authorized"})
			c.Abort()
		}
	}
}
