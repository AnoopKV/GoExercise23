package routes

import (
	"github.com/AnoopKV/GoExercise23/controllers"
	"github.com/AnoopKV/GoExercise23/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, u controllers.UserController) {
	user := r.Group("/api/user")             //open API
	user.POST("/register", u.HandleRegister) //open API
	user.POST("/login", u.HandleLogin)       //open API
	r.Use(middleware.Authenticate())
	r.Use(middleware.Authorize())
	r.GET("/api/user/logout", u.HandleLogout) //Only authenticated as well as authorized(userType is user) only can logout
}
