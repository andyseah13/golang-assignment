package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	var latestProduct Product
	var newProduct Product
	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(400, gin.H{
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
			c.JSON(500, gin.H{
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
			c.JSON(500, gin.H{
				"message": err,
			})
		}
		prodId = "PROD" + fmt.Sprintf("%05d", latestCount+1)
	}
	newProduct.Id = prodId
	db.Save(&newProduct)
	c.JSON(200, newProduct)
}

func CreateOrder(c *gin.Context) {
	var orders []Order
	var product Product
	var newOrder Order
	err := c.BindJSON(&newOrder)

	if err != nil {
		c.JSON(400, gin.H{
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
	c.JSON(200, newOrder)
}
