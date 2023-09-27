package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/AnoopKV/GoExercise23/entities"
	"github.com/AnoopKV/GoExercise23/interfaces"
	"github.com/AnoopKV/GoExercise23/utils"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService interfaces.IProduct
}

func InitProductController(productService interfaces.IProduct) *ProductController {
	return &ProductController{productService: productService}
}

func (p *ProductController) HandleAddProduct(c *gin.Context) {
	var product entities.Product
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
		if primitiveId, err := utils.SetId(prodId); err == nil {
			res, _ := p.productService.GetProductById(*primitiveId)
			c.JSON(http.StatusFound, res)
		} else {
			log.Println("SetId Exception" + err.Error())
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User ID Parameter needed"})
		}
	}
}

func (p *ProductController) HandleSearch(c *gin.Context) {
	searchVal, _ok := c.GetQuery("name")
	if !_ok {
		log.Println("name Query String needed")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "key as well as value Query String needed"})
	} else {
		if res, _error := p.productService.SearchProducts(strings.TrimSpace(searchVal)); _error != nil {
			log.Println("SetSearch Exception" + _error.Error())
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": _error.Error()})
		} else {
			c.JSON(http.StatusFound, res)
		}

	}
}
