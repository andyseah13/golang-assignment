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
	r.GET("/product/:id", GetProductById)
	r.GET("/products", GetAllProducts)
	r.PATCH("/product/:id", UpdateProduct)

	r.POST("/order", CreateOrder)
	r.GET("/order/:id", GetOrderById)

	r.Run()
}
