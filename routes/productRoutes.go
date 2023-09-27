package routes

import (
	"github.com/AnoopKV/GoExercise23/controllers"
	"github.com/AnoopKV/GoExercise23/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, p controllers.ProductController) {
	r.Use(middleware.Authenticate())
	r.Use(middleware.Authorize())
	r.POST("api/product/add", p.HandleAddProduct)    //Only authenticated as well as authorized(userType is user) only can logout
	r.GET("api/product/:id", p.HandleGetProductById) //Only authenticated as well as authorized(userType is user) only can logout
	r.GET("api/product/search", p.HandleSearch)      //Only authenticated as well as authorized(userType is user) only can logout
}
