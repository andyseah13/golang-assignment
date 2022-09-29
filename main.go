package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	// initialize database connection 
  getDatabaseConnection()

	r := gin.Default()

	// Router
	r.POST("/product", CreateProduct)
	r.Run()
}