package controllers

import (
	"log"
	"net/http"
	"strings"

	grpcclient "github.com/AnoopKV/GoExercise23/gRPCClient"
	"github.com/AnoopKV/GoExercise23/gRPCClient/proto/output/proto"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *grpcclient.GRPCCLientService
	//productService interfaces.IProduct
}

func InitProductController(productService *grpcclient.GRPCCLientService) *ProductController {
	return &ProductController{productService: productService}
}

func (p *ProductController) HandleAddProduct(c *gin.Context) {
	var product proto.Product
	if err := c.BindJSON(&product); err != nil { //Convert json into struct user
		log.Println("HandleAddProduct BindJSON Exception::" + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if res, err := p.productService.AddProduct(&product); err != nil {
		log.Println("HandleRegister userService Exception::" + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if res != nil {
			c.IndentedJSON(http.StatusCreated, gin.H{"Id": res})
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Already Registered"})
		}
	}
}

func (p *ProductController) HandleGetProductById(c *gin.Context) {
	if prodId := c.Param("id"); prodId == "" {
		log.Println("User ID Parameter needed")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User ID Parameter needed"})
	} else {
		res, _ := p.productService.GetProductById(&proto.ProductValue{Val: prodId})
		c.JSON(http.StatusFound, res)
	}
}

func (p *ProductController) HandleSearch(c *gin.Context) {
	searchVal, _ok := c.GetQuery("name")
	if !_ok {
		log.Println("name Query String needed")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "key as well as value Query String needed"})
	} else {
		if res, _error := p.productService.SearchProduct(&proto.ProductValue{Val: strings.TrimSpace(searchVal)}); _error != nil {
			log.Println("SetSearch Exception" + _error.Error())
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": _error.Error()})
		} else {
			c.JSON(http.StatusFound, res)
		}
	}
}
