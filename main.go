package main

import (
	"log"

	config "github.com/AnoopKV/GoExercise23/configs"
	"github.com/AnoopKV/GoExercise23/controllers"
	gRPCserver "github.com/AnoopKV/GoExercise23/gRPCServer"
	"github.com/AnoopKV/GoExercise23/routes"
	"github.com/AnoopKV/GoExercise23/services"
	"github.com/AnoopKV/GoExercise23/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	server      *gin.Engine
)

func init() {
	log.Println("Welcome to Exercise!")
	godotenv.Load()
	if mongoClient, err = config.Connect2DB(utils.GetEnvVal("MONGO_CONNECTION_STRING")); err != nil {
		log.Panic(err.Error())
	}
}

func main() {
	server = gin.Default()
	server.GET("/", home)
	initializeUser()
	initializeProduct()
	initializeGRPC(utils.GetEnvVal("GRPC_SERVER_PORT"), utils.GetEnvVal("SECRET_KEY"))
	log.Println(utils.GetEnvVal("PORT"))
	server.Run(":" + utils.GetEnvVal("PORT"))
}
func home(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Ekart is Up and Running..."})
}

func initializeUser() {
	userCollection := config.GetCollection(mongoClient, utils.GetEnvVal("USER_COLLECTION_NAME"), utils.GetEnvVal("DB_NAME"))
	userService := services.InitUserService(userCollection)
	userController := controllers.InitUserController(userService)
	routes.UserRoutes(server, *userController)
}

func initializeProduct() {
	propductCollection := config.GetCollection(mongoClient, utils.GetEnvVal("PRODUCT_COLLECTION_NAME"), utils.GetEnvVal("DB_NAME"))
	productService := services.InitProductService(propductCollection)
	productController := controllers.InitProductController(productService)
	routes.ProductRoutes(server, *productController)
}

func initializeGRPC(port string, key string) {
	go gRPCserver.Start(port, key) //seperate thread
}
