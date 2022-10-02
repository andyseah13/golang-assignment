package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*********************
  PRODUCT FUNCTIONS
**********************/

func CreateProduct(c *gin.Context) {
	var latestProduct Product
	var newProduct Product
	err := c.BindJSON(&newProduct)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": err,
		})
		return
	}
	result := db.Last(&latestProduct)
	prodId := "PROD00001"
	hasRecord := true
	if result.Error != nil {
		// if it's NOT recordnotfound error, proceed to create product
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			hasRecord = false
		} else {
			c.IndentedJSON(500, gin.H{
				"message": result.Error,
			})
			return
		}
	}
	// assign a new product id based on the latest product id
	if hasRecord {
		latestId := latestProduct.Id
		latestCount, err := strconv.Atoi(latestId[len(latestId)-5:])
		if err != nil {
			c.IndentedJSON(500, gin.H{
				"message": err,
			})
		}
		prodId = "PROD" + fmt.Sprintf("%05d", latestCount+1)
	}
	newProduct.Id = prodId
	db.Save(&newProduct)
	c.IndentedJSON(200, newProduct)
}

func GetProductById(c *gin.Context) {
	var product Product
	productId := c.Param("id")

	result := db.First(&product, "id = ?", productId)
	if result.Error != nil {
		// if it's NOT recordnotfound error, proceed to create product
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(404, gin.H{
				"message": "Product not found",
			})
			return
		} else {
			c.IndentedJSON(500, gin.H{
				"message": result.Error,
			})
			return
		}
	}
	c.IndentedJSON(200, product)
}

func GetAllProducts(c *gin.Context) {
	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		c.IndentedJSON(500, gin.H{
			"message": result.Error,
		})
		return
	}
	productListResponse := ProductListResponse{
		Products: products,
	}

	c.IndentedJSON(200, productListResponse)
}

func UpdateProduct(c *gin.Context) {
	// get product to update
	var product Product
	productId := c.Param("id")
	result := db.First(&product, "id = ?", productId)
	if result.Error != nil {
		// if it's NOT recordnotfound error, proceed to create product
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(404, gin.H{
				"message": "Product not found",
			})
			return
		} else {
			c.IndentedJSON(500, gin.H{
				"message": result.Error,
			})
			return
		}
	}

	// bind fields to be updated to product object
	var inputProduct Product
	err := c.BindJSON(&inputProduct)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": err,
		})
		return
	}

	db.Model(&product).Updates(inputProduct)
	c.IndentedJSON(200, product)
}

/*********************
  ORDER FUNCTIONS
**********************/

func CreateOrder(c *gin.Context) {
	var orders []Order
	var product Product
	var newOrder Order
	err := c.BindJSON(&newOrder)

	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": err,
		})
		return
	}

	// assign order ID
	result := db.Find(&orders)
	count := result.RowsAffected
	orderId := "ORD" + fmt.Sprintf("%05d", count)
	newOrder.Id = orderId

	// get product name
	db.First(&product, "id = ?", newOrder.ProductId)
	newOrder.Status = "order placed"
	// newOrder.ProductName = product.Name

	db.Save(&newOrder)
	c.IndentedJSON(200, newOrder)
}
