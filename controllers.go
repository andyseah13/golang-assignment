package main

import (
	"github.com/gin-gonic/gin"
  	"fmt"
)

func CreateProduct(c * gin.Context) {
	var products []Product;
	var newProduct Product;
	err := c.BindJSON(&newProduct);

	result := db.Find(&products)
	count := result.RowsAffected
	fmt.Println("count is ", count)

	prodId := "PROD" + fmt.Sprintf("%05d", count)
	newProduct.Id = prodId

   	if err != nil {
       c.JSON(400, gin.H {
			"message": err,
		});
		return
   }
	db.Save(&newProduct)
	c.JSON(200, newProduct);
}