package routes

import (
	"github.com/AnoopKV/GoExercise23/controllers"
	"github.com/AnoopKV/GoExercise23/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, u controllers.UserController) {
	user := r.Group("/api/user")
	user.POST("/register", u.HandleRegister)
	user.POST("/login", u.HandleLogin)
	r.Use(middleware.Authenticate())
	r.Use(middleware.Authorize())
	r.GET("/api/user/logout", controllers.HandleLogout)
}
